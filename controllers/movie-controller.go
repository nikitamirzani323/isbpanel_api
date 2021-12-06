package controllers

import (
	"log"
	"strconv"
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
const Fieldseason_home_redis = "LISTSEASON_FRONTEND_ISBPANEL"
const Fieldepisode_home_redis = "LISTEPISODE_FRONTEND_ISBPANEL"

const Fieldmovie_mobile_redis = "LISTMOVIE-MOBILE"
const Fieldmoviegenre_mobile_redis = "LISTMOVIEGENRE-MOBILE"
const Fieldmoviedetail_mobile_redis = "LISTMOVIEDETAIL-MOBILE"
const Fieldfrontpagemovie_mobile_redis = "LISTFRONTPAGE-MOBILE"
const Fieldseason_mobile_redis = "LISTSEASON_MOBILE"
const Fieldepisode_mobile_redis = "LISTSEASONEPISODE_MOBILE"

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

	flag_client := models.Get_Domain(temp_decp)

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
func Movieseason(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_season)
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
	flag_client := models.Get_Domain(temp_decp)
	if !flag_client {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}

	var obj entities.Model_movieseason
	var arraobj []entities.Model_movieseason
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldseason_home_redis + "_" + strconv.Itoa(client.Movie_id))
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		season_id, _ := jsonparser.GetInt(value, "season_id")
		season_title, _ := jsonparser.GetString(value, "season_title")

		obj.Season_id = int(season_id)
		obj.Season_title = season_title
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.SeasonMovie(client.Movie_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldseason_home_redis+"_"+strconv.Itoa(client.Movie_id), result, time.Minute*1)
		log.Println("MOVIE SEASON MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE SEASON CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Movieepisode(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_episode)
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
	flag_client := models.Get_Domain(temp_decp)
	if !flag_client {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "NOT REGISTER",
			"record":  nil,
		})
	}

	var obj entities.Model_movieepisode
	var arraobj []entities.Model_movieepisode
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldepisode_home_redis + "_" + strconv.Itoa(client.Season_id))
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		episode_id, _ := jsonparser.GetInt(value, "episode_id")
		episode_title, _ := jsonparser.GetString(value, "episode_title")
		episode_src, _ := jsonparser.GetString(value, "episode_src")

		obj.Episode_id = int(episode_id)
		obj.Episode_title = episode_title
		obj.Episode_src = episode_src
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.EpisodeMovie(client.Season_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldepisode_home_redis+"_"+strconv.Itoa(client.Season_id), result, time.Minute*1)
		log.Println("MOVIE SEASON MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE SEASON CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}

//MOBILE
func Moviemobile(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_clientmobilemovie)
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

	var obj entities.Model_movielist
	var arraobj []entities.Model_movielist
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmovie_mobile_redis + "_" + client.Client_type)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_id, _ := jsonparser.GetInt(value, "movie_id")
		movie_type, _ := jsonparser.GetString(value, "movie_type")
		movie_title, _ := jsonparser.GetString(value, "movie_title")
		movie_label, _ := jsonparser.GetString(value, "movie_label")
		movie_descp, _ := jsonparser.GetString(value, "movie_descp")
		movie_year, _ := jsonparser.GetInt(value, "movie_year")
		movie_view, _ := jsonparser.GetInt(value, "movie_view")
		movie_img, _ := jsonparser.GetString(value, "movie_img")

		obj.Movie_id = int(movie_id)
		obj.Movie_type = movie_type
		obj.Movie_title = movie_title
		obj.Movie_label = movie_label
		obj.Movie_descp = movie_descp
		obj.Movie_year = int(movie_year)
		obj.Movie_view = int(movie_view)
		obj.Movie_img = movie_img
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_movielist(client.Client_type)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		if client.Client_type != "random" {
			helpers.SetRedis(Fieldmovie_mobile_redis+"_"+client.Client_type, result, time.Minute*120)
		} else {
			helpers.SetRedis(Fieldmovie_mobile_redis+"_"+client.Client_type, result, time.Minute*5)
		}
		log.Println("MOVIE MOBILE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE MOBILE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Moviedetailmobile(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_clientmobiledetailmobile)
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

	flag_updateview := models.Update_movieview(client.Client_username, client.Client_idmovie)
	log.Printf("Update View Mobile %d - %s - %t", client.Client_idmovie, client.Client_username, flag_updateview)

	var obj entities.Model_moviedetail
	var arraobj []entities.Model_moviedetail
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmoviedetail_mobile_redis + "_" + strconv.Itoa(client.Client_idmovie))
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_id, _ := jsonparser.GetInt(value, "movie_id")
		movie_type, _ := jsonparser.GetString(value, "movie_type")
		movie_title, _ := jsonparser.GetString(value, "movie_title")
		movie_label, _ := jsonparser.GetString(value, "movie_label")
		movie_descp, _ := jsonparser.GetString(value, "movie_descp")
		movie_year, _ := jsonparser.GetInt(value, "movie_year")
		movie_view, _ := jsonparser.GetInt(value, "movie_view")
		movie_src, _ := jsonparser.GetString(value, "movie_src")
		movie_img, _ := jsonparser.GetString(value, "movie_img")
		movie_genre, _ := jsonparser.GetString(value, "movie_genre")
		movie_totalsource, _ := jsonparser.GetInt(value, "movie_totalsource")

		movie_video, _, _, _ := jsonparser.Get(value, "movie_video")
		var objmoviesrc entities.Model_movievideo
		var arraobjmoviesrc []entities.Model_movievideo
		jsonparser.ArrayEach(movie_video, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			movie_src, _ := jsonparser.GetString(value, "movie_src")
			objmoviesrc.Movie_src = movie_src
			arraobjmoviesrc = append(arraobjmoviesrc, objmoviesrc)
		})

		obj.Movie_id = int(movie_id)
		obj.Movie_type = movie_type
		obj.Movie_title = movie_title
		obj.Movie_label = movie_label
		obj.Movie_descp = movie_descp
		obj.Movie_year = int(movie_year)
		obj.Movie_view = int(movie_view)
		obj.Movie_genre = movie_genre
		obj.Movie_img = movie_img
		obj.Movie_src = movie_src
		obj.Movie_video = arraobjmoviesrc
		obj.Movie_totalsource = int(movie_totalsource)
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_moviedetail(client.Client_idmovie)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmoviedetail_mobile_redis+"_"+strconv.Itoa(client.Client_idmovie), result, time.Minute*120)
		log.Println("MOVIE MOBILE DETAIL MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE MOBILE DETAIL CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Moviegenremobile(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_clientmobilegenremovie)
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

	var obj entities.Model_movielist
	var arraobj []entities.Model_movielist
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldmoviegenre_mobile_redis + "_" + strconv.Itoa(client.Client_genre))
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_id, _ := jsonparser.GetInt(value, "movie_id")
		movie_type, _ := jsonparser.GetString(value, "movie_type")
		movie_title, _ := jsonparser.GetString(value, "movie_title")
		movie_label, _ := jsonparser.GetString(value, "movie_label")
		movie_descp, _ := jsonparser.GetString(value, "movie_descp")
		movie_year, _ := jsonparser.GetInt(value, "movie_year")
		movie_view, _ := jsonparser.GetInt(value, "movie_view")
		movie_img, _ := jsonparser.GetString(value, "movie_img")

		obj.Movie_id = int(movie_id)
		obj.Movie_type = movie_type
		obj.Movie_title = movie_title
		obj.Movie_label = movie_label
		obj.Movie_descp = movie_descp
		obj.Movie_year = int(movie_year)
		obj.Movie_view = int(movie_view)
		obj.Movie_img = movie_img
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_moviegenre(client.Client_genre)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldmoviegenre_mobile_redis+"_"+strconv.Itoa(client.Client_genre), result, time.Minute*120)
		log.Println("MOVIE GENRE MOBILE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE MOBILE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Moviefrontpagemobile(c *fiber.Ctx) error {
	var obj entities.Model_mobilemoviecategory
	var arraobj []entities.Model_mobilemoviecategory
	var objslider entities.Model_movie
	var arraobjslider []entities.Model_movie
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldfrontpagemovie_mobile_redis)
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	recordslider_RD, _, _, _ := jsonparser.Get(jsonredis, "slider")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "genre")
	jsonparser.ArrayEach(recordslider_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
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
		objslider.Movie_id = int(movie_id)
		objslider.Movie_type = movie_type
		objslider.Movie_title = movie_title
		objslider.Movie_label = movie_label
		objslider.Movie_thumbnail = movie_thumbnail
		objslider.Movie_video = arraobjmoviesrc
		arraobjslider = append(arraobjslider, objslider)
	})
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		movie_idcategory, _ := jsonparser.GetInt(value, "movie_idcategory")
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

		obj.Movie_idcategory = int(movie_idcategory)
		obj.Movie_category = movie_category
		obj.Movie_list = arraobjchild
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.Fetch_frontpage()
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldfrontpagemovie_mobile_redis, result, time.Minute*120)
		log.Println("MOVIE FRONTPAGE MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE FRONTPAGE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"slider":  arraobjslider,
			"genre":   arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Moviemobileseason(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_mobileseason)
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

	var obj entities.Model_movieseason
	var arraobj []entities.Model_movieseason
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldseason_mobile_redis + "_" + strconv.Itoa(client.Movie_id))
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		season_id, _ := jsonparser.GetInt(value, "season_id")
		season_title, _ := jsonparser.GetString(value, "season_title")

		obj.Season_id = int(season_id)
		obj.Season_title = season_title
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.SeasonMovie(client.Movie_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldseason_mobile_redis+"_"+strconv.Itoa(client.Movie_id), result, time.Minute*30)
		log.Println("MOVIE SEASON MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE SEASON CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Moviemobileepisode(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_mobileepisode)
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

	var obj entities.Model_movieepisode
	var arraobj []entities.Model_movieepisode
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldepisode_mobile_redis + "_" + strconv.Itoa(client.Season_id))
	jsonredis := []byte(resultredis)
	message_RD, _ := jsonparser.GetString(jsonredis, "message")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		episode_id, _ := jsonparser.GetInt(value, "episode_id")
		episode_title, _ := jsonparser.GetString(value, "episode_title")
		episode_src, _ := jsonparser.GetString(value, "episode_src")

		obj.Episode_id = int(episode_id)
		obj.Episode_title = episode_title
		obj.Episode_src = episode_src
		arraobj = append(arraobj, obj)
	})
	if !flag {
		result, err := models.EpisodeMovie(client.Season_id)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldepisode_mobile_redis+"_"+strconv.Itoa(client.Season_id), result, time.Minute*30)
		log.Println("MOVIE SEASON MYSQL")
		return c.JSON(result)
	} else {
		log.Println("MOVIE SEASON CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": message_RD,
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
