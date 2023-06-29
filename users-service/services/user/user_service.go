package services

import (
	"strconv"
	userCliente "user/clients/user"

	userDtos "user/dtos/user"
	userModel "user/models"
	jwtUtils "user/utils/jwt"
)

type userService struct{}

type userServiceInterface interface {
	GetUser(string) (userDtos.UsersResponseDto, error)
	InsertUser(userDtos.UserInsertDto) (userDtos.UserInsertDto, error)
	UpdateUser(auth string, updateUser userDtos.UserUpdateDto) (userDtos.UsersResponseDto, error)
	DeleteUser(auth string) (userDtos.UsersResponseDto, error)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) InsertUser(newUser userDtos.UserInsertDto) (userDtos.UserInsertDto, error) {
	var newUserModel userModel.User
	newUserModel.Name = newUser.Name
	newUserModel.LastName = newUser.LastName
	newUserModel.Email = newUser.Email
	newUserModel.UserName = newUser.UserName
	newUserModel.Pwd = newUser.Password

	err := userCliente.InsertUser(newUserModel)

	if err != nil {
		return userDtos.UserInsertDto{}, err
	}

	return newUser, nil
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

func (s *userService) UpdateUser(authToken string, updateUser userDtos.UserUpdateDto) (userDtos.UsersResponseDto, error) {
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

	if updateUser.Name == "" {
		updateUser.Name = user.Name
	}

	if updateUser.LastName == "" {
		updateUser.LastName = user.LastName
	}

	if updateUser.Password == "" {
		updateUser.Password = user.Pwd
	}

	user, err = userCliente.UpdateUser(id, updateUser)

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

func (s *userService) DeleteUser(authToken string) (userDtos.UsersResponseDto, error) {
	var userDto userDtos.UsersResponseDto

	// Verificar el token de autenticación
	claims, err := jwtUtils.VerifyToken(authToken)
	if err != nil {
		return userDto, err
	}

	// Obtener el ID del usuario del token
	id, err := strconv.Atoi(claims.Id)
	if err != nil {
		return userDto, err
	}

	// Eliminar el usuario
	user, err := userCliente.DeleteUser(id)
	if err != nil {
		return userDto, err
	}

	// Asignar los detalles del usuario eliminado al objeto de respuesta
	userDto.Id = user.UserID
	userDto.Email = user.Email
	userDto.LastName = user.LastName
	userDto.Name = user.Name
	userDto.UserName = user.UserName

	return userDto, nil
}
