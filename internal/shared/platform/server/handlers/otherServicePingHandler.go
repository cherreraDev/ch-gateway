package handlers

import (
	"ch-gateway/internal/shared/domain/discovery"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OtherServicePingHandler(service discovery.DiscoveryServer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url, err := service.GetServiceAdress("tasks-service")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error discovering ping service"})
			return
		}

		resp, err := http.Get(fmt.Sprintf("%s/api/v1/ping", url))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error contacting ping service"})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response"})
			return
		}
		ctx.Data(resp.StatusCode, "application/json", body)
	}
}
