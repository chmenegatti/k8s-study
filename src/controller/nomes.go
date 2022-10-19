package controller

import (
	"encoding/json"
	"fmt"

	"github.com/chmenegatti/k8s-study/src/database"
	"github.com/chmenegatti/k8s-study/src/models"
	responses "github.com/chmenegatti/k8s-study/src/responses/nomes"
	"github.com/chmenegatti/k8s-study/src/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateNome(c *fiber.Ctx) error {
	info := utils.Trace()
	db := database.GetDatabase()

	var nome models.Nomes

	if err := c.BodyParser(&nome); err != nil {
		utils.Logger("Erro", info, "Json mal formado")
		return c.Status(400).JSON(err.Error())
	}
	db.Create(&nome)
	response := responses.NomesResponse(nome)

	payload, err := json.Marshal(response)

	if err != nil {
		utils.Logger("Erro", info, "Json mal formado")
		return c.Status(400).JSON(err.Error())
	}

	utils.Logger("Info", info, "Payload"+fmt.Sprintf("%v", string(payload)))

	return c.Status(200).JSON(response)
}
