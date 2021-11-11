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
const Field_tafsirmimpihome_redis = "LISTTAFSIRMIMPI_FRONTEND_ISBPANEL"

type clienrequest struct {
	Tipe string `json:"tipe"`
	Nama string `json:"nama"`
}
type clientafsirmimpirequest struct {
	Search string `json:"search"`
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
func TafsirMimpihome(c *fiber.Ctx) error {
	client := new(clientafsirmimpirequest)
	if err := c.BodyParser(client); err != nil {
		return err
	}
	var obj entities.Model_tafsirmimpi
	var arraobj []entities.Model_tafsirmimpi
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_tafsirmimpihome_redis + "-" + client.Search)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		tafsirmimpi_mimpi, _ := jsonparser.GetString(value, "tafsirmimpi_mimpi")
		tafsirmimpi_artimimpi, _ := jsonparser.GetString(value, "tafsirmimpi_artimimpi")
		tafsirmimpi_angka2d, _ := jsonparser.GetString(value, "tafsirmimpi_angka2d")
		tafsirmimpi_angka3d, _ := jsonparser.GetString(value, "tafsirmimpi_angka3d")
		tafsirmimpi_angka4d, _ := jsonparser.GetString(value, "tafsirmimpi_angka4d")

		obj.Tafsirmimpi_mimpi = tafsirmimpi_mimpi
		obj.Tafsirmimpi_artimimpi = tafsirmimpi_artimimpi
		obj.Tafsirmimpi_angka4d = tafsirmimpi_angka4d
		obj.Tafsirmimpi_angka3d = tafsirmimpi_angka3d
		obj.Tafsirmimpi_angka2d = tafsirmimpi_angka2d
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_tafsirmimpiHome(client.Search)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_tafsirmimpihome_redis+"-"+client.Search, result, 1*time.Minute)
		log.Printf("TAFSIR MIMPI MYSQL %s\n", client.Search)
		return c.JSON(result)
	} else {
		log.Printf("TAFSIR MIMPI CACHE %s\n", client.Search)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
