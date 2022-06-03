package main

import (
	"log"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/db"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/routers"
)

func main() {
	db.Init()
	app := routers.Init()
	log.Fatal(app.Listen(":7072"))
}
