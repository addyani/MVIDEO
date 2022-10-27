package controllers //controler akan mengambil model

import (
	"github.com/gofiber/fiber/v2"
)

type VideoController struct { // semacam kelas
	//deklarasi data/variabel

}

func InitVideoController() *VideoController {

	return &VideoController{}
}

func (controller *VideoController) HelloVideo(c *fiber.Ctx) error {
	return c.Render("homevideo", fiber.Map{
		"Title": "ini judul...",
	})

}
