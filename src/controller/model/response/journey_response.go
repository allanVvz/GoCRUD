package response

type JourneyResponse struct {
	Id              string  `json:"id"`
	DriverId        string  `json:"driver_id"`
	VehicleId       string  `json:"vehicle_id"`
	InitialDate     string  `json:"initial_date"`
	FinalDate       string  `json:"final_date"`
	InitialLocation string  `json:"initial_location"`
	FinalLocation   string  `json:"final_location"`
	Duration        float64 `json:"total_value"`
}
