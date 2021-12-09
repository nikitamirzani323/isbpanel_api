package models

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_api/config"
	"github.com/nikitamirzani323/isbpanel_api/db"
	"github.com/nikitamirzani323/isbpanel_api/entities"
	"github.com/nikitamirzani323/isbpanel_api/helpers"
	"github.com/nleeper/goment"
)

func Fetch_movieHome() (helpers.Response, error) {
	var obj entities.Model_moviecategory
	var arraobj []entities.Model_moviecategory
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_toprated := `SELECT 
		SUM(A.ratingposter) as total, 
		B.movieid, B.movietitle , B.movietype, B.label, COALESCE(B.posted_id,0) , B.urlthumbnail  
		FROM ` + config.DB_tbl_trx_rate + ` as A 
		JOIN ` + config.DB_tbl_trx_movie + ` as B ON B.movieid = A.idposter 
		WHERE B.enabled = 1 
		GROUP BY A.idposter 
		ORDER BY total DESC LIMIT 24     
	`
	row_toprated, err_toprated := con.QueryContext(ctx, sql_toprated)
	helpers.ErrorCheck(err_toprated)
	var objtoprated entities.Model_movie
	var arratoprated []entities.Model_movie
	for row_toprated.Next() {
		var (
			movieid_db, posted_id_db, total_db                     int
			movietitle_db, movietype_db, label_db, urlthumbnail_db string
		)

		err := row_toprated.Scan(&total_db, &movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db, &urlthumbnail_db)
		helpers.ErrorCheck(err)
		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}

		movie_url, _, _ := _GetVideo(movieid_db, "")

		objtoprated.Movie_id = movieid_db
		objtoprated.Movie_type = movietype_db
		objtoprated.Movie_title = movietitle_db
		objtoprated.Movie_label = label_db
		objtoprated.Movie_thumbnail = path_image
		objtoprated.Movie_video = movie_url
		arratoprated = append(arratoprated, objtoprated)
		msg = "Success"
	}
	defer row_toprated.Close()
	obj.Movie_idcategory = "toprated"
	obj.Movie_category = "Top Rated"
	obj.Movie_list = arratoprated
	arraobj = append(arraobj, obj)

	sql_selectview := `SELECT 
		movieid, movietitle , movietype, label, COALESCE(posted_id,0) , urlthumbnail    
		FROM ` + config.DB_tbl_trx_movie + ` 
		WHERE enabled = 1 
		ORDER BY views DESC LIMIT 24     
	`
	row, err := con.QueryContext(ctx, sql_selectview)
	helpers.ErrorCheck(err)
	var objpopular entities.Model_movie
	var arraobjpopular []entities.Model_movie
	for row.Next() {
		var (
			movieid_db, posted_id_db                               int
			movietitle_db, movietype_db, label_db, urlthumbnail_db string
		)

		err = row.Scan(&movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db, &urlthumbnail_db)
		helpers.ErrorCheck(err)
		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}

		movie_url, _, _ := _GetVideo(movieid_db, "")

		objpopular.Movie_id = movieid_db
		objpopular.Movie_type = movietype_db
		objpopular.Movie_title = movietitle_db
		objpopular.Movie_label = label_db
		objpopular.Movie_thumbnail = path_image
		objpopular.Movie_video = movie_url
		arraobjpopular = append(arraobjpopular, objpopular)
		msg = "Success"
	}
	defer row.Close()
	obj.Movie_idcategory = "popular"
	obj.Movie_category = "Popular"
	obj.Movie_list = arraobjpopular
	arraobj = append(arraobj, obj)

	sql_selectnew := `SELECT 
		movieid, movietitle , movietype, label, COALESCE(posted_id,0) , urlthumbnail 
		FROM ` + config.DB_tbl_trx_movie + ` 
		WHERE enabled = 1 
		ORDER BY createdatemovie DESC LIMIT 24	     
	`
	var objchildbaru entities.Model_movie
	var arraobjchildbaru []entities.Model_movie
	row_new, err_new := con.QueryContext(ctx, sql_selectnew)
	helpers.ErrorCheck(err_new)
	for row_new.Next() {
		var (
			movieid_db, posted_id_db                               int
			movietitle_db, movietype_db, label_db, urlthumbnail_db string
		)

		err = row_new.Scan(&movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db, &urlthumbnail_db)
		helpers.ErrorCheck(err)
		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}
		movie_url, _, _ := _GetVideo(movieid_db, "")

		objchildbaru.Movie_id = movieid_db
		objchildbaru.Movie_type = movietype_db
		objchildbaru.Movie_title = movietitle_db
		objchildbaru.Movie_label = label_db
		objchildbaru.Movie_thumbnail = path_image
		objchildbaru.Movie_video = movie_url
		arraobjchildbaru = append(arraobjchildbaru, objchildbaru)
		msg = "Success"
	}
	defer row_new.Close()
	obj.Movie_idcategory = "new"
	obj.Movie_category = "Terbaru"
	obj.Movie_list = arraobjchildbaru
	arraobj = append(arraobj, obj)

	sql_selectrekomendasi := `SELECT 
		movieid, movietitle , movietype, label, COALESCE(posted_id,0) , urlthumbnail    
		FROM ` + config.DB_tbl_trx_movie + ` 
		WHERE enabled = 1 
		ORDER BY RAND() LIMIT 72       
	`
	var objrekomendasi entities.Model_movie
	var arraobjrekomendasi []entities.Model_movie
	row_rekomendasi, err_rekomendasi := con.QueryContext(ctx, sql_selectrekomendasi)
	helpers.ErrorCheck(err_rekomendasi)
	for row_rekomendasi.Next() {
		var (
			movieid_db, posted_id_db                               int
			movietitle_db, movietype_db, label_db, urlthumbnail_db string
		)

		err = row_rekomendasi.Scan(&movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db, &urlthumbnail_db)
		helpers.ErrorCheck(err)
		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}

		movie_url, _, _ := _GetVideo(movieid_db, "")

		objrekomendasi.Movie_id = movieid_db
		objrekomendasi.Movie_type = movietype_db
		objrekomendasi.Movie_title = movietitle_db
		objrekomendasi.Movie_label = label_db
		objrekomendasi.Movie_thumbnail = path_image
		objrekomendasi.Movie_video = movie_url
		arraobjrekomendasi = append(arraobjrekomendasi, objrekomendasi)
		msg = "Success"
	}
	defer row_new.Close()
	obj.Movie_idcategory = "rekomendasi"
	obj.Movie_category = "Rekomendasi"
	obj.Movie_list = arraobjrekomendasi
	arraobj = append(arraobj, obj)

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func SeasonMovie(idmovie int) (helpers.Response, error) {
	var obj entities.Model_movieseason
	var arraobj []entities.Model_movieseason
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_season := `SELECT 
		id, title   
		FROM ` + config.DB_tbl_mst_series_season + ` 
		WHERE poster_id=?  
		ORDER BY position ASC      
	`
	row, err := con.QueryContext(ctx, sql_season, idmovie)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			id_db    int
			title_db string
		)

		err = row.Scan(&id_db, &title_db)
		helpers.ErrorCheck(err)

		obj.Season_id = id_db
		obj.Season_title = title_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func EpisodeMovie(idseason int) (helpers.Response, error) {
	var obj entities.Model_movieepisode
	var arraobj []entities.Model_movieepisode
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_season := `SELECT 
		id, title   
		FROM ` + config.DB_tbl_mst_series_episode + ` 
		WHERE season_id=?  
		ORDER BY position ASC      
	`
	row, err := con.QueryContext(ctx, sql_season, idseason)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			id_db    int
			title_db string
		)

		err = row.Scan(&id_db, &title_db)
		helpers.ErrorCheck(err)

		obj.Episode_id = id_db
		obj.Episode_title = title_db
		obj.Episode_src = _GetVideoEpisode(id_db)
		arraobj = append(arraobj, obj)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}

