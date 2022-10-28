package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"ilmudata/task1/database"
	"ilmudata/task1/models"

	"github.com/golang-jwt/jwt/v4"
)

type UserController struct {
	// declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitUserController(s *session.Store) *UserController {
	db := database.InitDb()
	return &UserController{Db: db, store: s}
}

// GET /login
func (controller *UserController) ViewHome(c *fiber.Ctx) error {
	return c.Render("myview", fiber.Map{
		"Title": "Halaman Depan",
	})
}

// GET /login
func (controller *UserController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// post /login
func (controller *UserController) LoginPostVerify(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	var names models.LoginForm
	names.Name = name

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	println(name)
	sess.Set("name", "names.Name")
	sess.Save()

	val := sess.Get("name")

	str := fmt.Sprintf("%v", val)
	fmt.Println(str)
	return c.SendString("Welcome " + str)
}

// /logout
func (controller *UserController) Logout(c *fiber.Ctx) error {

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Destroy()
	return c.Render("login", fiber.Map{
		"Title": "Telah Logout",
	})
}

func (controller *UserController) AuthVerify(c *fiber.Ctx) error {
	sess, _ := controller.store.Get(c)
	val := sess.Get("name")
	if val != nil {
		return c.Next()
	}
	return c.Redirect("/login")
}

func (controller *UserController) DashboardUser(c *fiber.Ctx) error {
	return c.Render("dashboarduser", fiber.Map{
		"Title": "Dashboard User",
	})
}
