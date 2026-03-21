package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"adcar-crm/backend-go/internal/models"
	// "adcar-crm/backend-go/internal/repository"
)

func GetCars(c *gin.Context) {
	cars := []models.Car{
		{Id: 1, Brand: "Lexus", Model: "Camry", Year: 2020, PriceRub: 1500000},
		{Id: 2, Brand: "Lada", Model: "Kalina", Year: 2015, PriceRub: 300000},
	}
	c.JSON(http.StatusOK, cars)
}

// func GetCars(c *gin.Context) {
// 	cars := repository.GetAllCars()
// 	c.JSON(http.StatusOK, cars)
// }

// func CreateCar(c *gin.Context) {
// 	var dto struct {
// 		Brand    string `json:"brand"`
// 		Model    string `json:"model"`
// 		Year     int    `json:"year"`
// 		PriceRub int    `json:"priceRub"`
// 	}
// 	if err := c.BindJSON(&dto); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json"})
// 		return
// 	}

// 	newCar := repository.CreateCar(dto.Brand, dto.Model, dto.Year, dto.PriceRub)
// 	c.JSON(http.StatusCreated, newCar)
// }
