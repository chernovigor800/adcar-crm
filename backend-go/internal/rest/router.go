package rest

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"adcar-crm/backend-go/internal/handlers"
	"adcar-crm/backend-go/internal/service"
)

func SetupHTTPRouter(r *gin.Engine, svc *service.CarService) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	{
		api.GET("/cars",  handlers.GetCarsHandler(svc))   // тут handler
		api.POST("/cars", handlers.CreateCarHandler(svc)) // тут handler
	}
}
