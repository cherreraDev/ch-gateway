package routes

import (
	dependencycontainer "ch-gateway/internal/shared/dependencyContainer"
	"ch-gateway/internal/user/platform/server/handlers"

	"github.com/gin-gonic/gin"
)

func SetUp(engine *gin.Engine, container dependencycontainer.Container) {
	userGroup := engine.Group("/api/v1/user")
	{
		userGroup.POST("/login", handlers.LoginHandler(container.Services.LoginService))
		userGroup.POST("/new", handlers.CreateUserHandler(container.Services.UserService))
	}
}
