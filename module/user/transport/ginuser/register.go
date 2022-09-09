package ginuser

import (
	"github.com/gin-gonic/gin"
	"myapp/common"
	"myapp/component/appctx"
	"myapp/component/hasher"
	userbiz "myapp/module/user/biz"
	usermodel "myapp/module/user/model"
	userstore "myapp/module/user/store"
	"net/http"
)

func Register(ctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
