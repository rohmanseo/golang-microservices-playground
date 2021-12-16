package permanent

import (
	"database/sql"
	"fmt"
	"github.com/rohmanseo/golang-clean-arch/entity"
)

type permanentDbImpl struct {
	pgConnection *sql.DB
}

func (p *permanentDbImpl) CreateUser(user entity.User) (entity.User, error) {
	insertQuery := fmt.Sprintf("INSERT INTO \"user\"(name, email,password) VALUES ('%s', '%s','%s') RETURNING id,name,email;", user.Name, user.Email, user.Password)
	user = entity.User{
		Id:    0,
		Name:  "",
		Email: "",
	}
	err := p.pgConnection.QueryRow(insertQuery).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		fmt.Println(err)
		return user, err
	} else {
		return user, nil
	}
}

func (p *permanentDbImpl) GetUser(user entity.User) (entity.User, error) {
	query := fmt.Sprintf("SELECT id,name,email FROM public.\"user\" WHERE email='%s' AND password='%s' LIMIT 1;", user.Email, user.Password)
	user = entity.User{
		Id:    0,
		Name:  "",
		Email: "",
	}
	err := p.pgConnection.QueryRow(query).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func NewPermanentDb(db *sql.DB) IPermanentDb {
	return &permanentDbImpl{pgConnection: db}
}
