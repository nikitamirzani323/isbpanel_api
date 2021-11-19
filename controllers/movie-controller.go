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

const Fieldmovie_home_redis = "LISTMOVIE_FRONTEND_ISBPANEL"

func Moviehome(c *fiber.Ctx) error {
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
			movie_video, _ := jsonparser.GetString(value, "movie_video")
			objchild.Movie_id = int(movie_id)
			objchild.Movie_type = movie_type
			objchild.Movie_title = movie_title
			objchild.Movie_label = movie_label
			objchild.Movie_thumbnail = movie_thumbnail
			objchild.Movie_video = movie_video
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
		helpers.SetRedis(Fieldmovie_home_redis, result, time.Minute*60)
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
