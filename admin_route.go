package main

import (
	"github.com/gin-gonic/gin"
	"myapp/component/appctx"
	"myapp/middleware"
	"myapp/module/user/transport/ginuser"
)

func setupAdminRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin",
		middleware.RequiredAuth(appContext),
		middleware.RoleRequired(appContext, "admin", "mod"),
	)

	{
		admin.GET("/profile", ginuser.Profile(appContext))
	}
}
