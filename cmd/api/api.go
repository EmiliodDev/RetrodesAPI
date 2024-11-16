package api

import (
	"database/sql"
	"log"

	"github.com/EmiliodDev/gofeed/service/employee"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
    addr    string
    db      *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
    return &APIServer{
        addr:   addr,
        db:     db,
    }
}

func (s *APIServer) Run() error {
    router := gin.New()

    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    err := router.SetTrustedProxies(nil)
    if err != nil {
        log.Fatalf("Failed to set trusted proxies: %v", err)
    }

    api := router.Group("/api/v1")

    userStore := employee.NewStore(s.db)
    userHandler := employee.NewHandler(userStore)
    userHandler.RegisterRoutes(api)

    log.Println("Listening on: ", s.addr)

    return router.Run(s.addr)
}