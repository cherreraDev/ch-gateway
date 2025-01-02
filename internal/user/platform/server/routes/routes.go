package routes

import (
	dependencycontainer "ch-gateway/internal/shared/dependencyContainer"
	"ch-gateway/internal/shared/platform/server/middleware"
	"ch-gateway/internal/user/platform/server/handlers"

	"github.com/gin-gonic/gin"
)

func SetUp(engine *gin.Engine, container dependencycontainer.Container) {
	userGroup := engine.Group("/api/v1/user")
	{
		userGroup.POST("/login", handlers.LoginHandler(container.Services.LoginService))

		userGroup.POST("/new", handlers.CreateUserHandler(container.Services.UserService))

		userGroup.GET("/:id", handlers.GetUserByIdHandler(container.Services.UserService))

		userGroup.GET("/name/:userName", handlers.GetUserByUserNameHandler(container.Services.UserService))

		userGroup.PUT("/:id", middleware.AuthMiddleware(container.SigningKey),
			handlers.UpdateUserHandler(container.Services.UserService))

		userGroup.DELETE("/:id", middleware.AuthMiddleware(container.SigningKey),
			handlers.DeleteUserHandler(container.Services.UserService))
	}
}
