package response

type VehicleResponse struct {
	Id       string  `json:"id"`
	Status   int8    `json:"status"`
	DriverId string  `json:"driver_id"`
	Model    string  `json:"model"`
	Brand    string  `json:"brand"`
	Plate    string  `json:"plate"`
	Year     string  `json:"year"`
	Price    float64 `json:"price"`
}
