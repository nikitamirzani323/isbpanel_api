package models

import (
	"context"
	"database/sql"

	"github.com/nikitamirzani323/isbpanel_api/config"
	"github.com/nikitamirzani323/isbpanel_api/db"
)

func Get_Domain(nmdomain string) bool {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	statusdomain := ""

	sql_select := `SELECT
		statusdomain  
		FROM ` + config.DB_tbl_mst_domain + `  
		WHERE nmdomain = ? 
		AND statusdomain = 'RUNNING' 
	`
	row := con.QueryRowContext(ctx, sql_select, nmdomain)
	switch e := row.Scan(&statusdomain); e {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true
	default:
		flag = false
	}

	return flag
}
