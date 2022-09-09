package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"myapp/common"
	"myapp/component/appctx"
)

func RoleRequired(ctx appctx.AppContext, allowRoles ...string) gin.HandlerFunc {

	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, item := range allowRoles {
			if u.GetRole() == item {
				hasFound = true
				break
			}
		}
		if !hasFound {
			panic(common.ErrNoPermission(errors.New("invalid role user")))
		}
		c.Next()
	}
}
