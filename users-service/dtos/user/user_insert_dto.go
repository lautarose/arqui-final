package dto

type UserInsertDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UsersInsertDto []UserInsertDto
