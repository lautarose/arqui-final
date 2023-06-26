package services

import (
	"strconv"
	userCliente "user/clients"
	userDtos "user/dtos/user"
	jwtUtils "user/utils/jwt"
)

type userService struct{}

type userServiceInterface interface {
	GetUser(string) (userDtos.UsersResponseDto, error)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUser(authToken string) (userDtos.UsersResponseDto, error) {
	//controlar authToken
	var userDto userDtos.UsersResponseDto

	claims, err := jwtUtils.VerifyToken(authToken)

	if err != nil {
		return userDto, err
	}

	id, err := strconv.Atoi(claims.Id)

	if err != nil {
		return userDto, err
	}

	user, err := userCliente.GetUserById(id)

	if err != nil {
		return userDto, err
	}

	userDto.Id = user.UserID
	userDto.Email = user.Email
	userDto.LastName = user.LastName
	userDto.Name = user.Name
	userDto.UserName = user.UserName

	return userDto, nil
}
