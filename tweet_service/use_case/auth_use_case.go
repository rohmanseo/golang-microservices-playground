package use_case

import "github.com/rohmanseo/golang-clean-arch/model"

type IAuthUseCase interface {
	Login(request model.LoginRequest) (model.LoginResponse, error)
	Register(request model.RegisterRequest) (model.RegisterResponse, error)
	Logout(request model.LogoutRequest) (bool, error)
}
