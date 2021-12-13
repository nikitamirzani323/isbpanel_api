package controllers

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_api/config"
	"github.com/nikitamirzani323/isbpanel_api/entities"
	"github.com/nikitamirzani323/isbpanel_api/helpers"
	"github.com/nikitamirzani323/isbpanel_api/models"
)

func CheckLoginmobile(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Loginmobile)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result := models.Loginmobile_Model(client.Username)
	version := models.Mobileversion_Model()

	if result {
		flag_point := models.Save_moviepoint(client.Username, "POINT_LOGIN", 0, config.POINT_LOGIN)
		log.Printf("POINT_LOGIN STATUS : %t", flag_point)
		dataclient := client.Username + "==" + client.Name + "==" + client.Device
		dataclient_encr, keymap := helpers.Encryption(dataclient)
		dataclient_encr_final := dataclient_encr + "|" + strconv.Itoa(keymap)
		t, err := helpers.GenerateNewAccessToken(dataclient_encr_final)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"status":  "Y",
			"version": version,
			"message": "updated",
			"token":   t,
		})
	} else {
		return c.JSON(fiber.Map{
			"status":  "N",
			"version": version,
			"message": "",
		})
	}
}
