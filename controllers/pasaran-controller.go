package controllers

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_api/entities"
	"github.com/nikitamirzani323/isbpanel_api/helpers"
	"github.com/nikitamirzani323/isbpanel_api/models"
)

const Field_home_redis = "LISTPASARAN_FRONTEND_ISBPANEL"
const Field_keluaran_redis = "LISTKELUARAN_FRONTEND_ISBPANEL"

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
func Keluaranhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_keluaran)
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

	var obj entities.Model_keluaran
	var arraobj []entities.Model_keluaran
	var obj_minggu entities.Model_keluaranpaitominggu
	var arraobj_minggu []entities.Model_keluaranpaitominggu
	var obj_senin entities.Model_keluaranpaitosenin
	var arraobj_senin []entities.Model_keluaranpaitosenin
	var obj_selasa entities.Model_keluaranpaitoselasa
	var arraobj_selasa []entities.Model_keluaranpaitoselasa
	var obj_rabu entities.Model_keluaranpaitorabu
	var arraobj_rabu []entities.Model_keluaranpaitorabu
	var obj_kamis entities.Model_keluaranpaitokamis
	var arraobj_kamis []entities.Model_keluaranpaitokamis
	var obj_jumat entities.Model_keluaranpaitojumat
	var arraobj_jumat []entities.Model_keluaranpaitojumat
	var obj_sabtu entities.Model_keluaranpaitosabtu
	var arraobj_sabtu []entities.Model_keluaranpaitosabtu
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Field_keluaran_redis + "_" + client.Pasaran_id)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	paito_minggu_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_minggu")
	paito_senin_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_senin")
	paito_selasa_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_selasa")
	paito_rabu_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_rabu")
	paito_kamis_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_kamis")
	paito_jumat_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_jumat")
	paito_sabtu_RD, _, _, _ := jsonparser.Get(jsonredis, "paito_sabtu")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_datekeluaran, _ := jsonparser.GetString(value, "keluaran_datekeluaran")
		keluaran_periode, _ := jsonparser.GetString(value, "keluaran_periode")
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")

		obj.Keluaran_datekeluaran = keluaran_datekeluaran
		obj.Keluaran_periode = keluaran_periode
		obj.Keluaran_nomor = keluaran_nomor
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(paito_minggu_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_minggu.Keluaran_nomor = keluaran_nomor
		arraobj_minggu = append(arraobj_minggu, obj_minggu)
	})
	jsonparser.ArrayEach(paito_senin_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_senin.Keluaran_nomor = keluaran_nomor
		arraobj_senin = append(arraobj_senin, obj_senin)
	})
	jsonparser.ArrayEach(paito_selasa_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_selasa.Keluaran_nomor = keluaran_nomor
		arraobj_selasa = append(arraobj_selasa, obj_selasa)
	})
	jsonparser.ArrayEach(paito_rabu_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_rabu.Keluaran_nomor = keluaran_nomor
		arraobj_rabu = append(arraobj_rabu, obj_rabu)
	})
	jsonparser.ArrayEach(paito_kamis_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_kamis.Keluaran_nomor = keluaran_nomor
		arraobj_kamis = append(arraobj_kamis, obj_kamis)
	})
	jsonparser.ArrayEach(paito_jumat_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_jumat.Keluaran_nomor = keluaran_nomor
		arraobj_jumat = append(arraobj_jumat, obj_jumat)
	})
	jsonparser.ArrayEach(paito_sabtu_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		keluaran_nomor, _ := jsonparser.GetString(value, "keluaran_nomor")
		obj_sabtu.Keluaran_nomor = keluaran_nomor
		arraobj_sabtu = append(arraobj_sabtu, obj_sabtu)
	})
	if !flag {
		result, err := models.Fetch_keluaran(client.Pasaran_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Field_keluaran_redis+"_"+client.Pasaran_id, result, 0)
		log.Println("KELUARAN MYSQL")
		return c.JSON(result)
	} else {
		log.Println("KELUARAN CACHE")
		return c.JSON(fiber.Map{
			"status":       fiber.StatusOK,
			"message":      message_RD,
			"record":       arraobj,
			"paito_minggu": arraobj_minggu,
			"paito_senin":  arraobj_senin,
			"paito_selasa": arraobj_selasa,
			"paito_rabu":   arraobj_rabu,
			"paito_kamis":  arraobj_kamis,
			"paito_jumat":  arraobj_jumat,
			"paito_sabtu":  arraobj_sabtu,
			"time":         time.Since(render_page).String(),
		})
	}
}
