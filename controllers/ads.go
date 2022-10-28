package controllers

import (
	"encoding/json"
	"ilmudata/task1/database"
	"io/ioutil"

	"net/http"

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
	Id_Iklan  int64  `json:"id_iklan"`
	Id_User   int64  `json:"id_user"`
	ImagePath string `json:"image_path"`
	VideoPath string `json:"video_path"`
}

type GetExample struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// func (controllers *AdsController) GetAds(c *fiber.Ctx) error {
// 	var url string
// 	// var id = c.Query("id")

// 	if url != "" {
// 		url = "http://localhost:3000/api/products"
// 	} else {
// 		return c.JSON(fiber.Map{
// 			"message": "Error 500 Id Not Found",
// 		})
// 	}

// 	client := &http.Client{}
// 	reques, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return c.SendString("Error Id Not Found")
// 	}
// 	reques.Header.Add("Accept", "application/json")
// 	reques.Header.Add("Content-Type", "application/json")
// 	resp, err := client.Do(reques)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	defer resp.Body.Close()

// 	bodyBytes, err := ioutil.ReadAll(reques.Body)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	var responseObject Ads
// 	json.Unmarshal(bodyBytes, &responseObject)

// 	return c.JSON(responseObject)
// }

// var BASE_URL = ""

func (controllers *AdsController) Ads(c *fiber.Ctx) error {
	resd, _ := http.Get("http://127.0.0.1:3000/api/products")
	data, _ := ioutil.ReadAll(resd.Body)
	var api []GetExample
	json.Unmarshal(data, &api)
	veri, _ := json.Marshal(api)
	return c.Format(veri)
	// return c.Render("ads", fiber.Map{
	// 	"message": "Data Ads",
	// 	"Data":    c.Format(veri),
	// })

	// 	var ad []Ads //untuk menampung postingan dari api
	// 	// response, err := http.Get(BASE_URL + "/posts")
	// 	response, err := http.Get(BASE_URL + "/posts")
	// 	if err != nil {
	// 		log.Print(err)
	// 	}
	// 	defer response.Body.Close()

	// 	decoder := json.NewDecoder(response.Body)
	// 	if err := decoder.Decode(&ad); err != nil {
	// 		log.Print(err)
	// 	}

	//	return c.Render("ads", fiber.Map{
	//		"message": "Data Ads",
	//		"data":    ad,
	//	})
}
