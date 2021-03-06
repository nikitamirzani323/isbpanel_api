package routers

import (
	"time"

	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/controllers"
	"bitbucket.org/isbtotogroup/isbpanel_api_frontend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        200,
		Expiration: 20 * time.Second,
		LimitReached: func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"status":  fiber.StatusTooManyRequests,
				"message": "error many request",
				"record":  nil,
			})
		},
	}))
	app.Post("/api/init", controllers.CheckLogin)
	app.Post("/api/banner", controllers.Bannerhome)
	app.Post("/api/pasaran", controllers.Pasaranhome)
	app.Post("/api/keluaran", controllers.Keluaranhome)
	app.Post("/api/news", controllers.Newshome)
	app.Post("/api/newsmovie", controllers.Newsmoviehome)
	app.Post("/api/bukumimpi", controllers.Bukumimpihome)
	app.Post("/api/tafsirmimpi", controllers.TafsirMimpihome)
	app.Post("/api/providerslot", controllers.Providerslothome)
	app.Post("/api/providerslotdetail", controllers.Prediksislotdetail)
	app.Post("/api/prediksislot", controllers.Prediksislothome)

	app.Post("/api/home", middleware.JWTProtected(), controllers.Home)
	app.Post("/api/movie", middleware.JWTProtected(), controllers.Moviehome)
	app.Post("/api/season", middleware.JWTProtected(), controllers.Movieseason)
	app.Post("/api/episode", middleware.JWTProtected(), controllers.Movieepisode)

	//MOBILE
	app.Post("/api/mobile/login", controllers.CheckLoginmobile)
	app.Post("/api/mobile/frontpagemovie", controllers.Moviefrontpagemobile)
	app.Post("/api/mobile/listmovie", controllers.Moviemobile)
	app.Post("/api/mobile/listgenremovie", controllers.Moviegenremobile)
	app.Post("/api/mobile/moviedetail", controllers.Moviedetailmobile)
	app.Post("/api/mobile/season", controllers.Moviemobileseason)
	app.Post("/api/mobile/episode", controllers.Moviemobileepisode)
	app.Post("/api/mobile/comment", controllers.Moviecommentmobile)
	app.Post("/api/mobile/savecomment", controllers.Moviecommentsave)
	app.Post("/api/mobile/saverate", controllers.Movieratesave)
	app.Post("/api/mobile/savefavorite", controllers.Moviefavoritesave)
	app.Post("/api/mobile/deletefavorite", controllers.Moviefavoritedelete)
	app.Post("/api/mobile/savereport", controllers.Moviereportsave)

	app.Post("/api/mobile/userdetail", controllers.Movieuserdetail)
	app.Post("/api/mobile/saveuserclaim", controllers.Movieclaimsave)
	return app
}
