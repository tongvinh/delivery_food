package ginrstlike

import (
	"github.com/gin-gonic/gin"
	"myapp/common"
	"myapp/component/appctx"
	restaurantstorage "myapp/module/restaurant/storage"
	rstlikebiz "myapp/module/restaurantlike/biz"
	restaurantlikestorage "myapp/module/restaurantlike/storage"
	"net/http"
)

//DELETE /v1/restaurants/:id/unlike

func UserDislikeRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		decStore := restaurantstorage.NewSqlStore(ctx.GetMainDBConnection())
		biz := rstlikebiz.NewUserDislikeRestaurantBiz(store, decStore)

		if err := biz.DislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalId())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
