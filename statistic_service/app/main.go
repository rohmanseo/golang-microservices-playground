package main

import (
	"context"
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
	"github.com/rohmanseo/golang-clean-arch/repository/memory"
	"github.com/rohmanseo/golang-clean-arch/use_case"
)

func main() {
	configuration := config.LoadConfig()
	database := db.NewPostgreDbConnection(configuration)
	webFw := gin.Default()
	ctx := context.Background()
	redis := db.NewRedisDb(configuration)
	cacheDb := memory.NewCacheDataStore(redis, &ctx)
	natsConn := msg_broker.NewNats(configuration)
	repo := repository.NewStatisticRepository(&cacheDb, natsConn)
	msgUseCase := use_case.NewMsgUseCaseImpl(&repo)
	nats_subsciber.SetupSubscriber(natsConn, msgUseCase)
	useCase := use_case.NewStatisticUseCaseImpl(&repo)
	controller := user_controller.NewStatisticController(&useCase)

	router.SetupRouter(webFw, controller)

	defer func(database *sql.DB) {
		err := database.Close()
		exception.PanicIfNeeded(err)
	}(database)

	err := webFw.Run(configuration.Get("HOST_URL"))
	exception.PanicIfNeeded(err)

}
