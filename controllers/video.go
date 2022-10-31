package controllers //controler akan mengambil model

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"ilmudata/task1/database"
	"ilmudata/task1/models"
)

type VideoController struct {
	Db    *gorm.DB
	store *session.Store
}

func InitVideoController(s *session.Store) *VideoController {
	db := database.InitDb()

	db.AutoMigrate(&models.Video{})
	return &VideoController{Db: db, store: s}
}

// routing
// GET /videos
func (controller *VideoController) IndexVideo(c *fiber.Ctx) error {
	var videos []models.Video
	err := models.ReadVideo(controller.Db, &videos)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("homevideo", fiber.Map{
		"Title":  "Dasboard",
		"Videos": videos,
	})

}

// GET /videos/create
func (controller *VideoController) AddVideo(c *fiber.Ctx) error {
	return c.Render("uploadvideo", fiber.Map{
		"Title": "Upload Video",
	})
}

// POST /videos/create
func (controller *VideoController) AddPostedVideo(c *fiber.Ctx) error {
	var myform models.Video

	//upload Tumb
	file, errFile := c.FormFile("tumb")
	if errFile != nil {
		fmt.Println("Error File =", errFile)
	}
	var filename string = file.Filename
	if file != nil {

		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/Tumb/%s", filename))
		if errSaveFile != nil {
			fmt.Println("404")
		}
	} else {
		fmt.Println("404")
	}

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/videos")
	}

	//upload Video
	filee, errFilee := c.FormFile("video")
	if errFilee != nil {
		fmt.Println("Error File =", errFilee)
	}
	var filenamee string = filee.Filename
	if filee != nil {
		errSaveFilee := c.SaveFile(filee, fmt.Sprintf("./public/Video/%s", filenamee))
		if errSaveFilee != nil {
			fmt.Println("404")
		}
	} else {
		fmt.Println("404")
	}

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/videos")
	}

	myform.Tumb = filename
	myform.Video = filenamee

	// save product
	err := models.CreateVideo(controller.Db, &myform)
	if err != nil {
		return c.Redirect("/videos")
	}
	// if succeed
	return c.Redirect("/videos")
}

// GET /videos/detailvideo/:id
func (controller *VideoController) GetViewVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var videos models.Video
	err := models.ReadVideoById(controller.Db, &videos, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("viewvideo", fiber.Map{
		"Title":  "Detail Video",
		"Videos": videos,
	})
}

// GET /videos/editvideo/:id
func (controller *VideoController) EditVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var videos models.Video
	err := models.ReadVideoById(controller.Db, &videos, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	return c.Render("editvideo", fiber.Map{
		"Title":  "Edit Video",
		"Videos": videos,
	})
}

// POST /videos/editvideo/:id
func (controller *VideoController) EditPostedVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var videos models.Video
	err := models.ReadVideoById(controller.Db, &videos, idn)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	var myform models.Video

	if err := c.BodyParser(&myform); err != nil {
		return c.Redirect("/videos")
	}

	file, errFile := c.FormFile("tumb")
	if errFile != nil {
		fmt.Println("Error File =", errFile)
	}
	var filename string = file.Filename
	if file != nil {

		errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/Tumb/%s", filename))
		if errSaveFile != nil {
			fmt.Println("404")
		}
	} else {
		fmt.Println("404")
	}

	//upload Video
	filee, errFilee := c.FormFile("video")
	if errFilee != nil {
		fmt.Println("Error File =", errFilee)
	}
	var filenamee string = filee.Filename
	if filee != nil {
		errSaveFilee := c.SaveFile(filee, fmt.Sprintf("./public/Video/%s", filenamee))
		if errSaveFilee != nil {
			fmt.Println("404")
		}
	} else {
		fmt.Println("404")
	}

	myform.Tumb = filename
	myform.Video = filenamee
	videos.Title = myform.Title
	videos.Description = myform.Description
	videos.Video = myform.Video
	videos.Tumb = myform.Tumb
	// save product
	models.UpdateVideo(controller.Db, &videos)
	return c.Redirect("/videos")
}

// GET /videos/deletevideo/:id
func (controller *VideoController) DeleteVideo(c *fiber.Ctx) error {
	id := c.Params("id")
	idn, _ := strconv.Atoi(id)

	var videos models.Video
	models.DeleteVideoById(controller.Db, &videos, idn)
	return c.Redirect("/videos")
}

