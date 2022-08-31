package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"myapp/common"
	"myapp/component/appctx"
	restaurantbiz "myapp/module/restaurant/biz"
	restaurantmodel "myapp/module/restaurant/model"
	restaurantstorage "myapp/module/restaurant/storage"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			/*c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})*/
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
