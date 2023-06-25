package services

import (
	userCliente "user/clients"
	userDtos "user/dtos/user"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (userDtos.UserDto, error)
	GetUserByUsername(string) (userDtos.UserDto, error)
	GetUsers() (userDtos.UsersResponseDtos, error)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (userDtos.UserDto, error) {

	user, err := userCliente.GetUserById(id)
	var userDto userDtos.UserDto

	if err != nil {
		return userDto, err
	}
	if user.UserID == 0 {
		return userDto, nil
	}
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.UserID
	userDto.Email = user.Email
	userDto.Password = user.Pwd
	return userDto, nil
}

func (s *userService) GetUsers() (userDtos.UsersResponseDtos, error) {
	users, err := userCliente.GetUsers()
	var usersDto userDtos.UsersResponseDtos

	if err != nil {
		return usersDto, err
	}

	for _, user := range users {
		var userDto userDtos.UsersResponseDto
		userDto.Name = user.Name
		userDto.LastName = user.LastName
		userDto.UserName = user.UserName
		userDto.Id = user.UserID
		userDto.Email = user.Email
		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) GetUserByUsername(username string) (userDtos.UserDto, error) {

	user, err := userCliente.GetUserByUsername(username)
	var userDto userDtos.UserDto

	if err != nil {
		return userDto, err
	}

	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.UserName
	userDto.Id = user.UserID
	userDto.Email = user.Email
	return userDto, nil
}
