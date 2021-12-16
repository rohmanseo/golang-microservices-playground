package repository

import (
	_ "github.com/lib/pq"
	"github.com/rohmanseo/golang-clean-arch/entity"
	"github.com/rohmanseo/golang-clean-arch/repository/memory"
	"github.com/rohmanseo/golang-clean-arch/security"
)

type TokenRepositoryImpl struct {
	jwtManager security.IJwtToken
	cacheDb    memory.ICacheDataStore
}

func (t *TokenRepositoryImpl) Remove(tkn entity.Token) (bool, error) {
	return t.cacheDb.RevokeToken(tkn)
}

func (t *TokenRepositoryImpl) Create(user entity.User) (entity.Token, error) {
	token, err := t.jwtManager.GenerateToken(user)
	t.cacheDb.AddToken(token)
	return token, err
}

func (t *TokenRepositoryImpl) Validate(token entity.Token) (bool, error) {
	return t.cacheDb.ValidateToken(token), nil
}

func NewTokenRepository(jwtMngr security.IJwtToken, cache memory.ICacheDataStore) ITokenRepository {
	return &TokenRepositoryImpl{
		jwtManager: jwtMngr,
		cacheDb:    cache,
	}
}
