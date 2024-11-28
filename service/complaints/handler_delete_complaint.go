package complaints

import (
	"net/http"
	"strconv"

	"github.com/EmiliodDev/gofeed/service/auth"
	"github.com/gin-gonic/gin"
)

func (h *Handler) handleDeleteComplaint(c *gin.Context) {
	employeeID := c.Request.Context().Value(auth.EmployeeKey)
	if employeeID == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	complaintIDStr := c.Param("id")
	complaintID, err := strconv.Atoi(complaintIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid complaint ID"})
		return
	}

	eID := employeeID.(int)
	complaint, err := h.store.GetComplaintByID(complaintID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "complaint not found"})
		return
	}

	if complaint.EmployeeID != eID {
		c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
		return
	}

	err = h.store.DeleteComplaint(complaintID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "complaint deleted"})
}