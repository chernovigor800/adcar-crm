package rest

import (	
	"github.com/gin-gonic/gin"
	"adcar-crm/backend-go/internal/handlers"
	"adcar-crm/backend-go/internal/service"
)

func SetupHTTPRouter(r *gin.Engine, svc *service.CarService) {
	api := r.Group("/api/v1")
	cars := api.Group("/cars")
	{
		cars.GET("/",  handlers.GetCarsHandler(svc))
		cars.POST("/", handlers.CreateCarHandler(svc))
	}
}