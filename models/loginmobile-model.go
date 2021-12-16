package models

import (
	"context"
	"database/sql"
	"log"

	"github.com/nikitamirzani323/isbpanel_api/config"
	"github.com/nikitamirzani323/isbpanel_api/db"
	"github.com/nikitamirzani323/isbpanel_api/helpers"
	"github.com/nleeper/goment"
)

func Loginmobile_Model(username, name, tipe string) bool {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	tglnow, _ := goment.New()
	if tipe == "LOGIN" {
		sql_select := `
			SELECT
			username     
			FROM ` + config.DB_tbl_trx_user + ` 
			WHERE username  = ?
			AND statususer = "Y" 
		`

		row := con.QueryRowContext(ctx, sql_select, username)
		switch e := row.Scan(&username); e {
		case sql.ErrNoRows:
		case nil:
			flag = true
		default:
		}

		if flag {
			//UPDATE USER
			sql_update := `
				UPDATE ` + config.DB_tbl_trx_user + ` 
				SET lastlogin=?
				WHERE username  = ? 
			`
			flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_trx_user, "UPDATE",
				tglnow.Format("YYYY-MM-DD HH:mm:ss"),
				username)

			if flag_update {
				log.Println(msg_update)
			} else {
				log.Println(msg_update)
			}
		}
	} else {
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
		sql_insert := `
				INSERT INTO ` + config.DB_tbl_trx_user + ` 
				(username, nmuser,typeuser,coderef,point_in,point_out,statususer,lastlogin,createdateuser)
				VALUES (?,?,?,?,?,?,?,?,?)
			`
		flag_insert, msg_insert := Exec_SQL(sql_insert, config.DB_tbl_trx_user, "INSERT",
			username, name, "phone", numbertemp, 0, 0, "Y", tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			log.Println(msg_insert)
		} else {
			log.Println(msg_insert)
		}
	}

	return flag
}
func Mobileversion_Model() string {
	con := db.CreateCon()
	ctx := context.Background()
	version_db := ""
	sql_select := `
			SELECT
			version     
			FROM ` + config.DB_tbl_mst_version + ` 
			WHERE idversion  = '1' 
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&version_db); e {
	case sql.ErrNoRows:

	case nil:

	default:

	}

	return version_db
}
