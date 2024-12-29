package handlers

import (
	"ch-gateway/internal/user/domain"
	httprequests "ch-gateway/internal/user/platform/server/handlers/httpRequests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(service domain.LoginService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req httprequests.LoginRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		credentials := map[string]string{"username": req.Username, "password": req.Password}
		authResponse, err := service.Authenticate(credentials)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "err"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": authResponse})

	}
}
