package model

type User struct {
	Id        int    `json:"id"`
	Login     string `json:"login" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	RoleID    int    `json:"role_id"`
}
