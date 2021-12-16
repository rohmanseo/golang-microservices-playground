package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rohmanseo/golang-clean-arch/model"
	"github.com/rohmanseo/golang-clean-arch/use_case"
	"net/http"
	"strings"
)

func NewAuthControllerImpl(authUseCase *use_case.IAuthUseCase) IAuthController {
	return &authControllerImpl{
		AuthUseCase: *authUseCase,
	}
}

//TODO: input validation, req&res mapper, response helper

type authControllerImpl struct {
	AuthUseCase use_case.IAuthUseCase
}

func (a *authControllerImpl) Register(ctx *gin.Context) {
	req := model.RegisterRequest{
		Name:     ctx.PostForm("name"),
		Email:    ctx.PostForm("email"),
		Password: ctx.PostForm("password"),
	}
	response, err := a.AuthUseCase.Register(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Data:   nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, model.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
	return
}

func (a *authControllerImpl) Login(ctx *gin.Context) {
	req := model.LoginRequest{
		Email:    ctx.PostForm("email"),
		Password: ctx.PostForm("password"),
	}
	response, err := a.AuthUseCase.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.ApiResponse{
			Code:   http.StatusUnauthorized,
			Status: "Error",
			Data:   nil,
		})
		fmt.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, model.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
	return
}

func (a *authControllerImpl) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	token = strings.Split(token, " ")[1]
	req := model.LogoutRequest{AccessToken: token}
	_, err := a.AuthUseCase.Logout(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.ApiResponse{
			Code:   http.StatusUnauthorized,
			Status: "Error",
			Data:   nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, model.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	})
	return
}
