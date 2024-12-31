package handlers

import (
	"ch-gateway/internal/user/domain"
	httprequests "ch-gateway/internal/user/platform/server/handlers/httpRequests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req httprequests.CreateUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		err := service.CreateUser(req.Id, req.UserName, req.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating the user"})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"success": "User created succesfully"})
	}
}
