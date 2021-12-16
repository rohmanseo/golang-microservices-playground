package use_case

import "fmt"

type msgUseCaseImpl struct {
}

func (m *msgUseCaseImpl) TestMessageReceived(msg string) {
	fmt.Println("Received test msg", msg)
}

func NewMsgUseCaseImpl() IMsgUseCase {
	return &msgUseCaseImpl{}
}
