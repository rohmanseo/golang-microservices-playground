package permanent

import "github.com/rohmanseo/golang-clean-arch/entity"

type IPermanentDb interface {
	GetUser(user entity.User) (entity.User, error)
	CreateUser(user entity.User) (entity.User, error)
}
