package request

type UserRequest struct {
	Email    string `json:"email" binding:"required, email"`
	Password string `json:"password" binding:"required, min=2, max=10, containsNumber, containsUppercase, containsLowercase, containsSpecialCharacter"`
	Name     string `json:"name" binding:"required,min=4,max=50"`
	Age      int8   `json:"age" binding:"required, min=18, max=120"`
}
