package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=@#$!?*"`
	Name     string `json:"name" binding:"required,min=4,max=40"`
	Age      int8   `json:"age" binding:"required,min=12,max=150"`
}

type UserRequestUpdate struct {
	Name     string `json:"name" binding:"omitempty,min=4,max=40"`
	Age      int8   `json:"age" binding:"omitempty,min=12,max=150"`
}