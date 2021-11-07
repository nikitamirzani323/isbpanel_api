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

func Fetch_pasaranHome() (helpers.Response, error) {
	var obj entities.Model_pasaran
	var arraobj []entities.Model_pasaran
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			idpasarantogel , nmpasarantogel, 
			urlpasaran , pasarandiundi, jamjadwal 
			FROM ` + config.DB_tbl_mst_pasaran + ` 
			WHERE statuspasaran = 'Y' 
			ORDER BY displaypasaran ASC  
		`

	row, err := con.QueryContext(ctx, sql_select)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idpasarantogel_db, nmpasarantogel_db, urlpasaran_db, pasarandiundi_db, jamjadwal_db string
		)

		err = row.Scan(&idpasarantogel_db, &nmpasarantogel_db, &urlpasaran_db, &pasarandiundi_db, &jamjadwal_db)

		helpers.ErrorCheck(err)

		var (
			datekeluaran_db, nomorkeluaran_db                  string
			dateprediksi_db, bbfsprediksi_db, nomorprediksi_db string
		)
		sql_selectpasaran := `SELECT 
			datekeluaran , nomorkeluaran
			FROM ` + config.DB_tbl_trx_keluaran + ` 
			WHERE idpasarantogel = ? 
			ORDER BY datekeluaran DESC LIMIT 1
		`
		row_keluaran := con.QueryRowContext(ctx, sql_selectpasaran, idpasarantogel_db)
		switch e_keluaran := row_keluaran.Scan(&datekeluaran_db, &nomorkeluaran_db); e_keluaran {
		case sql.ErrNoRows:
		case nil:
		default:
			helpers.ErrorCheck(e_keluaran)
		}

		sql_selectprediksi := `SELECT 
			dateprediksi , bbfsprediksi, nomorprediksi
			FROM ` + config.DB_tbl_trx_prediksi + ` 
			WHERE idpasarantogel = ? 
			ORDER BY dateprediksi DESC LIMIT 1
		`
		row_prediksi := con.QueryRowContext(ctx, sql_selectprediksi, idpasarantogel_db)
		switch e_prediksi := row_prediksi.Scan(&dateprediksi_db, &bbfsprediksi_db, &nomorprediksi_db); e_prediksi {
		case sql.ErrNoRows:
		case nil:
		default:
			helpers.ErrorCheck(e_prediksi)
		}

		obj.Pasaran_id = idpasarantogel_db
		obj.Pasaran_name = nmpasarantogel_db
		obj.Pasaran_url = urlpasaran_db
		obj.Pasaran_diundi = pasarandiundi_db
		obj.Pasaran_jamjadwal = jamjadwal_db
		obj.Pasaran_datekeluaran = datekeluaran_db
		obj.Pasaran_keluaran = nomorkeluaran_db
		obj.Pasaran_dateprediksi = dateprediksi_db
		obj.Pasaran_bbfsprediksi = bbfsprediksi_db
		obj.Pasaran_nomorprediksi = nomorprediksi_db
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
