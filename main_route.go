package main

import (
	"github.com/gin-gonic/gin"
	"myapp/component/appctx"
	"myapp/middleware"
	"myapp/module/restaurant/transport/ginrestaurant"
	"myapp/module/restaurantlike/transport/ginrstlike"
	"myapp/module/upload/uploadtransport/ginupload"
	"myapp/module/user/transport/ginuser"
	"net/http"
	"strconv"
)

func setupRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {

	v1.POST("/upload", ginupload.Upload(appContext))

	v1.POST("/register", ginuser.Register(appContext))

	v1.POST("/authenticate", ginuser.Login(appContext))

	v1.GET("/profile", middleware.RequiredAuth(appContext), ginuser.Profile(appContext))

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

		appContext.GetMainDBConnection().Where("id = ?", id).First(&data)

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

		appContext.GetMainDBConnection().Where("id = ?", id).Updates(&data)

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

	restaurants.POST("/:id/liked-users", ginrstlike.UserLikeRestaurant(appContext))
	restaurants.DELETE("/:id/liked-users", ginrstlike.UserDislikeRestaurant(appContext))
	restaurants.GET("/:id/liked-users", ginrstlike.ListUser(appContext))
	//	GET /v1/restaurants/:id/liked-users
}
