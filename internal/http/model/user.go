package model

type User struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
