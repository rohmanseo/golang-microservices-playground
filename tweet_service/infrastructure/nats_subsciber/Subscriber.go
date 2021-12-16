package nats_subsciber

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/rohmanseo/golang-clean-arch/exception"
	"github.com/rohmanseo/golang-clean-arch/use_case"
	"strconv"
)

func SetupSubscriber(conn *nats.Conn, msgUseCase use_case.IMsgUseCase) {
	asyncSubscribe(conn, "check.subject", func(msg *nats.Msg) {
		msgUseCase.TestMessageReceived(string(msg.Data))
	})
	asyncSubscribe(conn, "tweet.total", func(msg *nats.Msg) {
		fmt.Println("Received tweet total subject", msg.Data)
		totalTweet := msgUseCase.GetTotalTweet()
		conn.Publish(msg.Reply, []byte(strconv.Itoa(totalTweet)))
	})
}

func asyncSubscribe(conn *nats.Conn, subject string, action nats.MsgHandler) {
	_, err := conn.Subscribe(subject, action)
	exception.PanicIfNeeded(err)
}
