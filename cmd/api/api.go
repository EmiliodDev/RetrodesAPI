package api

import (
	"database/sql"
	"log"
	"time"

	c_service "github.com/EmiliodDev/gofeed/service/complaints"
	e_service "github.com/EmiliodDev/gofeed/service/employee"
	"github.com/gin-contrib/cors"
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

    configCors(router)
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    err := router.SetTrustedProxies(nil)
    if err != nil {
        log.Fatalf("Failed to set trusted proxies: %v", err)
    }

    // Groups
    api := router.Group("/api/v1")

    // Stores
    userStore := e_service.NewStore(s.db)
    complaintsStore := c_service.NewStore(s.db)

    // Handlers
    userHandler := e_service.NewHandler(userStore)
    complaintsHandler := c_service.NewHandler(complaintsStore)

    // Registser routes
    userHandler.RegisterRoutes(api)
    complaintsHandler.RegisterRoutes(api, userStore)

    log.Println("Listening on: ", s.addr)

    return router.Run(s.addr)
}

func configCors(r *gin.Engine) {
    r.Use(cors.New(cors.Config{
        AllowOrigins:   []string{"http://localhost:5173"},
        AllowMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:   []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:  []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
}