package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"myapp/component/appctx"
	"myapp/component/uploadprovider"
	"myapp/middleware"
	"os"
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

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3Domain := os.Getenv("S3Domain")
	s3SecretKey := os.Getenv("S3SecretKey")
	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)
	appContext := appctx.NewAppContext(db, s3Provider, secretKey)

	r := gin.Default()
	r.Use(middleware.Recover(appContext))

	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})*/

	r.Static("/static", "./static")

	//POST /restaurants

	v1 := r.Group("/v1", middleware.RequiredAuth(appContext))

	setupRoute(appContext, v1)
	setupAdminRoute(appContext, v1)

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
