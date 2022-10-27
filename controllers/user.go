package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
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

func (controller *UserController) LoginPosted(c *fiber.Ctx) error {
	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	var user models.User
	var myform models.LoginForm

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/login")
	}

	// Find user
	errs := models.FindUserByUsername(controller.Db, &user, myform.Username)
	if errs != nil {
		return c.Redirect("/login") // Unsuccessful login (cannot find user)
	}

	// Compare password
	compare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(myform.Password))
	if compare == nil { // compare == nil artinya hasil compare di atas true
		sess.Set("username", user.Username)
		sess.Set("userId", user.Id)
		sess.Save()

		idn := strconv.FormatUint(uint64(user.Id), 10)
		return c.Redirect("/products/" + idn)
	}

	return c.Redirect("/login")
}
