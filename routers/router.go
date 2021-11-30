package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/isbpanel_api/controllers"
	"github.com/nikitamirzani323/isbpanel_api/middleware"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Post("/api/init", controllers.CheckLogin)
	app.Post("/api/home", middleware.JWTProtected(), controllers.Home)
	app.Post("/api/pasaran", controllers.Pasaranhome)
	app.Post("/api/keluaran", controllers.Keluaranhome)
	app.Post("/api/news", controllers.Newshome)
	app.Post("/api/newsmovie", controllers.Newsmoviehome)
	app.Post("/api/bukumimpi", controllers.Bukumimpihome)
	app.Post("/api/tafsirmimpi", controllers.TafsirMimpihome)

	app.Post("/api/movie", middleware.JWTProtected(), controllers.Moviehome)
	app.Post("/api/season", middleware.JWTProtected(), controllers.Movieseason)
	app.Post("/api/episode", middleware.JWTProtected(), controllers.Movieepisode)

	//MOBILE
	app.Post("/api/mobile/login", controllers.CheckLoginmobile)
	app.Post("/api/mobile/listmovie", controllers.Moviemobile)
	return app
}
