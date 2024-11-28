package complaints

import (
	"net/http"

	"github.com/EmiliodDev/gofeed/service/auth"
	"github.com/gin-gonic/gin"
)

func (h *Handler) handleGetComplaints(c *gin.Context) {
	employeeID := c.Request.Context().Value(auth.EmployeeKey)
	if employeeID == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	complaints, err := h.store.GetAllComplaints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"complaints": complaints})
}