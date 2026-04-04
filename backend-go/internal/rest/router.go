package rest

import (	
	"github.com/gin-gonic/gin"
	"adcar-crm/backend-go/internal/handlers"
	"adcar-crm/backend-go/internal/service"
)

func SetupHTTPRouter(r *gin.Engine, svc *service.CarService) {
	r.GET("/",   handlers.GetCarsHandler(svc))
	r.POST("/",  handlers.CreateCarHandler(svc))
}