package dto

type CreateUserDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Address  string `json:"address" form:"address" binding:"required"`
}
