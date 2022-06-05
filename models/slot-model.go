package models

import (
	"context"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"github.com/gofiber/fiber/v2"
)

func Fetch_providerslotHome() (helpers.Response, error) {
	var obj entities.Model_providerslot
	var arraobj []entities.Model_providerslot
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
		nmproviderslot , providerslot_slug, providerslot_image, 
		providerslot_title , providerslot_descp 
		FROM ` + config.DB_tbl_mst_providerslot + ` 
		WHERE providerslot_status = 'Y'  
		ORDER BY providerslot_display ASC    
	`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			nmproviderslot_db, providerslot_slug_db, providerslot_image_db string
			providerslot_title_db, providerslot_descp_db                   string
		)

		err = row.Scan(&nmproviderslot_db, &providerslot_slug_db, &providerslot_image_db,
			&providerslot_title_db, &providerslot_descp_db)

		helpers.ErrorCheck(err)

		obj.Providerslot_name = nmproviderslot_db
		obj.Providerslot_slug = providerslot_slug_db
		obj.Providerslot_image = providerslot_image_db
		obj.Providerslot_title = providerslot_title_db
		obj.Providerslot_descp = providerslot_descp_db
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
func Fetch_prediksislot(idprovider, limit int) (helpers.Response, error) {
	var obj entities.Model_prediksislot
	var arraobj []entities.Model_prediksislot
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := ""
	sql_select += ""
	sql_select += "SELECT "
	sql_select += "nmgameslot , gameslot_image, gameslot_prediksi "
	sql_select += "FROM " + config.DB_tbl_trx_prediksislot + "   "
	if idprovider > 0 {
		sql_select += "WHERE idproviderslot ='" + strconv.Itoa(idprovider) + "'   "
	}
	if limit > 0 {
		sql_select += "ORDER BY gameslot_prediksi DESC LIMIT " + strconv.Itoa(limit) + "   "
	} else {
		sql_select += "ORDER BY gameslot_prediksi DESC    "
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			nmgameslot_db, gameslot_image_db string
			gameslot_prediksi_db             int
		)

		err = row.Scan(&nmgameslot_db, &gameslot_image_db, &gameslot_prediksi_db)

		helpers.ErrorCheck(err)

		obj.Prediksislot_name = nmgameslot_db
		obj.Prediksislot_image = gameslot_image_db
		obj.Prediksislot_prediksi = gameslot_prediksi_db
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
