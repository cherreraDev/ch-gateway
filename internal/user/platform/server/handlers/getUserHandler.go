package handlers

import (
	"ch-gateway/internal/user/domain"
	httprequests "ch-gateway/internal/user/platform/server/handlers/httpRequests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByIdHandler(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//ToDo  do it with query param
		var req httprequests.GetByIdRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		user, err := service.GetUserById(req.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding the user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"succes": user})
	}
}
func GetUserByUserNameHandler(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//ToDo  do it with query param
		var req httprequests.GetByUserNameRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		user, err := service.GetUserByUserName(req.UserName)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding the user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"succes": user})
	}
}
