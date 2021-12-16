package config

import (
	"github.com/joho/godotenv"
	"github.com/rohmanseo/golang-clean-arch/exception"
	"os"
)

type IConfig interface {
	Get(key string) string
}
type confImpl struct {
}

func (config *confImpl) Get(key string) string {
	return os.Getenv(key)
}

func LoadConfig() IConfig {
	err := godotenv.Load()
	exception.PanicIfNeeded(err)
	return &confImpl{}
}
