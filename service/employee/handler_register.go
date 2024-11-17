package employee

import (
	"net/http"

	"github.com/EmiliodDev/gofeed/service/auth"
	"github.com/EmiliodDev/gofeed/types"
	"github.com/EmiliodDev/gofeed/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) handleRegister(c *gin.Context) {
	var payload types.RegisterEmployeePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid peyload"})
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	_, err := h.store.GetEmployeeByEmail(payload.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": "email already used"})
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	err = h.store.CreateEmployee(types.Employee{
		Name: payload.Name,
		LastName: payload.LastName,
		Email: payload.Email,
		Department: payload.Department,
		Position: payload.Position,
		Password: hashedPassword,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "employee created"})
}