package ginuser

import (
	"github.com/gin-gonic/gin"
	"myapp/common"
	"myapp/component/appctx"
	"myapp/component/hasher"
	"myapp/component/tokenprovider/jwt"
	userbiz "myapp/module/user/biz"
	usermodel "myapp/module/user/model"
	userstore "myapp/module/user/store"
	"net/http"
)

func Login(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := ctx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(ctx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
