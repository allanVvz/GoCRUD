package request

type DriverRequest struct {
	Id       string `json:"id"binding:"required" min:"1" max:"100"`
	Name     string `json:"name" min:"4" max:"100"`
	Rg	   string `json:"rg" min:"9" max:"9"`
	Registration string `json:"registration" min:"3" max:"5"`
	Salary  float64 `json:"salario" min:"10" max:"100000"`
}
