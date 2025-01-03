package routers

import (
	"GoDemo/gin08/controllers/admin"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1、我是一个中间件")
	c.Next() //调用该请求的剩余处理程序
	//c.Abort()//终止调用该请求的剩余处理程序
	fmt.Println("2、我是一个中间件")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}
func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", initMiddleware, func(c *gin.Context) {
			fmt.Println("这是一个首页")
			time.Sleep(time.Second * 5)
			c.String(http.StatusOK, "后台首页")
		})
		adminRouters.GET("/user", admin.IndexController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.POST("/user/doUpload", admin.UserController{}.DoUpload)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
	}
}
