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
	_, err := r.pool.Exec(ctx, 
		`INSERT INTO cars (
			department, resource, from_country, car_id, link, vehicle_type, vin, make, model,
			month, year, age, body_type, is_right_steering, color, trim, mileage,
			fuel, engine_volume, horse_power, transmission, drive_type, photos,
			price, price_currency, additional_context
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
                 $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)`,
		car.Department, car.Resource, car.FromCountry, car.CarId, car.Link, car.VehicleType, 
		car.Vin, car.Make, car.Model, car.Month, car.Year, car.Age, car.BodyType, 
		car.IsRightSteering, car.Color, car.Trim, car.Mileage, car.Fuel, car.EngineVolume, 
		car.HorsePower, car.Transmission, car.DriveType, car.Photos, car.Price, 
		car.PriceCurrency, car.AdditionalContext,
	)
	
	if err != nil {
		return models.Car{}, err
	}
	
	return car, nil
}

func (r *CarRepository) GetCars(ctx context.Context) ([]models.Car, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT department, resource, from_country, car_id, link, vehicle_type, vin, make,
		        model, month, year, age, body_type, is_right_steering, color, trim, mileage,
		        fuel, engine_volume, horse_power, transmission, drive_type, photos,
		        price, price_currency, additional_context
		 FROM cars ORDER BY car_id DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		// Scan in same order as SELECT; pgx uses positional binding
		err := rows.Scan(
			&car.Department, &car.Resource, &car.FromCountry, &car.CarId, &car.Link,
			&car.VehicleType, &car.Vin, &car.Make, &car.Model, &car.Month, &car.Year,
			&car.Age, &car.BodyType, &car.IsRightSteering, &car.Color, &car.Trim,
			&car.Mileage, &car.Fuel, &car.EngineVolume, &car.HorsePower, &car.Transmission,
			&car.DriveType, &car.Photos, &car.Price, &car.PriceCurrency, &car.AdditionalContext,
		)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, rows.Err()
}
