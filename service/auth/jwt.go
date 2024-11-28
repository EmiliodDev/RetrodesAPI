package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/EmiliodDev/gofeed/config"
	"github.com/EmiliodDev/gofeed/types"
	"github.com/EmiliodDev/gofeed/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const EmployeeKey contextKey = "employeeID"

func WithJWTAuth(handlerFunc gin.HandlerFunc, store types.EmployeeStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := utils.GetTokenFromRequest(c)

		token, err := validateJWT(tokenString)
		if err != nil {
			log.Printf("failed to validate token %v: ", err)
			permissionDenied(c)
			return 
		}

		if !token.Valid {
			log.Printf("invalid token")
			permissionDenied(c)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["employeeID"].(string)

		employeeID, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert employeeID to int %v: ", err)
			permissionDenied(c)
			return
		}

		e, err := store.GetEmployeeByID(employeeID)
		if err != nil {
			log.Printf("failed to get user by id %v: ", err)
			permissionDenied(c)
			return
		}

		ctx := context.WithValue(c.Request.Context(), EmployeeKey, e.ID)
		c.Request = c.Request.WithContext(ctx)

		handlerFunc(c)
	}
}

func CreateJWT(secret []byte, employeeID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"employeeID": strconv.Itoa(int(employeeID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})	
}

func permissionDenied(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
}