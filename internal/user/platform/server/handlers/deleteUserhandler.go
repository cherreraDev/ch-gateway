package handlers

import (
	"ch-gateway/internal/user/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteUserHandler(service domain.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.DefaultQuery("id", "")

		parsedID, err := uuid.Parse(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "id param 'id' should be a valid UUID"})
			return
		}
		err = service.DeleteUser(parsedID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting the user"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": "user deleted successfully"})
	}
}
