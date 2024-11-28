package complaints

import (
	"net/http"

	"github.com/EmiliodDev/gofeed/service/auth"
	"github.com/EmiliodDev/gofeed/types"
	"github.com/EmiliodDev/gofeed/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) handleCreateComplaint(c *gin.Context) {
	employeeID := c.Request.Context().Value(auth.EmployeeKey)
	if employeeID == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	var payload types.CreateComplaintPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload1"})
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload2"})
		return
	}

	eID := employeeID.(int)
	payload.EmployeeID = eID

	err := h.store.CreateComplaint(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "complaint created"})
}