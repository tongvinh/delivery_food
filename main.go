package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"myapp/component/appctx"
	"myapp/middleware"
	"myapp/module/restaurant/transport/ginrestaurant"
	"net/http"
	"os"
	"strconv"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()
	appContext := appctx.NewAppContext(db)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})*/

	//POST /restaurants

	v1 := r.Group("/v1")

	restaurants := v1.Group("/restaurants")

	/*restaurants.POST("", func(c *gin.Context) {
		var data Restaurant

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db.Create(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})*/
	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

	//GET /restaurants /:id = 4
	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data Restaurant

		db.Where("id = ?", id).First(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	//GET /restaurants

	/*restaurants.GET("", func(c *gin.Context) {
		var data []Restaurant

		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}

		var pagingData Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}

		if pagingData.Limit <= 0 {
			pagingData.Limit = 2
		}

		db.Offset((pagingData.Page - 1) * pagingData.Limit).
			Order("id desc").
			Limit(pagingData.Limit).
			Find(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})*/
	restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
	//UPDATE /restaurants /:id
	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db.Where("id = ?", id).Updates(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	//DELETE /restaurants /:id
	/*restaurants.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil)

		c.JSON(http.StatusOK, gin.H{
			"data": 1,
		})
	})*/
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	/*newRestaurant := Restaurant{Name: "Tani", Addr: "9 Pham Van Hai"}

	if err := db.Create(&newRestaurant).Error; err != nil {
		log.Println(err)
	}

	log.Println("New id: ", newRestaurant.Id)*/

	/*var myRestaurant Restaurant

	if err := db.Where("id = ?", 2).First(&myRestaurant).Error; err != nil {
		log.Println(err)
	}

	log.Println(myRestaurant)

	newName := "200Lab"
	updateData := RestaurantUpdate{Name: &newName}

	if err := db.Where("id = ?", 2).Updates(&updateData).Error; err != nil {
		log.Println(err)
	}

	log.Println(myRestaurant)

	if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 3).Delete(nil).Error; err != nil {
		log.Println(err)
	}

	log.Println(myRestaurant)*/

}
