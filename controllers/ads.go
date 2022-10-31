package controllers

import (
	"encoding/json"
	"ilmudata/task1/database"
	"ilmudata/task1/models"
	"io/ioutil"

	// "io/ioutil"

	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type AdsController struct {
	// declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitAdsController(s *session.Store) *AdsController {
	db := database.InitDb()
	return &AdsController{Db: db, store: s}
}

type Ads struct {
	Code      int64  `json:"code"`
	Id_Iklan  int64  `json:"idIklan"`
	Id_User   int64  `json:"idUser"`
	ImagePath string `json:"image_path"`
	VideoPath string `json:"video_path"`
}

func (controllers *AdsController) Ads(c *fiber.Ctx) error {
	resd, _ := http.Get("http://127.0.0.1:3000/api/products")
	data, _ := ioutil.ReadAll(resd.Body)
	var api []Ads
	json.Unmarshal(data, &api)
	veri, _ := json.Marshal(api)
	// return c.Render("ads", fiber.Map{
	// 	"message": "Data Ads",
	// 	"data":    veri,
	// })
	formatData := c.Format(veri)
	return c.Render("ads", fiber.Map{
		"message": "Data Iklan",
		"data":    formatData,
	})
}

// login dahulu sebelum melihat tabel view iklan
func (controller *AdsController) GetAds(c *fiber.Ctx) error {
	userId := c.Query("userid")
	user_id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(fiber.Map{
			"message": "Anda harus login",
		})
	}

	var ads []models.Advert
	err = models.ViewAdvert(controller.Db, &ads, user_id)

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(ads)
}

// var BASE_URL = "URL API M IKLAN"

// func (controllers *AdsController) Ads(c *fiber.Ctx) error {
// 	// resd, _ := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
// 	// data, _ := ioutil.ReadAll(resd.Body)
// 	// var api ApiAds
// 	// json.Unmarshal(data, &api)
// 	// veri, _ := json.Marshal(api.Bpi)
// 	// return c.Render("ads", fiber.Map{
// 	// 	"message": "Data Ads",
// 	// 	"Data":    c.Format(veri),
// 	// })
// 	var ad []Ads //untuk menampung postingan dari api
// 	// response, err := http.Get(BASE_URL + "/posts")
// 	response, err := http.Get(BASE_URL + "/ROUTER API")
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	defer response.Body.Close()

// 	decoder := json.NewDecoder(response.Body)
// 	if err := decoder.Decode(&ad); err != nil {
// 		log.Print(err)
// 	}

// 	// update view iklan jadi +1
// 	var data models.Advert
// 	id := c.Params("id") 	// id iklan
// 	idn, _ := strconv.Atoi(id)

// 	err3 := models.ReadAdsById(controllers.Db, &data, idn)
// 	if err3 != nil {
// 		return c.SendStatus(500) // http 500 internal server error
// 	}
// 	data.View++
// 	err = models.UpdateAds(controllers.Db, &data)
// 	if err != nil {
// 		return c.JSON(data)
// 	}

// 	return c.Render("ads", fiber.Map{
// 		"message": "Data Ads",
// 		"data":    ad,
// 	})
// }
