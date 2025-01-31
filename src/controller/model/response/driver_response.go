package response

type DriverResponse struct {
	Id           string  `json:"id"`
	Status       bool    `json:"status"`
	Name         string  `json:"nome"`
	Rg           string  `json:"rg"`
	Registration string  `json:"registro"`
	Salary       float64 `json:"salario"`
}

// Error implements error.
func (d DriverResponse) Error() string {
	panic("unimplemented")
}
