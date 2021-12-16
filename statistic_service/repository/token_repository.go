package repository

import (
	"github.com/rohmanseo/golang-clean-arch/entity"
)

type ITokenRepository interface {
	Create(user entity.User) (token entity.Token, err error)
	Remove(tkn entity.Token) (bool, error)
	Validate(token entity.Token) (bool, error)
}
