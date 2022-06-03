package models

import (
	"context"
	"database/sql"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/config"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/db"
)

func Get_Domain(nmdomain string) bool {
	con := db.CreateCon()
	ctx := context.Background()
	flag := false
	statusdomain := ""

	sql_select := `SELECT
		statusdomain  
		FROM ` + config.DB_tbl_mst_domain + `  
		WHERE nmdomain = $1 
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
