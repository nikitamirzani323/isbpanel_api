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

const Field_home_redis = "LISTPASARAN_FRONTEND_ISBPANEL"

func Pasaranhome(c *fiber.Ctx) error {
	var obj entities.Model_pasaran
	var arraobj []entities.Model_pasaran
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		pasaran_id, _ := jsonparser.GetString(value, "pasaran_id")
		pasaran_name, _ := jsonparser.GetString(value, "pasaran_name")
		pasaran_url, _ := jsonparser.GetString(value, "pasaran_url")
		pasaran_diundi, _ := jsonparser.GetString(value, "pasaran_diundi")
		pasaran_jamjadwal, _ := jsonparser.GetString(value, "pasaran_jamjadwal")
		pasaran_datekeluaran, _ := jsonparser.GetString(value, "pasaran_datekeluaran")
		pasaran_keluaran, _ := jsonparser.GetString(value, "pasaran_keluaran")
		pasaran_dateprediksi, _ := jsonparser.GetString(value, "pasaran_dateprediksi")
		pasaran_bbfsprediksi, _ := jsonparser.GetString(value, "pasaran_bbfsprediksi")
		pasaran_nomorprediksi, _ := jsonparser.GetString(value, "pasaran_nomorprediksi")

		obj.Pasaran_id = pasaran_id
		obj.Pasaran_name = pasaran_name
		obj.Pasaran_url = pasaran_url
		obj.Pasaran_diundi = pasaran_diundi
		obj.Pasaran_jamjadwal = pasaran_jamjadwal
		obj.Pasaran_datekeluaran = pasaran_datekeluaran
		obj.Pasaran_keluaran = pasaran_keluaran
		obj.Pasaran_dateprediksi = pasaran_dateprediksi
		obj.Pasaran_bbfsprediksi = pasaran_bbfsprediksi
		obj.Pasaran_nomorprediksi = pasaran_nomorprediksi
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_pasaranHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_home_redis, result, 0)
		log.Println("PASARAN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("PASARAN CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
