package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/isbpanel_api/controllers"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Post("/api/pasaran", controllers.Pasaranhome)
	app.Post("/api/keluaran", controllers.Keluaranhome)
	app.Post("/api/news", controllers.Newshome)
	app.Post("/api/bukumimpi", controllers.Bukumimpihome)
	app.Post("/api/tafsirmimpi", controllers.TafsirMimpihome)

	app.Post("/api/movie", controllers.Moviehome)
	return app
}
