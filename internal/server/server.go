package server

import (
	"go-rest-api-kafka/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	router := gin.Default()
	
	plansHandler := handlers.NewPlansHandler(db)
	
	router.GET("/health", handlers.HealthCheck)
	// router.PUT("/release/save", plansHandler.Release)
	// router.PATCH("/release/update", plansHandler.Release)
	// router.PATCH("/plan/update", plansHandler.Plans)
	router.PUT("/plan/save", plansHandler.CreatePlan)

	return &Server{
		router: router,
		db:     db,
	}
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}