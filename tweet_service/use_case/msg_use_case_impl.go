package use_case

import (
	"fmt"
	"github.com/rohmanseo/golang-clean-arch/repository"
)

type msgUseCaseImpl struct {
	repo repository.ITweetRepository
}

func (m *msgUseCaseImpl) GetTotalTweet() int {
	return m.repo.GetTotalTweet()
}

func (m *msgUseCaseImpl) TestMessageReceived(msg string) {
	fmt.Println("Received test msg", msg)
}

func NewMsgUseCaseImpl(tweetRepo *repository.ITweetRepository) IMsgUseCase {
	return &msgUseCaseImpl{repo: *tweetRepo}
}
