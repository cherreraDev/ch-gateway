package server

import (
	dependencycontainer "ch-gateway/internal/shared/dependencyContainer"
	"ch-gateway/internal/shared/platform/server/handlers"
	userRoutes "ch-gateway/internal/user/platform/server/routes"

	"github.com/gin-gonic/gin"
)

func registerRoutes(s *Server, container dependencycontainer.Container) {
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	s.engine.GET("/micro/ping", handlers.OtherServicePingHandler(container.Services.DiscoveryService))
	userRoutes.SetUp(s.engine, container)
}
