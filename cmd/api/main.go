package main

import (
	"github.com/ffelipelimao/bank/internal/controllers"
	db "github.com/ffelipelimao/bank/internal/database"
	"github.com/ffelipelimao/bank/internal/repository"
	"github.com/ffelipelimao/bank/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := db.NewDatabase()
	defer db.Close()

	balanceRepository := repository.NewBalanceRepository(db)
	transferRepository := repository.NewTransferRepository(db)

	saveTransferUseCase := usecases.NewSaveTransfer(transferRepository, balanceRepository)
	saveTransferController := controllers.NewSaveTransferController(saveTransferUseCase)

	getExtractUseCase := usecases.NewGetExtract(transferRepository, balanceRepository)
	getExtractController := controllers.NewGetExtractController(getExtractUseCase)

	app := fiber.New()

	router := controllers.Router{
		SaveTransfer: saveTransferController.Handle,
		GetExtract:   getExtractController.Handle,
	}

	router.Register(app)

	app.Listen(":3000")
}
