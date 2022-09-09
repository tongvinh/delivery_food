package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"myapp/common"
	"myapp/component/appctx"
	restaurantbiz "myapp/module/restaurant/biz"
	restaurantmodel "myapp/module/restaurant/model"
	restaurantrepo "myapp/module/restaurant/repository"
	restaurantstorage "myapp/module/restaurant/storage"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.FullFill()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var result []restaurantmodel.Restaurant

		store := restaurantstorage.NewSqlStore(db)
		//likeStore := restaurantlikestorage.NewSQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}

}
