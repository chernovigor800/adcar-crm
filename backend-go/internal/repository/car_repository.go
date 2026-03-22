package repository

import (
	"context"
	"adcar-crm/backend-go/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CarRepository struct {
	pool *pgxpool.Pool
}

func NewCarRepository(pool *pgxpool.Pool) *CarRepository {
	return &CarRepository{pool: pool}
}

func (r *CarRepository) CreateCar(ctx context.Context, car models.Car) (models.Car, error) {
	var id int
	err := r.pool.QueryRow(
		ctx,
		`INSERT INTO cars (brand, model, year, price_rub) VALUES ($1, $2, $3, $4) RETURNING id`,
		car.Brand, car.Model, car.Year, car.PriceRub,
	).Scan(&id)

	if err != nil {
		return models.Car{}, err
	}

	car.Id = id
	return car, nil
}

func (r *CarRepository) GetCars(ctx context.Context) ([]models.Car, error) {
	rows, err := r.pool.Query(
		ctx,
		`SELECT id, brand, model, year, price_rub FROM cars ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		err := rows.Scan(&car.Id, &car.Brand, &car.Model, &car.Year, &car.PriceRub)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, rows.Err()
}
