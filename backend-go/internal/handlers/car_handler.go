package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"adcar-crm/backend-go/internal/models"
	"adcar-crm/backend-go/internal/service"
)

func GetCarsHandler(svc *service.CarService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cars, err := svc.GetCars(ctx.Request.Context())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, cars)
	}
}

func CreateCarHandler(svc *service.CarService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.NewCarRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad json"})
			return
		}

		car := models.Car{
			Brand:    req.Brand,
			Model:    req.Model,
			Year:     req.Year,
			PriceRub: req.PriceRub,
		}

		car, err := svc.CreateCar(ctx.Request.Context(), car)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, car)
	}
}
