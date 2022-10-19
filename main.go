package main

import (
	"log"

	"github.com/chmenegatti/k8s-study/src/database"
	"github.com/chmenegatti/k8s-study/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDb()

	routes.Routes(app)

	log.Fatal(app.Listen(":3001"))
}
