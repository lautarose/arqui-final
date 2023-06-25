package dto

type UsersResponseDto struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type UsersResponseDtos []UsersResponseDto
