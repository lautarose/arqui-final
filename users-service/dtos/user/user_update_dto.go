package dto

type UserUpdateDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
}

type UsersUpdateDto []UserUpdateDto
