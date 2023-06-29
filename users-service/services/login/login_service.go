package services

import (
	"errors"
	"fmt"

	userCliente "user/clients/user"
	loginDto "user/dtos/login"
	jwtUtils "user/utils/jwt"

	log "github.com/sirupsen/logrus"
)

type loginService struct{}

type loginServiceInterface interface {
	Login(loginDto.LoginRequestDto) (loginDto.LoginResponseDto, error)
}

var (
	LoginService loginServiceInterface
)

func init() {
	LoginService = &loginService{}
}

func (s *loginService) Login(loginReq loginDto.LoginRequestDto) (loginDto.LoginResponseDto, error) {
	username := loginReq.Username
	pass := loginReq.Password
	log.Debug("pass", pass)
	user, err := userCliente.GetUserByUsername(username)
	var loginResp loginDto.LoginResponseDto

	if err != nil {
		log.Println(err)
		return loginResp, err
	}

	if user.Pwd != pass {
		err = errors.New("incorrect credentials")
		log.Println(err)
		return loginResp, err
	}
	
	token, err := jwtUtils.GenerateToken(user.UserID)
	if err != nil {
		err = errors.New("error generating token")
		return loginResp, err
	}

	fmt.Println("genera token")

	loginResp.Token = token

	return loginResp, nil
}
