package use_case

import (
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rohmanseo/golang-clean-arch/entity"
	"github.com/rohmanseo/golang-clean-arch/model"
	"github.com/rohmanseo/golang-clean-arch/repository"
	"github.com/rohmanseo/golang-clean-arch/security"
)

func NewAuthUseCase(tokenRepository *repository.ITokenRepository, userRepository *repository.IUserRepository, conn *nats.Conn, tkn *security.IJwtToken) IAuthUseCase {
	return &authUseCaseImpl{
		TokenRepository: *tokenRepository,
		UserRepository:  *userRepository,
		natsConn:        conn,
		tokenAuth:       *tkn,
	}
}

type authUseCaseImpl struct {
	TokenRepository repository.ITokenRepository
	UserRepository  repository.IUserRepository
	natsConn        *nats.Conn
	tokenAuth       security.IJwtToken
}

func (a *authUseCaseImpl) Register(request model.RegisterRequest) (model.RegisterResponse, error) {
	res, err := a.UserRepository.CreateUser(entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		fmt.Println(err)
		return model.RegisterResponse{}, err
	}
	a.natsConn.Publish("user.created", []byte("New user created"))
	return model.RegisterResponse{
		Id:    res.Id,
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

func (a *authUseCaseImpl) Login(request model.LoginRequest) (model.LoginResponse, error) {
	user, err := a.UserRepository.GetUser(entity.User{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return model.LoginResponse{}, err
	}
	if user.Id != 0 {
		token, err := a.TokenRepository.Create(user)
		if err != nil {
			return model.LoginResponse{}, err
		}
		return model.LoginResponse{AccessToken: token.AccessToken}, nil
	}
	return model.LoginResponse{}, errors.New("Invalid email/password")
}

func (a *authUseCaseImpl) Logout(request model.LogoutRequest) (bool, error) {
	isActive, _ := a.TokenRepository.Validate(entity.Token{AccessToken: request.AccessToken})
	if isActive {
		a.TokenRepository.Remove(entity.Token{AccessToken: request.AccessToken})
		return true, nil
	}
	return false, errors.New("Token not valid")
}
