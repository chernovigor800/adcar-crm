package rest

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"adcar-crm/backend-go/internal/service"
)

func Setup(r *gin.Engine) {
	// ping — можно оставить здесь или в handlers/ping.go
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// группа /api
	api := r.Group("/api")
	{
		api.GET("/cars", service.GetCars)
		// api.POST("/cars", handlers.CreateCar)
	}
}