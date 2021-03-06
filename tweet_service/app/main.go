package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/rohmanseo/golang-clean-arch/config"
	user_controller "github.com/rohmanseo/golang-clean-arch/controller"
	"github.com/rohmanseo/golang-clean-arch/exception"
	"github.com/rohmanseo/golang-clean-arch/infrastructure/db"
	"github.com/rohmanseo/golang-clean-arch/infrastructure/msg_broker"
	"github.com/rohmanseo/golang-clean-arch/infrastructure/nats_subsciber"
	"github.com/rohmanseo/golang-clean-arch/infrastructure/router"
	"github.com/rohmanseo/golang-clean-arch/repository"
	"github.com/rohmanseo/golang-clean-arch/repository/permanent"
	"github.com/rohmanseo/golang-clean-arch/security"
	"github.com/rohmanseo/golang-clean-arch/use_case"
)

func main() {
	configuration := config.LoadConfig()
	database := db.NewPostgreDbConnection(configuration)
	webFw := gin.Default()
	permanentDb := permanent.NewPermanentDb(database)
	jwtMngr := security.NewJwtTokenManager(configuration)
	natsConn := msg_broker.NewNats(configuration)
	repo := repository.NewTweetRepository(&permanentDb)
	useCase := use_case.NewTweetUseCaseImpl(&repo)
	controller := user_controller.NewTweetController(&useCase, jwtMngr)
	msgUseCase := use_case.NewMsgUseCaseImpl(&repo)
	nats_subsciber.SetupSubscriber(natsConn, msgUseCase)
	router.SetupRouter(webFw, controller)

	defer func(database *sql.DB) {
		err := database.Close()
		exception.PanicIfNeeded(err)
	}(database)

	err := webFw.Run(configuration.Get("HOST_URL"))
	exception.PanicIfNeeded(err)

}
