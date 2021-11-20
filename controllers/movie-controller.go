package controllers

import (
	"log"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/isbpanel_api/entities"
	"github.com/nikitamirzani323/isbpanel_api/helpers"
	"github.com/nikitamirzani323/isbpanel_api/models"
)

const Fieldmovie_home_redis = "LISTMOVIE_FRONTEND_ISBPANEL"

func Moviehome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_clientmovie)
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
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	log.Println("Client TOKEN : ", temp_decp)
	log.Println("Client BODYPARSE : ", client.Client_hostname)
	flag_client := false
	switch temp_decp {
	case "167.86.112.29":
		flag_client = true
	case "localhost:7075":
		flag_client = true
	}
	if temp_decp != client.Client_hostname {
		flag_client = false
	}

	if !flag_client {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}
	var obj entities.Model_moviecategory
	var arraobj []entities.Model_moviecategory
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmovie_home_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_idcategory, _ := jsonparser.GetString(value, "movie_idcategory")
		movie_category, _ := jsonparser.GetString(value, "movie_category")
		movie_list, _, _, _ := jsonparser.Get(value, "movie_list")
		var objchild entities.Model_movie
		var arraobjchild []entities.Model_movie
		jsonparser.ArrayEach(movie_list, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			movie_id, _ := jsonparser.GetInt(value, "movie_id")
			movie_type, _ := jsonparser.GetString(value, "movie_type")
			movie_title, _ := jsonparser.GetString(value, "movie_title")
			movie_label, _ := jsonparser.GetString(value, "movie_label")
			movie_thumbnail, _ := jsonparser.GetString(value, "movie_thumbnail")
			movie_video, _, _, _ := jsonparser.Get(value, "movie_video")
			var objmoviesrc entities.Model_movievideo
			var arraobjmoviesrc []entities.Model_movievideo
			jsonparser.ArrayEach(movie_video, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				movie_src, _ := jsonparser.GetString(value, "movie_src")
				objmoviesrc.Movie_src = movie_src
				arraobjmoviesrc = append(arraobjmoviesrc, objmoviesrc)
			})
			objchild.Movie_id = int(movie_id)
			objchild.Movie_type = movie_type
			objchild.Movie_title = movie_title
			objchild.Movie_label = movie_label
			objchild.Movie_thumbnail = movie_thumbnail
			objchild.Movie_video = arraobjmoviesrc
			arraobjchild = append(arraobjchild, objchild)
		})

		obj.Movie_idcategory = movie_idcategory
		obj.Movie_category = movie_category
		obj.Movie_list = arraobjchild
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_movieHome()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmovie_home_redis, result, time.Minute*120)
		log.Println("MOVIE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
