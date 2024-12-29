package server

import (
	dependencycontainer "ch-gateway/internal/shared/dependencyContainer"
	userRoutes "ch-gateway/internal/user/platform/server/routes"

	"github.com/gin-gonic/gin"
)

func registerRoutes(s *Server, container dependencycontainer.Container) {
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userRoutes.SetUp(s.engine, container)
}
