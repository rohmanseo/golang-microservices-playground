package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rohmanseo/golang-clean-arch/config"
	"github.com/rohmanseo/golang-clean-arch/exception"
)

func NewPostgreDbConnection(config config.IConfig) *sql.DB {
	/*	ops := &pg.Options{
			Database: config.Get("DB_NAME"),
			User:     config.Get("DB_USERNAME"),
			Password: config.Get("DB_PASSWORD"),
		}
		db := pg.Connect(ops)*/

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.Get("DB_USERNAME"),
		config.Get("DB_PASSWORD"),
		config.Get("DB_NAME"))
	db, err := sql.Open("postgres", dbinfo)
	exception.PanicIfNeeded(err)
	return db
}
