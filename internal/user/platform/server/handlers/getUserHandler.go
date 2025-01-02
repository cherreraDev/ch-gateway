package handlers

import (
	"ch-gateway/internal/user/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserByIdHandler(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.DefaultQuery("id", "")

		parsedID, err := uuid.Parse(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'id' debe ser un UUID válido"})
			return
		}

		user, err := service.GetUserById(parsedID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding the user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"succes": user})
	}
}
func GetUserByUserNameHandler(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userNameParam := ctx.DefaultQuery("userName", "")
		user, err := service.GetUserByUserName(userNameParam)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding the user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"succes": user})
	}
}
