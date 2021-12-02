package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_api/config"
	"github.com/nikitamirzani323/isbpanel_api/db"
	"github.com/nikitamirzani323/isbpanel_api/entities"
	"github.com/nikitamirzani323/isbpanel_api/helpers"
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

		movie_url, _ := _GetVideo(movieid_db)

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

		movie_url, _ := _GetVideo(movieid_db)

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
		movie_url, _ := _GetVideo(movieid_db)

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

		movie_url, _ := _GetVideo(movieid_db)

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
func Fetch_movielist(tipe string) (helpers.Response, error) {
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
		sql_select += "movieid, posted_id, movietitle, movietype, year, views, urlthumbnail, label, description "
		sql_select += "FROM " + config.DB_VIEW_MOVIE + " "
		sql_select += "ORDER BY RAND() DESC LIMIT 300 "
	} else if tipe == "series" {
		sql_select += "SELECT "
		sql_select += "movieid, posted_id, movietitle, movietype, year, views, urlthumbnail, label, description "
		sql_select += "FROM " + config.DB_VIEW_MOVIESERIES + " "
		sql_select += "ORDER BY RAND() DESC LIMIT 300 "
	} else {
		sql_select += "SELECT "
		sql_select += "movieid, posted_id, movietitle, movietype, year, views, urlthumbnail, label, description "
		sql_select += "FROM " + config.DB_tbl_trx_movie + " "
		sql_select += "WHERE enabled='1' "
		sql_select += "ORDER BY RAND() DESC LIMIT 30 "
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
func _GetVideo(idrecord int) (interface{}, int) {
	var obj entities.Model_movievideo
	var arraobj []entities.Model_movievideo
	con := db.CreateCon()
	ctx := context.Background()
	totalsource := 0
	sql_select := `SELECT
		url   
		FROM ` + config.DB_tbl_mst_movie_source + `  
		WHERE poster_id = ? 
	`
	row_select, err_select := con.QueryContext(ctx, sql_select, idrecord)
	helpers.ErrorCheck(err_select)
	for row_select.Next() {
		totalsource = totalsource + 1
		var url_db string

		err_select = row_select.Scan(&url_db)
		helpers.ErrorCheck(err_select)

		obj.Movie_src = url_db
		arraobj = append(arraobj, obj)
	}
	return arraobj, totalsource
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
