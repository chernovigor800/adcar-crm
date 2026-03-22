package service

import (
	"context"
	"adcar-crm/backend-go/internal/models"
	"adcar-crm/backend-go/internal/repository"
)

type CarService struct {
	repo *repository.CarRepository
}

// фабричная функция, создаёт сервис один раз.
func NewCarService(repo *repository.CarRepository) *CarService {
	return &CarService{repo: repo}
}

func (s *CarService) CreateCar(ctx context.Context, car models.Car) (models.Car, error) {
	return s.repo.CreateCar(ctx, car)
}

func (s *CarService) GetCars(ctx context.Context) ([]models.Car, error) {
	return s.repo.GetCars(ctx)
}