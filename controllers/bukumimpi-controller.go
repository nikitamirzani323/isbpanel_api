package controllers

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_api/entities"
	"github.com/nikitamirzani323/isbpanel_api/helpers"
	"github.com/nikitamirzani323/isbpanel_api/models"
)

const Field_bukumimpihome_redis = "LISTBUKUMIMPI_FRONTEND_ISBPANEL"

type clienrequest struct {
	Tipe string `json:"tipe"`
	Nama string `json:"nama"`
}

func Bukumimpihome(c *fiber.Ctx) error {
	client := new(clienrequest)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	var obj entities.Model_bukumimpi
	var arraobj []entities.Model_bukumimpi
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_bukumimpihome_redis + "-" + client.Tipe + "-" + client.Nama)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		bukumimpi_type, _ := jsonparser.GetString(value, "bukumimpi_type")
		bukumimpi_name, _ := jsonparser.GetString(value, "bukumimpi_name")
		bukumimpi_nomor, _ := jsonparser.GetString(value, "bukumimpi_nomor")

		obj.Bukumimpi_type = bukumimpi_type
		obj.Bukumimpi_name = bukumimpi_name
		obj.Bukumimpi_nomor = bukumimpi_nomor
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_bukumimpiHome(client.Tipe, client.Nama)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_bukumimpihome_redis+"-"+client.Tipe+"-"+client.Nama, result, 1*time.Minute)
		log.Printf("BUKUMIMPI MYSQL %s - %s\n", client.Tipe, client.Nama)
		return c.JSON(result)
	} else {
		log.Printf("BUKUMIMPI CACHE %s - %s\n", client.Tipe, client.Nama)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
