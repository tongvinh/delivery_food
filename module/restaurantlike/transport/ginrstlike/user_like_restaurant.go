package ginrstlike

import (
	"github.com/gin-gonic/gin"
	"myapp/common"
	"myapp/component/appctx"
	restaurantstorage "myapp/module/restaurant/storage"
	rstlikebiz "myapp/module/restaurantlike/biz"
	restaurantlikemodel "myapp/module/restaurantlike/model"
	restaurantlikestorage "myapp/module/restaurantlike/storage"
	"net/http"
)

//POST /v1/restaurants/:id/like

func UserLikeRestaurant(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalId()),
			UserId:       requester.GetUserId(),
		}
		store := restaurantlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		incStore := restaurantstorage.NewSqlStore(ctx.GetMainDBConnection())
		biz := rstlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
