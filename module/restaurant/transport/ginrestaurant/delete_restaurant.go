package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"myapp/common"
	"myapp/component/appctx"
	restaurantbiz "myapp/module/restaurant/biz"
	restaurantstorage "myapp/module/restaurant/storage"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	db := appCtx.GetMainDBConnection()

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
