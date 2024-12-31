package routes

import (
	dependencycontainer "ch-gateway/internal/shared/dependencyContainer"
	"ch-gateway/internal/user/platform/server/handlers"

	"github.com/gin-gonic/gin"
)

func SetUp(engine *gin.Engine, container *dependencycontainer.Container) {
	engine.POST("/login", handlers.LoginHandler(container.Services.LoginService))
}
