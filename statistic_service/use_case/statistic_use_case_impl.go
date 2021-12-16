package use_case

import (
	"github.com/rohmanseo/golang-clean-arch/model"
	"github.com/rohmanseo/golang-clean-arch/repository"
)

func NewStatisticUseCaseImpl(statisticRepository *repository.IStatisticRepository) IStatisticUseCase {
	return &statisticUseCaseImpl{
		StatisticRepository: *statisticRepository,
	}
}

type statisticUseCaseImpl struct {
	StatisticRepository repository.IStatisticRepository
}

func (s *statisticUseCaseImpl) GetStatistic() (model.StatisticResponse, error) {
	totalUser, err := s.StatisticRepository.GetTotalUsers()
	totalTweet, err := s.StatisticRepository.GetTotalTweets()
	res := model.StatisticResponse{
		TotalUser:  totalUser,
		TotalTweet: totalTweet,
	}
	return res, err
}
