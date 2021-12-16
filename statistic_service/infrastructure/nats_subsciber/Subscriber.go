package nats_subsciber

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rohmanseo/golang-clean-arch/exception"
	"github.com/rohmanseo/golang-clean-arch/use_case"
)

func SetupSubscriber(conn *nats.Conn, msgUseCase use_case.IMsgUseCase) {
	asyncSubscribe(conn, "user.created", func(msg *nats.Msg) {
		fmt.Println("Received user created", msg.Data)
		msgUseCase.UserCreatedMsg()
	})
}

func asyncSubscribe(conn *nats.Conn, subject string, action nats.MsgHandler) {
	_, err := conn.Subscribe(subject, action)
	exception.PanicIfNeeded(err)
}
