package request

type JourneyRequest struct {
	DriverID        string  `json:"driver_id" binding:"required" min:"1" max:"999"`
	VehicleID       string  `json:"vehicle_id" binding:"required" min:"1" max:"999"`
	InitialDate     string  `json:"initial_date" binding:"required" date:"true"`
	FinalDate       string  `json:"final_date" binding:"required" date:"true"`
	InitialLocation string  `json:"initial_location" binding:"required" min:"1" max:"100"`
	FinalLocation   string  `json:"final_location" binding:"required" min:"1" max:"100"`
	TotalValue      float64 `json:"total_value" binding:"required" min:"1" max:"99999"`
}
