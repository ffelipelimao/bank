package main

import (
	"github.com/ffelipelimao/bank/internal/controllers"
	db "github.com/ffelipelimao/bank/internal/database"
	"github.com/ffelipelimao/bank/internal/repository"
	"github.com/ffelipelimao/bank/internal/uow"
	"github.com/ffelipelimao/bank/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := db.NewDatabase()
	defer db.Close()

	uow := uow.NewUnityOfWork(db)

	balanceRepository := repository.NewBalanceRepository()
	transferRepository := repository.NewTransferRepository()

	saveTransferUseCase := usecases.NewSaveTransfer(transferRepository, balanceRepository, uow)
	saveTransferController := controllers.NewSaveTransferController(saveTransferUseCase)

	getExtractUseCase := usecases.NewGetExtract(transferRepository, balanceRepository, uow)
	getExtractController := controllers.NewGetExtractController(getExtractUseCase)

	app := fiber.New()

	router := controllers.Router{
		SaveTransfer: saveTransferController.Handle,
		GetExtract:   getExtractController.Handle,
	}

	router.Register(app)

	app.Listen(":3000")
}
