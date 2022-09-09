package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"myapp/common"
	"myapp/component/appctx"
	restaurantbiz "myapp/module/restaurant/biz"
	restaurantstorage "myapp/module/restaurant/storage"
	"net/http"
)

func DeleteRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {

	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		//id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store, requester)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalId())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
