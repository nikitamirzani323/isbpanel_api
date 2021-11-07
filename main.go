package main

import (
	"log"

	"github.com/nikitamirzani323/isbpanel_api/db"
	"github.com/nikitamirzani323/isbpanel_api/routers"
)

func main() {
	db.Init()
	app := routers.Init()
	log.Fatal(app.Listen(":7072"))
}
