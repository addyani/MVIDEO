package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"

	"ilmudata/task1/controllers"
	"ilmudata/task1/models"
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
	// videoController := controllers.InitProductController(store)
	// advertController := controllers.InitCartController(store)

	user := app.Group("")
	user.Post("/login", userController.LoginPosted)

	// video := app.Group("/videos")

	// advert := app.Group("/advert")

	app.Listen(":3000")
}
