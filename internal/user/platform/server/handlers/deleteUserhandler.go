package handlers

import (
	"ch-gateway/internal/user/domain"
	httprequests "ch-gateway/internal/user/platform/server/handlers/httpRequests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//ToDo  do it with query param
		var req httprequests.DeleteUserRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		err := service.DeleteUser(req.Id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting the user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": "user deleted successfully"})
	}
}
