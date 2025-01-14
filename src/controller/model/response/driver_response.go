package response

type DriverResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Rg	   string `json:"rg"`
	Registration string `json:"registration"`
	Salary  float64 `json:"salario"`
}
