package main

import (
	"GoDemo/gin05/routers"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
type Article struct {
	Title   string `xml:"title"`
	Content string `xml:"content"`
}

func main() {
	r := gin.Default()
	//配置模版文件路径
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static/css")

	routers.AdminRoutersInit(r)
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)

	r.Run()
}
