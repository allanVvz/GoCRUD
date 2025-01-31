package request

type DriverRequest struct {
	Id           string  `json:"id"`
	Status       bool    `json:"status" min:"1" max:"1"`
	Name         string  `json:"name" min:"3" max:"99"`
	Rg           string  `json:"rg" min:"9" max:"9"`
	Registration string  `json:"registration" min:"3" max:"5"`
	Salary       float64 `json:"salary" min:"10" max:"100000"`
}
