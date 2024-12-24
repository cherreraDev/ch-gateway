package server

import (
	dependencycontainer "ch-gateway/internal/shared/dependencyContainer"

	"github.com/gin-gonic/gin"
)

func registerRoutes(s *Server, container dependencycontainer.Container) {
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
