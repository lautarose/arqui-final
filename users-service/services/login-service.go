package services

import (
	"errors"
	"os"
	userCliente "user/clients"
	loginDto "user/dtos/login"

	"github.com/dgrijalva/jwt-go"
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
		err = errors.New("incorrect password")
		log.Println(err)
		return loginResp, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginReq.Username,
		"password": loginReq.Password,
	})

	secretKey := os.Getenv("PROY_JWT_KEY")

	var jwtKey = []byte(secretKey)

	loginResp.Token, _ = token.SignedString(jwtKey)

	return loginResp, nil
}
