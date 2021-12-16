package repository

import (
	_ "github.com/lib/pq"
	"github.com/rohmanseo/golang-clean-arch/entity"
	"github.com/rohmanseo/golang-clean-arch/repository/permanent"
)

type UserRepositoryImpl struct {
	permanentDb permanent.IPermanentDb
}

func (u *UserRepositoryImpl) GetUser(user entity.User) (entity.User, error) {
	return u.permanentDb.GetUser(user)
}

func (u *UserRepositoryImpl) CreateUser(user entity.User) (entity.User, error) {
	return u.permanentDb.CreateUser(user)
}

func NewUserRepositoryImpl(db *permanent.IPermanentDb) IUserRepository {
	return &UserRepositoryImpl{
		permanentDb: *db,
	}
}
