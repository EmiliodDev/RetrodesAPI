package employee

import (
	"fmt"
	"net/http"

	"github.com/EmiliodDev/gofeed/config"
	"github.com/EmiliodDev/gofeed/service/auth"
	"github.com/EmiliodDev/gofeed/types"
	"github.com/EmiliodDev/gofeed/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) handleLogin(c *gin.Context) {
	var payload types.LoginEmployeePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid payload: %v", errors)})
		return
	}

	e, err := h.store.GetEmployeeByEmail(payload.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not found, invalid email or password"})
		return
	}

	if !auth.ComparePasswords(e.Password, []byte(payload.Password)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email or password"})
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, e.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	isHR := e.Department == "Human Resources"

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"isHR":  isHR,
	})
}