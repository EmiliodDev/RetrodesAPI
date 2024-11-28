package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func GetTokenFromRequest(c *gin.Context) string {
    tokenAuth := c.GetHeader("Authorization")
    if tokenAuth != "" {
        if strings.HasPrefix(tokenAuth, "Bearer ") {
            return tokenAuth[len("Bearer "):]
        }
        return ""
    }

    tokenQuery := c.Query("token")
    if tokenQuery != "" {
        return tokenQuery
    }

    return ""
}