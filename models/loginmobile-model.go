package models

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/nikitamirzani323/isbpanel_api/config"
	"github.com/nikitamirzani323/isbpanel_api/db"
	"github.com/nleeper/goment"
)

func Loginmobile_Model(username string) (bool, error) {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	tglnow, _ := goment.New()
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
		return false, errors.New("Username and Password Not Found")
	case nil:
		flag = true
	default:
		return false, errors.New("Username and Password Not Found")
	}

	if flag {
		sql_update := `
			UPDATE ` + config.DB_tbl_trx_user + ` 
			SET lastlogin=?
			WHERE username  = ? 
		`
		flag_update, msg_update := Exec_SQL(sql_update, config.DB_tbl_trx_user, "UPDATE",
			tglnow.Format("YYYY-MM-DD HH:mm:ss"),
			username)

		if flag_update {
			flag = true
			log.Println(msg_update)
		} else {
			log.Println(msg_update)
		}
	}

	return true, nil
}
