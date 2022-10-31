package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	"ilmudata/task1/database"
	"ilmudata/task1/models"
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

	// var names models.LoginForm
	// names.Name = name

	// sess, err := controller.store.Get(c)
	// if err != nil {
	// 	panic(err)
	// }
	// println(name)
	// sess.Set("name", "names.Name")
	// sess.Save()

	// val := sess.Get("name")

	// str := fmt.Sprintf("%v", val)
	// fmt.Println(str)

	//deklarasi algoritma sign in
	// tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY3MjI2MDIyLCJuYW1lIjoidXNlcjEifQ.IzusrDy1UP1Sz8wPRWFIC3uezjJC19FG6tR5ehnj1uQ")
	// //signed token
	// _, err := tokenAlgo.SignedString(config.JWT_KEY)
	// if err != nil {
	// 	response := map[string]string{"message": err.Error()}
	// 	helper.ResponseJSON(w, http.StatusInternalServerError, response)
	// 	return err
	// }

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "token",
	// 	Path:     "/",
	// 	Value:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY3MjI2MDIyLCJuYW1lIjoidXNlcjEifQ.IzusrDy1UP1Sz8wPRWFIC3uezjJC19FG6tR5ehnj1uQ",
	// 	HTTPOnly: true,
	// })

	var myform models.LoginForm
	if err := c.BodyParser(&myform.Token); err != nil {
		return c.SendString("Not Token Detect")
	}

	return c.SendString("Add JWT To Cookie")
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
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.Render("dashboarduser", fiber.Map{
		"Title": "Dashboard User",
		"Name":  name,
	})
}
