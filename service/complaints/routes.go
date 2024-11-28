package complaints

import (
	"github.com/EmiliodDev/gofeed/service/auth"
	"github.com/EmiliodDev/gofeed/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store types.ComplaintStore
}

func NewHandler(store types.ComplaintStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup, employeeStore types.EmployeeStore) {
	router.POST("/complaints", auth.WithJWTAuth(h.handleCreateComplaint, employeeStore))
	router.GET("/complaints", auth.WithJWTAuth(h.handleGetComplaints, employeeStore))
	router.PUT("/complaints/:id", auth.WithJWTAuth(h.handleUpdateComplaint, employeeStore))
	router.DELETE("/complaints/:id", auth.WithJWTAuth(h.handleDeleteComplaint, employeeStore))
}