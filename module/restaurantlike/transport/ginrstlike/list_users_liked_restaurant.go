package ginrstlike

import (
	"github.com/gin-gonic/gin"
	"myapp/common"
	"myapp/component/appctx"
	rstlikebiz "myapp/module/restaurantlike/biz"
	restaurantlikemodel "myapp/module/restaurantlike/model"
	restaurantlikestorage "myapp/module/restaurantlike/storage"
	"net/http"
)

func ListUser(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		//var filter restaurantmodel.Filter

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalId()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.FullFill()

		store := restaurantlikestorage.NewSQLStore(ctx.GetMainDBConnection())
		biz := rstlikebiz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
