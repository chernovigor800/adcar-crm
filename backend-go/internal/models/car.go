package models

type Car struct {
	Department        string   `json:"department" binding:"required"`
	Resource          string   `json:"resource" binding:"required"`
	FromCountry       string   `json:"from_country" binding:"required"`
	CarId             int64    `json:"car_id"`
	Link              string   `json:"link"`
	VehicleType       string   `json:"vehicle_type" binding:"required"`
	Vin               string   `json:"vin"`
	Make              string   `json:"make" binding:"required"`
	Model             string   `json:"model" binding:"required"`
	Month             int64    `json:"month"`
	Year              int64    `json:"year"`
	Age               int64    `json:"age"`
	BodyType          string   `json:"body_type"`
	IsRightSteering   bool     `json:"is_right_steering"`
	Color             string   `json:"color"`
	Trim              string   `json:"trim"`
	Mileage           int64    `json:"mileage"`
	Fuel              string   `json:"fuel"`
	EngineVolume      int64    `json:"engine_volume"`
	HorsePower        int64    `json:"horse_power"`
	Transmission      string   `json:"transmission"`
	DriveType         string   `json:"drive_type"`
	Photos            []string `json:"photos"`
	Price             int64    `json:"price" binding:"gte=0"`
	PriceCurrency     string   `json:"price_currency" binding:"required"`
	AdditionalContext string   `json:"additional_context"`
}