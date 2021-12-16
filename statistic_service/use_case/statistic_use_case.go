package use_case

import (
	"github.com/rohmanseo/golang-clean-arch/model"
)

type IStatisticUseCase interface {
	GetStatistic() (model.StatisticResponse, error)
}
