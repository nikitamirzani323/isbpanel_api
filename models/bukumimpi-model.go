package models

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/isbpanel_api/config"
	"github.com/nikitamirzani323/isbpanel_api/db"
	"github.com/nikitamirzani323/isbpanel_api/entities"
	"github.com/nikitamirzani323/isbpanel_api/helpers"
)

func Fetch_bukumimpiHome(tipe, nama string) (helpers.Response, error) {
	var obj entities.Model_bukumimpi
	var arraobj []entities.Model_bukumimpi
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()
	sql_select := ""
	if tipe == "" {
		sql_select += ""
		sql_select += "SELECT "
		sql_select += "typebukumimpi, nmbukumimpi, nmrbukumimpi "
		sql_select += "FROM " + config.DB_tbl_trx_bukumimpi + " "
		if nama != "" {
			sql_select += "WHERE nmbukumimpi LIKE '%" + nama + "%' "
			sql_select += "ORDER BY nmbukumimpi ASC LIMIT 500 "
		} else {
			sql_select += "ORDER BY RAND() LIMIT 500 "
		}

	} else {
		sql_select += ""
		sql_select += "SELECT "
		sql_select += "typebukumimpi, nmbukumimpi, nmrbukumimpi "
		sql_select += "FROM " + config.DB_tbl_trx_bukumimpi + " "
		if nama != "" {
			sql_select += "WHERE nmbukumimpi LIKE '%" + nama + "%' "
			sql_select += "AND typebukumimpi='" + tipe + "' "
		} else {
			sql_select += "WHERE typebukumimpi='" + tipe + "' "
		}
		sql_select += "ORDER BY nmbukumimpi ASC LIMIT 500 "
	}

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			typebukumimpi_db, nmbukumimpi_db, nmrbukumimpi_db string
		)

		err = row.Scan(&typebukumimpi_db, &nmbukumimpi_db, &nmrbukumimpi_db)

		helpers.ErrorCheck(err)

		obj.Bukumimpi_type = typebukumimpi_db
		obj.Bukumimpi_name = nmbukumimpi_db
		obj.Bukumimpi_nomor = nmrbukumimpi_db
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
