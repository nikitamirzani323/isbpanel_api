package models

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/entities"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/nleeper/goment"
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
			urlpasaran , pasarandiundi, jamjadwal::text as jamjadwal  
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
			to_char(datekeluaran, 'YYYY-MM-DD') , nomorkeluaran
			FROM ` + config.DB_tbl_trx_keluaran + ` 
			WHERE idpasarantogel = $1  
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
			to_char(dateprediksi, 'YYYY-MM-DD') , bbfsprediksi, nomorprediksi
			FROM ` + config.DB_tbl_trx_prediksi + ` 
			WHERE idpasarantogel = $1  
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
func Fetch_keluaran(idpasaran string) (helpers.ResponseKeluaran, error) {
	var obj entities.Model_keluaran
	var arraobj []entities.Model_keluaran
	var res helpers.ResponseKeluaran
	var myDays = []string{"minggu", "senin", "selasa", "rabu", "kamis", "jumat", "sabtu"}
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()
	tglnow, _ := goment.New()
	startyear := tglnow.Format("YYYY") + "-01-01"
	endyear := tglnow.Format("YYYY") + "-12-31"

	sql_select := `SELECT 
			to_char(datekeluaran, 'YYYY-MM-DD') as datekeluaran , periodekeluaran ,nomorkeluaran
			FROM ` + config.DB_tbl_trx_keluaran + ` 
			WHERE idpasarantogel=$1  
			AND datekeluaran >= $2  
			AND datekeluaran <= $3  
			ORDER BY datekeluaran DESC   
		`

	row, err := con.QueryContext(ctx, sql_select, idpasaran, startyear, endyear)
	helpers.ErrorCheck(err)
	var objpaito_minggu entities.Model_keluaranpaitominggu
	var arraobjpaito_minggu []entities.Model_keluaranpaitominggu
	var objpaito_senin entities.Model_keluaranpaitosenin
	var arraobjpaito_senin []entities.Model_keluaranpaitosenin
	var objpaito_selasa entities.Model_keluaranpaitoselasa
	var arraobjpaito_selasa []entities.Model_keluaranpaitoselasa
	var objpaito_rabu entities.Model_keluaranpaitorabu
	var arraobjpaito_rabu []entities.Model_keluaranpaitorabu
	var objpaito_kamis entities.Model_keluaranpaitokamis
	var arraobjpaito_kamis []entities.Model_keluaranpaitokamis
	var objpaito_jumat entities.Model_keluaranpaitojumat
	var arraobjpaito_jumat []entities.Model_keluaranpaitojumat
	var objpaito_sabtu entities.Model_keluaranpaitosabtu
	var arraobjpaito_sabtu []entities.Model_keluaranpaitosabtu
	for row.Next() {
		var (
			datekeluaran_db, periodekeluaran_db, nomorkeluaran_db string
		)

		err = row.Scan(&datekeluaran_db, &periodekeluaran_db, &nomorkeluaran_db)
		helpers.ErrorCheck(err)

		tgldatekeluaran, _ := goment.New(datekeluaran_db)
		daynow := tgldatekeluaran.Format("d")
		intVar, _ := strconv.ParseInt(daynow, 0, 8)
		daynowhari := myDays[intVar]

		switch daynowhari {
		case "minggu":
			objpaito_minggu.Keluaran_nomor = nomorkeluaran_db
			arraobjpaito_minggu = append(arraobjpaito_minggu, objpaito_minggu)
		case "senin":
			objpaito_senin.Keluaran_nomor = nomorkeluaran_db
			arraobjpaito_senin = append(arraobjpaito_senin, objpaito_senin)
		case "selasa":
			objpaito_selasa.Keluaran_nomor = nomorkeluaran_db
			arraobjpaito_selasa = append(arraobjpaito_selasa, objpaito_selasa)
		case "rabu":
			objpaito_rabu.Keluaran_nomor = nomorkeluaran_db
			arraobjpaito_rabu = append(arraobjpaito_rabu, objpaito_rabu)
		case "kamis":
			objpaito_kamis.Keluaran_nomor = nomorkeluaran_db
			arraobjpaito_kamis = append(arraobjpaito_kamis, objpaito_kamis)
		case "jumat":
			objpaito_jumat.Keluaran_nomor = nomorkeluaran_db
			arraobjpaito_jumat = append(arraobjpaito_jumat, objpaito_jumat)
		case "sabtu":
			objpaito_sabtu.Keluaran_nomor = nomorkeluaran_db
			arraobjpaito_sabtu = append(arraobjpaito_sabtu, objpaito_sabtu)
		}
		obj.Keluaran_datekeluaran = datekeluaran_db
		obj.Keluaran_periode = idpasaran + "-" + periodekeluaran_db
		obj.Keluaran_nomor = nomorkeluaran_db
		arraobj = append(arraobj, obj)
		nomorkeluaran_db = ""
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Paito_minggu = arraobjpaito_minggu
	res.Paito_senin = arraobjpaito_senin
	res.Paito_selasa = arraobjpaito_selasa
	res.Paito_rabu = arraobjpaito_rabu
	res.Paito_kamis = arraobjpaito_kamis
	res.Paito_jumat = arraobjpaito_jumat
	res.Paito_sabtu = arraobjpaito_sabtu
	res.Time = time.Since(start).String()

	return res, nil
}
