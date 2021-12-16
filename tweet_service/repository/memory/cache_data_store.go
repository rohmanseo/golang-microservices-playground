package memory

import (
	"github.com/rohmanseo/golang-clean-arch/entity"
)

type ICacheDataStore interface {
	AddToken(token entity.Token)
	RevokeToken(token entity.Token) (bool, error)
	ValidateToken(token entity.Token) bool
}
