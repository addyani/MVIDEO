package controllers

import (
	"encoding/json"
	"ilmudata/task1/database"
	"ilmudata/task1/models"
	// "io/ioutil"
	"log"
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

// type Time struct {
// 	Updated    string
// 	UpdatedISO string
// 	Updateduk  string
// }

// type Cru struct {
// 	Code       string
// 	Symbol     string
// 	Rate       string
// 	Rate_Float float64
// }

// type Bpi struct {
// 	USD Cru
// 	EUR Cru
// 	GBP Cru
// }

// type ApiAds struct {
// 	Time     Time
// 	CharName string
// 	Bpi      Bpi
// }

// type Post struct {
// 	Id     int64  `json:"id"`
// 	Title  string `json:"title`
// 	Body   string `json:"body"`
// 	UserId int64  `json:"userId"`
// }

type Ads struct {
	Code      int64  `json:"code"`
	Id_Iklan  int64  `json:"idIklan"`
	Id_User   int64  `json:"idUser"`
	ImagePath string `json:"image"`
	VideoPath string `json:"video"`
}

var BASE_URL = "URL API M IKLAN"

func (controllers *AdsController) Ads(c *fiber.Ctx) error {
	// resd, _ := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	// data, _ := ioutil.ReadAll(resd.Body)
	// var api ApiAds
	// json.Unmarshal(data, &api)
	// veri, _ := json.Marshal(api.Bpi)
	// return c.Render("ads", fiber.Map{
	// 	"message": "Data Ads",
	// 	"Data":    c.Format(veri),
	// })
	
	
	var ad []Ads //untuk menampung postingan dari api
	// response, err := http.Get(BASE_URL + "/posts")
	response, err := http.Get(BASE_URL + "/ROUTER API")
	if err != nil {
		log.Print(err)
	}
	defer response.Body.Close()
	
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&ad); err != nil {
		log.Print(err)
	}
	
	// update view iklan jadi +1
	var data models.Advert
	id := c.Params("id") 	// id iklan
	idn, _ := strconv.Atoi(id)
	
	err3 := models.ReadAdsById(controllers.Db, &data, idn)
	if err3 != nil {
		return c.SendStatus(500) // http 500 internal server error
	}
	data.View++ 
	err = models.UpdateAds(controllers.Db, &data)
	if err != nil {
		return c.JSON(data)
	}

	
	return c.Render("ads", fiber.Map{
		"message": "Data Ads",
		"data":    ad,
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