package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
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
	//r.Static("/static", "./static/css")
	//GET传值
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")
		c.JSON(200, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "itying/user.html", gin.H{})
	})

	r.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")
		c.JSON(200, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	r.GET("/getUser", func(c *gin.Context) {
		user := &Person{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(200, user)
		} else {
			c.JSON(200, gin.H{
				"err": err.Error(),
			})
		}
	})
	r.POST("/doAddUser2", func(c *gin.Context) {
		user := &Person{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(200, user)
		} else {
			c.JSON(200, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.POST("/xml", func(c *gin.Context) {
		article := &Article{}
		xmlSliceDate, _ := c.GetRawData()
		if err := xml.Unmarshal(xmlSliceDate, article); err == nil {
			c.JSON(http.StatusOK, article)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})
	//动态路由
	r.GET("/user/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		c.String(http.StatusOK, uid)
	})
	r.Run()
}
