package request


type VehicleRequest struct {
	Id        string `json:"id" binding:"required" min:"1" max:"999"`
	DriverId  string `json:"driver_id" binding:"required" min:"1" max:"999"`
	Model     string `json:"model" 	binding:"required" min:"1" max:"99"`
	Brand     string `json:"brand" 		binding:"required" min:"1" max:"99"`
	Plate    string `json:"plate" 		binding:"required" min:"1" max:"99"`
	Year     string `json:"year" 		binding:"required" min:"1" max:"99"`
	Price	float64 `json:"price" 		binding:"required" min:"1" max:"99999"`
}