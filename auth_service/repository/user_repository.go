package repository

import (
	"github.com/rohmanseo/golang-clean-arch/entity"
)

type IUserRepository interface {
	GetUser(user entity.User) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
}
