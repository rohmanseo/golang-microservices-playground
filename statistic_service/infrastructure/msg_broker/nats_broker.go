package msg_broker

import (
	"github.com/nats-io/nats.go"
	"github.com/rohmanseo/golang-clean-arch/config"
	"github.com/rohmanseo/golang-clean-arch/exception"
)

func NewNats(config config.IConfig) *nats.Conn {
	nc, err := nats.Connect(config.Get("NATS_URL"))
	exception.PanicIfNeeded(err)
	return nc
}
