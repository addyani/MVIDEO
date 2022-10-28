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

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("mysecretpassword"),
	}))

	// static
	app.Static("/public", "./public")
	models.InitDbModels()

	userController := controllers.InitUserController(store)
	videoController := controllers.InitVideoController(store)
	advertController := controllers.InitAdsController(store)

	ads := app.Group("/ads")
	ads.Get("/", advertController.Ads)

	user := app.Group("")
	user.Get("/login", userController.Login)
	user.Get("/loginverify", userController.LoginPostVerify)
	user.Get("/logout", userController.Logout)
	//userController.AuthVerify for auth
	user.Get("/dashboarduser", userController.AuthVerify, userController.DashboardUser)

	//Untuk testing harus ada user register agar foreignKey relasi user to video dan advert bisa tersambung
	// user.Post("/register", userController.AddRegisteredUser)

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
