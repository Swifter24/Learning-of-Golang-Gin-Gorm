package routers

import (
	"GoDemo/gin08/controllers/api"
	"GoDemo/gin08/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	apiRouters.Use(middlewares.InitMiddleware) //路由分组定义中间件
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userlist", api.ApiController{}.UserList)
		apiRouters.GET("/plist", api.ApiController{}.Plist)
		apiRouters.GET("/cart", api.ApiController{}.Cart)
	}
}
