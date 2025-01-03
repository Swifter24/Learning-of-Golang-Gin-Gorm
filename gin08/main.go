package main

import (
	"GoDemo/gin08/models"
	"GoDemo/gin08/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

type Person struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
type Article struct {
	Title   string `xml:"title"`
	Content string `xml:"content"`
}

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1、我是一个中间件")
	c.Next() //调用该请求的剩余处理程序
	//c.Abort()//终止调用该请求的剩余处理程序
	fmt.Println("2、我是一个中间件")
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}
func initMiddlewareTwo(c *gin.Context) {
	fmt.Println("1、我是一个中间件two")
}
func main() {
	r := gin.Default()
	//自定义模版函数	放在加载模版之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	//r:=gin.New()	//不自带中间件的引擎
	//配置模版文件路径
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static/css")

	//r.Use(initMiddleware, initMiddlewareTwo) //全局中间件
	routers.AdminRoutersInit(r)
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)

	r.Run()
}
