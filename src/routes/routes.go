package routes

import (
	"github.com/chmenegatti/k8s-study/src/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	nomes := app.Group("/nomes")
	nomes.Post("/", controller.CreateNome)
}
