package use_case

import "github.com/rohmanseo/golang-clean-arch/repository"

type msgUseCaseImpl struct {
	statisticRepo repository.IStatisticRepository
}

func (m *msgUseCaseImpl) UserCreatedMsg() {
	m.statisticRepo.AddUser()
}

func NewMsgUseCaseImpl(statisticRepository *repository.IStatisticRepository) IMsgUseCase {
	return &msgUseCaseImpl{
		statisticRepo: *statisticRepository,
	}
}