//MOBILE
func Fetch_movielist(tipe, username string) (helpers.Response, error) {
	var obj entities.Model_movielist
	var arraobj []entities.Model_movielist
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""

	if tipe == "movie" {
		sql_select += "SELECT "
		sql_select += "movieid, COALESCE(posted_id,0), movietitle, movietype, year, views, urlthumbnail, label, description "
		sql_select += "FROM " + config.DB_VIEW_MOVIE + " "
		sql_select += "ORDER BY RAND()  LIMIT 300 "
	} else if tipe == "serie" {
		sql_select += "SELECT "
		sql_select += "movieid, COALESCE(posted_id,0), movietitle, movietype, year, views, urlthumbnail, label, description "
		sql_select += "FROM " + config.DB_VIEW_MOVIESERIES + " "
		sql_select += "ORDER BY RAND()  LIMIT 300 "
	} else if tipe == "favorite" {
		sql_select += "SELECT "
		sql_select += "idposter as movieid, COALESCE(posted_id,0), movietitle, movietype, year, views, urlthumbnail, label, description "
		sql_select += "FROM " + config.DB_VIEW_MOVIEFAVORITE + " "
		sql_select += "WHERE username='" + username + "' "
		sql_select += "ORDER BY createfavorite DESC LIMIT 300 "
		log.Println(sql_select)
	} else {
		sql_select += "SELECT "
		sql_select += "movieid, COALESCE(posted_id,0), movietitle, movietype, year, views, urlthumbnail, label, description "
		sql_select += "FROM " + config.DB_tbl_trx_movie + " "
		sql_select += "WHERE enabled='1' "
		sql_select += "ORDER BY RAND()  LIMIT 30 "
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			movieid_db, posted_id_db, year_db, views_db                            int
			movietitle_db, movietype_db, label_db, urlthumbnail_db, description_db string
		)

		err := row.Scan(&movieid_db, &posted_id_db, &movietitle_db, &movietype_db, &year_db, &views_db,
			&urlthumbnail_db, &label_db, &description_db)
		helpers.ErrorCheck(err)
		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}

		obj.Movie_id = movieid_db
		obj.Movie_type = movietype_db
		obj.Movie_title = movietitle_db
		obj.Movie_label = label_db
		obj.Movie_descp = description_db
		obj.Movie_year = year_db
		obj.Movie_view = views_db
		obj.Movie_img = path_image
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_moviegenre(genre int) (helpers.Response, error) {
	var obj entities.Model_movielist
	var arraobj []entities.Model_movielist
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	if genre > 0 {
		sql_select += " "
		sql_select += "SELECT "
		sql_select += "A.movieid, COALESCE(A.posted_id,0), A.movietitle, A.movietype, A.year, A.views, A.urlthumbnail, A.label, A.description "
		sql_select += "FROM " + config.DB_tbl_trx_movie + " as A "
		sql_select += "JOIN " + config.DB_tbl_trx_moviegenre + " as B on B.movieid = A.movieid  "
		sql_select += "JOIN " + config.DB_tbl_mst_movie_genre + " as C on C.idgenre = B.idgenre  "
		sql_select += "WHERE C.idgenre=? "
		sql_select += "ORDER BY RAND() DESC LIMIT 300 "
		row, err := con.QueryContext(ctx, sql_select, genre)
		helpers.ErrorCheck(err)

		for row.Next() {
			var (
				movieid_db, posted_id_db, year_db, views_db                            int
				movietitle_db, movietype_db, label_db, urlthumbnail_db, description_db string
			)

			err := row.Scan(&movieid_db, &posted_id_db, &movietitle_db, &movietype_db, &year_db, &views_db,
				&urlthumbnail_db, &label_db, &description_db)
			helpers.ErrorCheck(err)
			path_image := ""
			if urlthumbnail_db == "" {
				poster_image, poster_extension := _GetMedia(posted_id_db)
				path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
			} else {
				path_image = urlthumbnail_db
			}

			obj.Movie_id = movieid_db
			obj.Movie_type = movietype_db
			obj.Movie_title = movietitle_db
			obj.Movie_label = label_db
			obj.Movie_descp = description_db
			obj.Movie_year = year_db
			obj.Movie_view = views_db
			obj.Movie_img = path_image
			arraobj = append(arraobj, obj)
			msg = "Success"
		}
		defer row.Close()
	}
	if genre < 0 {
		sql_select += " "
		sql_select += "SELECT "
		sql_select += "SUM(A.ratingposter) as total, "
		sql_select += "B.movieid, COALESCE(B.posted_id,0), B.movietitle, B.movietype, B.year, B.views, B.urlthumbnail, B.label, B.description "
		sql_select += "FROM " + config.DB_tbl_trx_rate + " as A "
		sql_select += "JOIN " + config.DB_tbl_trx_movie + " as B ON B.movieid = A.idposter   "
		sql_select += "WHERE B.enabled = 1 "
		sql_select += "GROUP BY A.idposter "
		sql_select += "ORDER BY total DESC LIMIT 300 "
		row, err := con.QueryContext(ctx, sql_select)
		helpers.ErrorCheck(err)

		for row.Next() {
			var (
				movieid_db, posted_id_db, year_db, views_db, total_db                  int
				movietitle_db, movietype_db, label_db, urlthumbnail_db, description_db string
			)

			err := row.Scan(&total_db, &movieid_db, &posted_id_db, &movietitle_db, &movietype_db, &year_db, &views_db,
				&urlthumbnail_db, &label_db, &description_db)
			helpers.ErrorCheck(err)
			path_image := ""
			if urlthumbnail_db == "" {
				poster_image, poster_extension := _GetMedia(posted_id_db)
				path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
			} else {
				path_image = urlthumbnail_db
			}

			obj.Movie_id = movieid_db
			obj.Movie_type = movietype_db
			obj.Movie_title = movietitle_db
			obj.Movie_label = label_db
			obj.Movie_descp = description_db
			obj.Movie_year = year_db
			obj.Movie_view = views_db
			obj.Movie_img = path_image
			arraobj = append(arraobj, obj)
			msg = "Success"
		}
		defer row.Close()
	}
	if genre == 0 {
		sql_select += " "
		sql_select += "SELECT "
		sql_select += "movieid, COALESCE(posted_id,0), movietitle, movietype, year, views, urlthumbnail, label, description "
		sql_select += "FROM " + config.DB_tbl_trx_movie + "  "
		sql_select += "ORDER BY views DESC LIMIT 300 "

		row, err := con.QueryContext(ctx, sql_select)
		helpers.ErrorCheck(err)

		for row.Next() {
			var (
				movieid_db, posted_id_db, year_db, views_db                            int
				movietitle_db, movietype_db, label_db, urlthumbnail_db, description_db string
			)

			err := row.Scan(&movieid_db, &posted_id_db, &movietitle_db, &movietype_db, &year_db, &views_db,
				&urlthumbnail_db, &label_db, &description_db)
			helpers.ErrorCheck(err)
			path_image := ""
			if urlthumbnail_db == "" {
				poster_image, poster_extension := _GetMedia(posted_id_db)
				path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
			} else {
				path_image = urlthumbnail_db
			}

			obj.Movie_id = movieid_db
			obj.Movie_type = movietype_db
			obj.Movie_title = movietitle_db
			obj.Movie_label = label_db
			obj.Movie_descp = description_db
			obj.Movie_year = year_db
			obj.Movie_view = views_db
			obj.Movie_img = path_image
			arraobj = append(arraobj, obj)
			msg = "Success"
		}
		defer row.Close()
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_moviedetail(movieid int, username string) (helpers.Response, error) {
	var obj entities.Model_moviedetail
	var arraobj []entities.Model_moviedetail
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += "SELECT "
	sql_select += "movieid, COALESCE(posted_id,0), movietitle, movietype, year, views, urlthumbnail, label, description "
	sql_select += "FROM " + config.DB_tbl_trx_movie + " "
	sql_select += "WHERE enabled='1' "
	sql_select += "AND movieid=? "

	row, err := con.QueryContext(ctx, sql_select, movieid)
	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			movieid_db, posted_id_db, year_db, views_db                            int
			movietitle_db, movietype_db, label_db, urlthumbnail_db, description_db string
		)

		err := row.Scan(&movieid_db, &posted_id_db, &movietitle_db, &movietype_db, &year_db, &views_db,
			&urlthumbnail_db, &label_db, &description_db)
		helpers.ErrorCheck(err)
		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}
		movie_url, totalsource, _ := _GetVideo(movieid_db, "")
		_, _, movie_src := _GetVideo(movieid_db, "single")
		genre := ""
		//GENRE
		sql_genre := `SELECT 
			B.nmgenre   
			FROM ` + config.DB_tbl_trx_moviegenre + ` as A 
			JOIN ` + config.DB_tbl_mst_movie_genre + ` as B ON B.idgenre = A.idgenre 
			WHERE A.movieid = ?    
		`
		row_genre, err_genre := con.QueryContext(ctx, sql_genre, movieid_db)
		helpers.ErrorCheck(err_genre)
		for row_genre.Next() {
			var nmgenre_db string

			err := row_genre.Scan(&nmgenre_db)
			helpers.ErrorCheck(err)
			genre += nmgenre_db + ","
		}

		obj.Movie_id = movieid_db
		obj.Movie_type = movietype_db
		obj.Movie_title = movietitle_db
		obj.Movie_label = label_db
		obj.Movie_descp = description_db
		obj.Movie_year = year_db
		obj.Movie_view = views_db
		obj.Movie_img = path_image
		obj.Movie_genre = genre
		obj.Movie_src = movie_src
		obj.Movie_favorite = _GetFavorite(movieid_db, username)
		obj.Movie_totalsource = totalsource
		obj.Movie_video = movie_url
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_frontpage() (helpers.Responsemobilemovie, error) {
	var obj entities.Model_mobilemoviecategory
	var arraobj []entities.Model_mobilemoviecategory
	var res helpers.Responsemobilemovie
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_toprated := `SELECT 
		SUM(A.ratingposter) as total, 
		B.movieid, B.movietitle , B.movietype, B.label, COALESCE(B.posted_id,0) , B.urlthumbnail  
		FROM ` + config.DB_tbl_trx_rate + ` as A 
		JOIN ` + config.DB_tbl_trx_movie + ` as B ON B.movieid = A.idposter 
		WHERE B.enabled = 1 
		GROUP BY A.idposter 
		ORDER BY total DESC LIMIT 10     
	`
	row_toprated, err_toprated := con.QueryContext(ctx, sql_toprated)
	helpers.ErrorCheck(err_toprated)
	var objtoprated entities.Model_movie
	var arratoprated []entities.Model_movie
	for row_toprated.Next() {
		var (
			movieid_db, posted_id_db, total_db                     int
			movietitle_db, movietype_db, label_db, urlthumbnail_db string
		)

		err := row_toprated.Scan(&total_db, &movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db, &urlthumbnail_db)
		helpers.ErrorCheck(err)
		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}

		objtoprated.Movie_id = movieid_db
		objtoprated.Movie_type = movietype_db
		objtoprated.Movie_title = movietitle_db
		objtoprated.Movie_label = label_db
		objtoprated.Movie_thumbnail = path_image
		objtoprated.Movie_video = ""
		arratoprated = append(arratoprated, objtoprated)
		msg = "Success"
	}
	defer row_toprated.Close()
	obj.Movie_idcategory = -1
	obj.Movie_category = "Top Rated"
	obj.Movie_list = arratoprated
	arraobj = append(arraobj, obj)

	sql_selectview := `SELECT 
		movieid, movietitle , movietype, label, COALESCE(posted_id,0) , urlthumbnail    
		FROM ` + config.DB_tbl_trx_movie + ` 
		WHERE enabled = 1 
		ORDER BY views DESC LIMIT 10     
	`
	row, err := con.QueryContext(ctx, sql_selectview)
	helpers.ErrorCheck(err)
	var objpopular entities.Model_movie
	var arraobjpopular []entities.Model_movie
	for row.Next() {
		var (
			movieid_db, posted_id_db                               int
			movietitle_db, movietype_db, label_db, urlthumbnail_db string
		)

		err = row.Scan(&movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db, &urlthumbnail_db)
		helpers.ErrorCheck(err)
		path_image := ""
		if urlthumbnail_db == "" {
			poster_image, poster_extension := _GetMedia(posted_id_db)
			path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		} else {
			path_image = urlthumbnail_db
		}

		objpopular.Movie_id = movieid_db
		objpopular.Movie_type = movietype_db
		objpopular.Movie_title = movietitle_db
		objpopular.Movie_label = label_db
		objpopular.Movie_thumbnail = path_image
		objpopular.Movie_video = ""
		arraobjpopular = append(arraobjpopular, objpopular)
		msg = "Success"
	}
	defer row.Close()
	obj.Movie_idcategory = 0
	obj.Movie_category = "Popular"
	obj.Movie_list = arraobjpopular
	arraobj = append(arraobj, obj)

	sql_selectcate := `SELECT 
		idgenre, nmgenre     
		FROM ` + config.DB_tbl_mst_movie_genre + ` as A 
		ORDER BY genredisplay ASC     
	`
	row_cate, err_cate := con.QueryContext(ctx, sql_selectcate)
	helpers.ErrorCheck(err_cate)
	for row_cate.Next() {
		var (
			idgenre_db int
			nmgenre_db string
		)
		err = row_cate.Scan(&idgenre_db, &nmgenre_db)
		helpers.ErrorCheck(err)

		sql_selectmovie := `SELECT 
			A.movieid, A.movietitle , A.movietype, A.label, COALESCE(A.posted_id,0) , A.urlthumbnail    
			FROM ` + config.DB_tbl_trx_movie + ` as A 
			JOIN ` + config.DB_tbl_trx_moviegenre + ` as B ON B.movieid = A.movieid  
			WHERE A.enabled = 1 
			AND B.idgenre = ?
			ORDER BY A.createdatemovie DESC LIMIT 10     
		`
		row_movie, err_movie := con.QueryContext(ctx, sql_selectmovie, idgenre_db)
		helpers.ErrorCheck(err_movie)
		var objpopular entities.Model_movie
		var arraobjpopular []entities.Model_movie
		for row_movie.Next() {
			var (
				movieid_db, posted_id_db                               int
				movietitle_db, movietype_db, label_db, urlthumbnail_db string
			)

			err = row_movie.Scan(&movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db, &urlthumbnail_db)
			helpers.ErrorCheck(err)
			path_image := ""
			if urlthumbnail_db == "" {
				poster_image, poster_extension := _GetMedia(posted_id_db)
				path_image = "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
			} else {
				path_image = urlthumbnail_db
			}

			objpopular.Movie_id = movieid_db
			objpopular.Movie_type = movietype_db
			objpopular.Movie_title = movietitle_db
			objpopular.Movie_label = label_db
			objpopular.Movie_thumbnail = path_image
			objpopular.Movie_video = ""
			arraobjpopular = append(arraobjpopular, objpopular)
			msg = "Success"
		}

		defer row_cate.Close()
		obj.Movie_idcategory = idgenre_db
		obj.Movie_category = nmgenre_db
		obj.Movie_list = arraobjpopular
		arraobj = append(arraobj, obj)
	}

	sql_slider := `SELECT 
		B.movieid, B.movietitle , B.movietype, B.label, COALESCE(B.posted_id,0) , A.url  
		FROM ` + config.DB_tbl_trx_slide + ` as A 
		JOIN ` + config.DB_tbl_trx_movie + ` as B ON B.movieid = A.movieid 
		ORDER BY position ASC      
	`
	row_slider, err_slider := con.QueryContext(ctx, sql_slider)
	helpers.ErrorCheck(err_slider)
	var objslider entities.Model_movie
	var arraobjslider []entities.Model_movie
	for row_slider.Next() {
		var (
			movieid_db, posted_id_db                      int
			movietitle_db, movietype_db, label_db, url_db string
		)

		err = row_slider.Scan(&movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db, &url_db)
		helpers.ErrorCheck(err)

		objslider.Movie_id = movieid_db
		objslider.Movie_type = movietype_db
		objslider.Movie_title = movietitle_db
		objslider.Movie_label = label_db
		objslider.Movie_thumbnail = url_db
		objslider.Movie_video = ""
		arraobjslider = append(arraobjslider, objslider)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Slider = arraobjslider
	res.Genre = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_moviecomment(movieid int) (helpers.Response, error) {
	var obj entities.Model_mobilemoviecomment
	var arraobj []entities.Model_mobilemoviecomment
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
		A.idcomment, B.nmuser , A.comment, A.createcomment 
		FROM ` + config.DB_tbl_trx_comment + ` as A 
		JOIN ` + config.DB_tbl_trx_user + ` as B ON B.username = A.username  
		WHERE A.idposter=?  
		AND B.statususer='Y' 
		ORDER BY A.createcomment DESC LIMIT 100   
	`

	row, err := con.QueryContext(ctx, sql_select, movieid)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idcomment_db                            int
			nmuser_db, comment_db, createcomment_db string
		)

		err = row.Scan(&idcomment_db, &nmuser_db, &comment_db, &createcomment_db)

		helpers.ErrorCheck(err)

		obj.Movie_idcomment = idcomment_db
		obj.Movie_name = nmuser_db
		obj.Movie_comment = comment_db
		obj.Movie_create = createcomment_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}

func Update_movieview(username string, idmovie int) bool {
	flag := false

	viewmovielast := _GetMovie(idmovie, "")
	viewmovienow := viewmovielast + 1

	sql_update := `
			UPDATE 
			` + config.DB_tbl_trx_movie + `  
			SET views =? 
			WHERE movieid =? 
		`

	flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_trx_movie, "UPDATE", viewmovienow, idmovie)

	if flag_update {
		flag = true
		log.Println(msg_update)
	} else {
		log.Println(msg_update)
	}

	return flag
}
func Save_moviecomment(username, comment string, idmovie int) bool {
	tglnow, _ := goment.New()
	flag := false

	field_column := config.DB_tbl_trx_comment + tglnow.Format("YYYY")
	idrecord_counter := Get_counter(field_column)
	sql_insert := `
			insert into
			` + config.DB_tbl_trx_comment + ` (
				idcomment , idposter, username, comment, statusread, createcomment
			) values (
				?,?,?,?,?,?
			)
		`
	flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_comment, "INSERT", tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idmovie, username, comment, "Y", tglnow.Format("YYYY-MM-DD HH:mm:ss"))

	if flag_insert {
		flag = true
		log.Println(msg_insert)
	} else {
		log.Println(msg_insert)
	}

	return flag
}
func Save_movierate(username, rating string, idmovie int) bool {
	tglnow, _ := goment.New()
	flag := false

	flag_check := CheckDBTwoField(config.DB_tbl_trx_rate, "username", username, "idposter", strconv.Itoa(idmovie))
	if !flag_check {
		field_column := config.DB_tbl_trx_rate + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		sql_insert := `
			insert into
			` + config.DB_tbl_trx_rate + ` (
				idrate , username, idposter, ratingposter, createrate
			) values (
				?,?,?,?,?
			)
		`
		flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_rate, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter),
			username, idmovie, rating, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			flag = true
			log.Println(msg_insert)
		} else {
			log.Println(msg_insert)
		}
	}

	return flag
}
func Save_moviefavorite(username string, idmovie int) bool {
	tglnow, _ := goment.New()
	flag := false

	flag_check := CheckDBTwoField(config.DB_tbl_trx_favorite, "username", username, "idposter", strconv.Itoa(idmovie))
	if !flag_check {
		field_column := config.DB_tbl_trx_favorite + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		sql_insert := `
			insert into
			` + config.DB_tbl_trx_favorite + ` (
				idfavorite , idposter, username, createfavorite
			) values (
				?,?,?,?
			)
		`
		flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_favorite, "INSERT",
			tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
			idmovie, username, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			flag = true
			log.Println(msg_insert)
		} else {
			log.Println(msg_insert)
		}
	}

	return flag
}
func Delete_moviefavorite(username string, idmovie int) bool {
	flag := false

	flag_check := CheckDBTwoField(config.DB_tbl_trx_favorite, "username", username, "idposter", strconv.Itoa(idmovie))
	if flag_check {
		sql_delete := `
			DELETE FROM 
			` + config.DB_tbl_trx_favorite + ` 
			WHERE idposter=? AND username=?
		`
		flag_delete, msg_delete := Exec_SQL(sql_delete, config.DB_tbl_trx_favorite, "DELETE", idmovie, username)

		if flag_delete {
			flag = true
			log.Println(msg_delete)
		} else {
			log.Println(msg_delete)
		}
	}

	return flag
}
func Save_moviereport(username, info string, idmovie int) bool {
	tglnow, _ := goment.New()
	flag := false

	info_1 := strings.ToUpper(info)
	info_final := strings.Replace(info_1, " ", "_", -1)

	field_column := config.DB_tbl_mst_report + tglnow.Format("YYYY")
	idrecord_counter := Get_counter(field_column)
	sql_insert := `
			insert into
			` + config.DB_tbl_mst_report + ` (
				idreport , idmovie, username, inforeport, poinreport, statusreport, createdatereport 
			) values (
				?,?,?,?,?,?,?
			)
		`
	flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_favorite, "INSERT",
		tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
		idmovie, username, info_final, 0, "PROCESS", tglnow.Format("YYYY-MM-DD HH:mm:ss"))

	if flag_insert {
		flag = true
		log.Println(msg_insert)
	} else {
		log.Println(msg_insert)
	}

	return flag
}
func Fetch_usermovie(username string) (helpers.Responsemobileuser, error) {
	var obj entities.Model_mobileuser
	var arraobj []entities.Model_mobileuser
	var res helpers.Responsemobileuser
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()
	tglnow, _ := goment.New()

	sql_select := `SELECT 
		username , nmuser, coderef, 
		point_in , point_out 
		FROM ` + config.DB_tbl_trx_user + ` 
		WHERE username=? 
	`

	row, err := con.QueryContext(ctx, sql_select, username)
	helpers.ErrorCheck(err)

	for row.Next() {
		var (
			point_in_db, point_out_db          int
			username_db, nmuser_db, coderef_db string
		)

		err := row.Scan(&username_db, &nmuser_db, &coderef_db, &point_in_db, &point_out_db)
		helpers.ErrorCheck(err)

		if coderef_db == "" {
			flag_coderef := true
			numbertemp := ""
			for {
				numbergenerate := helpers.GenerateNumber(7)
				flag_coderef = CheckDB(config.DB_tbl_trx_user, "coderef", numbergenerate)
				if !flag_coderef {
					numbertemp = numbergenerate
					break
				}
			}
			if !flag_coderef {
				sql_update := `
					UPDATE 
					` + config.DB_tbl_trx_user + ` 
					SET coderef=?, updatedateuser=? 
					WHERE username=? 
				`
				log.Println(username)
				log.Println(numbertemp)
				flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_trx_user, "UPDATE",
					numbertemp, tglnow.Format("YYYY-MM-DD HH:mm:ss"), username)

				if flag_update {
					coderef_db = numbertemp
					log.Println(msg_update)

				} else {
					log.Println(msg_update)
				}
			}
		}

		obj.User_username = username_db
		obj.User_name = nmuser_db
		obj.User_coderef = coderef_db
		obj.User_pointin = point_in_db
		obj.User_pointout = point_out_db
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	var objclaim entities.Model_mobilelistclaim
	var arraobjclaim []entities.Model_mobilelistclaim
	sql_selectlistclaim := `SELECT 
		idlistclaim , nmlistclaim, pointlistclaim 
		FROM ` + config.DB_tbl_mst_listclaim + ` 
		WHERE statuslistclaim='Y' 
		ORDER BY pointlistclaim ASC 
	`

	row_listclaim, err_listclaim := con.QueryContext(ctx, sql_selectlistclaim)
	helpers.ErrorCheck(err_listclaim)
	for row_listclaim.Next() {
		var (
			idlistclaim_db, pointlistclaim_db int
			nmlistclaim_db                    string
		)

		err := row_listclaim.Scan(&idlistclaim_db, &nmlistclaim_db, &pointlistclaim_db)
		helpers.ErrorCheck(err)

		objclaim.Claim_id = idlistclaim_db
		objclaim.Claim_name = nmlistclaim_db
		objclaim.Claim_point = pointlistclaim_db
		arraobjclaim = append(arraobjclaim, objclaim)
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Claim = arraobjclaim
	res.Time = time.Since(start).String()

	return res, nil
}

func Save_moviepoint(username, nmpoint string, idmovie, point int) bool {
	tglnow, _ := goment.New()
	flag := false

	field_column := config.DB_tbl_mst_point + tglnow.Format("YYYY")
	idrecord_counter := Get_counter(field_column)
	sql_insert := `
			insert into
			` + config.DB_tbl_mst_point + ` (
				idpoint , username, nmpoint, posted_id, point, createdatepoint
			) values (
				?,?,?,?,?,?
			)
		`
	flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_mst_point, "INSERT",
		tglnow.Format("YY")+tglnow.Format("MM")+strconv.Itoa(idrecord_counter),
		username, nmpoint, idmovie, point, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

	if flag_insert {
		flag = true
		log.Println(msg_insert)

	} else {
		log.Println(msg_insert)
	}

	return flag
}
func Update_pointmember(username string) bool {
	tglnow, _ := goment.New()
	flag := false

	point_total := _GetPoint_In(username)

	sql_update := `
		UPDATE 
		` + config.DB_tbl_trx_user + ` 
		SET point_in=?, updatedateuser=? 
		WHERE username=? 
	`
	log.Println(username)
	log.Println(point_total)
	flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_trx_user, "UPDATE",
		point_total, tglnow.Format("YYYY-MM-DD HH:mm:ss"), username)

	if flag_update {
		flag = true
		log.Println(msg_update)

	} else {
		log.Println(msg_update)
	}

	return flag
}
func _GetMedia(idrecord int) (string, string) {
	con := db.CreateCon()
	ctx := context.Background()
	url := ""
	extension := ""

	sql_select := `SELECT
		url, extension   
		FROM ` + config.DB_tbl_mst_mediatable + `  
		WHERE idmediatable = ? 
	`
	row := con.QueryRowContext(ctx, sql_select, idrecord)
	switch e := row.Scan(&url, &extension); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return url, extension
}
func _GetVideo(idrecord int, tipe string) (interface{}, int, string) {
	var obj entities.Model_movievideo
	var arraobj []entities.Model_movievideo
	con := db.CreateCon()
	ctx := context.Background()
	totalsource := 0
	source := ""
	sql_select := ""
	sql_select += "SELECT "
	sql_select += "url "
	sql_select += "FROM " + config.DB_tbl_mst_movie_source + " "
	sql_select += "WHERE poster_id = ? "
	if tipe == "single" {
		sql_select += "ORDER BY RAND() DESC LIMIT 1 "
	}

	row_select, err_select := con.QueryContext(ctx, sql_select, idrecord)
	helpers.ErrorCheck(err_select)
	for row_select.Next() {
		totalsource = totalsource + 1
		var url_db string

		err_select = row_select.Scan(&url_db)
		helpers.ErrorCheck(err_select)
		source = url_db
		obj.Movie_src = url_db
		arraobj = append(arraobj, obj)
	}
	return arraobj, totalsource, source
}
func _GetVideoEpisode(idrecord int) string {
	con := db.CreateCon()
	ctx := context.Background()
	url := ""

	sql_select := `SELECT
		url   
		FROM ` + config.DB_tbl_mst_movie_source + `  
		WHERE episode_id = ? 
	`
	row := con.QueryRowContext(ctx, sql_select, idrecord)
	switch e := row.Scan(&url); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return url
}
func _GetMovie(idrecord int, tipe string) int {
	con := db.CreateCon()
	ctx := context.Background()
	views := 0

	sql_select := `SELECT
		views    
		FROM ` + config.DB_tbl_trx_movie + `  
		WHERE movieid = ? 
	`
	row := con.QueryRowContext(ctx, sql_select, idrecord)
	switch e := row.Scan(&views); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return views
}
func _GetPoint_In(username string) int {
	con := db.CreateCon()
	ctx := context.Background()
	totalpoint := 0

	sql_select := `SELECT
		SUM(point) as totalpoint     
		FROM ` + config.DB_tbl_mst_point + `  
		WHERE username = ? 
	`
	row := con.QueryRowContext(ctx, sql_select, username)
	switch e := row.Scan(&totalpoint); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}
	return totalpoint
}
func _GetFavorite(idrecord int, username string) string {
	con := db.CreateCon()
	ctx := context.Background()
	idfavorite := 0
	favorite := "N"

	sql_select := `SELECT
		idfavorite   
		FROM ` + config.DB_tbl_trx_favorite + `  
		WHERE idposter = ? AND username=? 
	`
	row := con.QueryRowContext(ctx, sql_select, idrecord, username)
	switch e := row.Scan(&idfavorite); e {
	case sql.ErrNoRows:
	case nil:
		favorite = "Y"
	default:
		helpers.ErrorCheck(e)
	}
	return favorite
}
