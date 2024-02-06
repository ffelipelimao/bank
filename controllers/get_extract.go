package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ffelipelimao/bank/entities"
	"github.com/gofiber/fiber/v2"
)

type GetExtractUseCase interface {
	Execute(ctx context.Context, userID *int64) (*entities.Extract, error)
}

type GetExtractController struct {
	getExtractUseCase GetExtractUseCase
}

func NewGetExtractController(getExtractUseCase GetExtractUseCase) *GetExtractController {
	return &GetExtractController{
		getExtractUseCase: getExtractUseCase,
	}
}

func (gx *GetExtractController) Handle(c *fiber.Ctx) error {
	ctx := c.Context()

	intNumber, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "invalid client_id",
		})
	}

	output, err := gx.getExtractUseCase.Execute(ctx, &intNumber)
	if err != nil {
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
