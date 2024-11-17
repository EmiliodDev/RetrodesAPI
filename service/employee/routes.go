package employee

import (
	"github.com/EmiliodDev/gofeed/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	store	types.EmployeeStore
}

func NewHandler(store types.EmployeeStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/register", h.handleRegister)
	router.POST("/login", h.handleLogin)
}