package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ffelipelimao/bank/entities"
	"github.com/ffelipelimao/bank/usecases"
	"github.com/gofiber/fiber/v2"
)

var statusCodeErrHandle = map[error]int{
	entities.ErrInvalidDescription: http.StatusBadRequest,
	entities.ErrInvalidType:        http.StatusBadRequest,
	entities.ErrInvalidValue:       http.StatusBadRequest,
	usecases.ErrInsufficientFunds:  http.StatusUnprocessableEntity,
}

type SaveTransferUseCase interface {
	Execute(ctx context.Context, transfer *entities.Transfer) (*entities.Balance, error)
}

type SaveTransferController struct {
	saveTransferUseCase SaveTransferUseCase
}

func NewSaveTransferController(saveTransferUseCase SaveTransferUseCase) *SaveTransferController {
	return &SaveTransferController{
		saveTransferUseCase: saveTransferUseCase,
	}
}

func (stc *SaveTransferController) Handle(c *fiber.Ctx) error {
	ctx := c.Context()

	var transfer entities.Transfer

	if err := c.BodyParser(&transfer); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "invalid transfer body",
		})
	}

	intNumber, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "invalid client_id",
		})
	}

	transfer.UserID = intNumber

	output, err := stc.saveTransferUseCase.Execute(ctx, &transfer)
	if err != nil {
		if statusCode, ok := statusCodeErrHandle[err]; ok {
			return c.Status(statusCode).JSON(fiber.Map{
				"code":    statusCode,
				"message": err.Error(),
			})
		}

		if err.Error() == "user does not exists" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"code":    http.StatusNotFound,
				"message": err.Error(),
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(output)
}
