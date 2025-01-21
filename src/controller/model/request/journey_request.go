package request

type JourneyRequest struct {
	Id              string  `json:"id"binding:"required",min:"1",max:"999"`
	DriverId        string  `json:"driver_id"binding:"required", min:"1", max:"999"`
	VehicleId       string  `json:"vehicle_id"binding:"required" min:"1" max:"999"`
	InitialDate     string  `json:"initial_date"binding:"required" date"`
	FinalDate       string  `json:"final_date"binding:"required" date"`
	InitialLocation string  `json:"initial_location"binding:"required" min:"1" max:"100"`
	FinalLocation   string  `json:"final_location"binding:"required" min:"1" max:"100"`
	Duration        float64 `json:"total_value"binding:"required" min:"1" max:"99999"`
}
