package handlers

import (	
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
		var req models.Car
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad json"})
			return
		}

		car := models.Car{
		    Department:      req.Department,
		    Resource:        req.Resource,
		    FromCountry:     req.FromCountry,
		    CarId:           req.CarId,
		    Link:            req.Link,
		    VehicleType:     req.VehicleType,
		    Vin:             req.Vin,
		    Make:            req.Make,
		    Model:           req.Model,
		    Month:           req.Month,
		    Year:            req.Year,
		    Age:             req.Age,
		    BodyType:        req.BodyType,
		    IsRightSteering: req.IsRightSteering,
		    Color:           req.Color,
		    Trim:            req.Trim,
		    Mileage:         req.Mileage,
		    Fuel:            req.Fuel,
		    EngineVolume:    req.EngineVolume,
		    HorsePower:      req.HorsePower,
		    Transmission:    req.Transmission,
		    DriveType:       req.DriveType,
		    Photos:          req.Photos,
		    Price:           req.Price,
		    PriceCurrency:   req.PriceCurrency,
		    AdditionalContext: req.AdditionalContext,
		}

		car, err := svc.CreateCar(ctx.Request.Context(), car)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, car)
	}
}
