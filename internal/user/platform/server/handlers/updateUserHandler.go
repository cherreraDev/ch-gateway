package handlers

import (
	"ch-gateway/internal/user/domain"
	httprequests "ch-gateway/internal/user/platform/server/handlers/httpRequests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req httprequests.UpdateUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		err := service.UpdateUser(req.Id, req.UserName, req.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating the user"})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"success": "User updated succesfully"})
	}
}
