package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "api")
		})
		apiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(http.StatusOK, "api/userlist")
		})
		apiRouters.GET("/plist", func(c *gin.Context) {
			c.String(http.StatusOK, "api/plist")
		})
		apiRouters.GET("/cart", func(c *gin.Context) {
			c.String(http.StatusOK, "api/cart")
		})
	}
}
