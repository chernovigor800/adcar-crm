package models

type Car struct {
	Id       int    `json:"id"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Year     int    `json:"year"`
	PriceRub int    `json:"priceRub"`
}

type NewCarRequest struct {
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Year     int    `json:"year"`
	PriceRub int    `json:"priceRub"`
}
// type Car struct {	
// 	Department        string           `json:"department"`
// 	Resource          string           `json:"resource"`
// 	FromCountry       string           `json:"from_country"`
// 	CarId             int64            `json:"car_id"`
// 	Link              string           `json:"link"`
// 	VehicleType       string           `json:"vehicle_type"`
// 	Vin               string           `json:"vin"`
// 	Make              string           `json:"make"`
// 	Model             string           `json:"model"`
// 	Month             int64            `json:"month"`
// 	Year              int64            `json:"year"`
// 	Age               int64            `json:"age"`
// 	BodyType          string           `json:"body_type"`
// 	IsRightSteering   bool             `json:"is_right_steering"`		
// 	Color             string           `json:"color"`
// 	Trim              string           `json:"trim"`	
// 	Millage           int64            `json:"millage"`	
// 	Fuel              string           `json:"fuel"`
// 	EngineVolume      int64            `json:"engine_volume"`
// 	HorsePower        int64            `json:"horse_power"`
// 	Transmission      string           `json:"transmission"`
// 	DriveType         string           `json:"drive_type"`
// 	Photos            []string         `json:"photos"`	
// 	Price             int64            `json:"price"`
// 	PriceCurrency 	  string           `json:"price_currency"`	
// 	AdditionalContext string           `json:"additional_context"`	
// }	
