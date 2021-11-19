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

	sql_selectview := `SELECT 
		movieid, movietitle , movietype, label, posted_id   
		FROM ` + config.DB_tbl_trx_movie + ` 
		ORDER BY views DESC LIMIT 24     
	`
	row, err := con.QueryContext(ctx, sql_selectview)
	helpers.ErrorCheck(err)
	var objpopular entities.Model_movie
	var arraobjpopular []entities.Model_movie
	for row.Next() {
		var (
			movieid_db, posted_id_db              int
			movietitle_db, movietype_db, label_db string
		)

		err = row.Scan(&movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db)
		helpers.ErrorCheck(err)
		poster_image, poster_extension := _GetMedia(posted_id_db)
		path_image := "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		movie_url := _GetVideo(movieid_db)

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
		movieid, movietitle , movietype, label, posted_id   
		FROM ` + config.DB_tbl_trx_movie + ` 
		ORDER BY createdatemovie DESC LIMIT 24	     
	`
	var objchildbaru entities.Model_movie
	var arraobjchildbaru []entities.Model_movie
	row_new, err_new := con.QueryContext(ctx, sql_selectnew)
	helpers.ErrorCheck(err_new)
	for row_new.Next() {
		var (
			movieid_db, posted_id_db              int
			movietitle_db, movietype_db, label_db string
		)

		err = row_new.Scan(&movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db)
		helpers.ErrorCheck(err)
		poster_image, poster_extension := _GetMedia(posted_id_db)
		path_image := "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		movie_url := _GetVideo(movieid_db)

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
		movieid, movietitle , movietype, label, posted_id   
		FROM ` + config.DB_tbl_trx_movie + ` 
		ORDER BY RAND() LIMIT 48      
	`
	var objrekomendasi entities.Model_movie
	var arraobjrekomendasi []entities.Model_movie
	row_rekomendasi, err_rekomendasi := con.QueryContext(ctx, sql_selectrekomendasi)
	helpers.ErrorCheck(err_rekomendasi)
	for row_rekomendasi.Next() {
		var (
			movieid_db, posted_id_db              int
			movietitle_db, movietype_db, label_db string
		)

		err = row_rekomendasi.Scan(&movieid_db, &movietitle_db, &movietype_db, &label_db, &posted_id_db)
		helpers.ErrorCheck(err)
		poster_image, poster_extension := _GetMedia(posted_id_db)
		path_image := "https://duniafilm.b-cdn.net/uploads/cache/poster_thumb/uploads/" + poster_extension + "/" + poster_image
		movie_url := _GetVideo(movieid_db)

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
func _GetVideo(idrecord int) string {
	con := db.CreateCon()
	ctx := context.Background()
	url := ""

	sql_select := `SELECT
		url   
		FROM ` + config.DB_tbl_mst_movie_source + `  
		WHERE poster_id = ? 
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
