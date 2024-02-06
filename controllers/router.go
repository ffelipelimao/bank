package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	SaveTransfer func(c *fiber.Ctx) error
	GetExtract   func(c *fiber.Ctx) error
}

func (r *Router) Register(app *fiber.App) {
	group := app.Group("/clientes")

	group.Post("/:id/transacoes", r.SaveTransfer)
	group.Get("/:id/extrato", r.GetExtract)
}
