package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"

	"ilmudata/task1/controllers"
	"ilmudata/task1/models"

	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	// session
	store := session.New()

	// load template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:     engine,
		BodyLimit: 200 * 1024 * 1024,
	})

	// static
	app.Static("/public", "./public")
	models.InitDbModels()

	userController := controllers.InitUserController(store)
	videoController := controllers.InitVideoController(store)
	advertController := controllers.InitAdsController(store)

	app.Get("/", userController.ViewHome)

	ads := app.Group("/ads")
	ads.Get("/", advertController.Ads)
	ads.Get("/viewiklan", advertController.GetAds)
	
	user := app.Group("")
	user.Get("/login", userController.Login)
	user.Post("/loginverify", userController.LoginPostVerify)
	user.Get("/logout", userController.Logout)

	//Testing JWT
	data := app.Group("/data")
	data.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mysecretpassword"),
	}))
	data.Get("/dashboarduser", userController.DashboardUser)

	video := app.Group("/videos")
	// video := app.Group("/videos")
	video.Get("/", videoController.IndexVideo)
	video.Get("/create", videoController.AddVideo)
	video.Post("/create", videoController.AddPostedVideo)
	video.Get("/editvideo/:id", videoController.EditVideo)
	video.Post("/editvideo/:id", videoController.EditPostedVideo)
	video.Get("/deletevideo/:id", videoController.DeleteVideo)
	video.Get("/detailvideo/:id", videoController.GetViewVideo)

	// advert := app.Group("/advert")

	app.Listen(":3001")
}
