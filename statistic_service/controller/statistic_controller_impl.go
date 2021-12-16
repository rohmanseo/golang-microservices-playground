package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rohmanseo/golang-clean-arch/model"
	"github.com/rohmanseo/golang-clean-arch/use_case"
	"net/http"
)

func NewStatisticController(statisticUseCase *use_case.IStatisticUseCase) IStatisticController {
	return &statisticController{
		StatisticUseCase: *statisticUseCase,
	}
}

//TODO: input validation, req&res mapper, response helper

type statisticController struct {
	StatisticUseCase use_case.IStatisticUseCase
}

func (s statisticController) GetStatistic(ctx *gin.Context) {
	statistic, err := s.StatisticUseCase.GetStatistic()
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
		Data:   statistic,
	})
	return
}
