package main

import "github.com/gin-gonic/gin"

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	//配置模版文件路径
	r.LoadHTMLGlob("templates/*")
	r.GET("/string", func(c *gin.Context) {
		c.String(200, "value: %v", "index") //c.string
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{ //c.json	map[string]string = gin.H
			"message": "hello json",
			"code":    200,
		})
	})
	r.GET("/json2", func(c *gin.Context) {
		a := &Article{ //struct
			Title:   "json2",
			Desc:    "json2",
			Content: "json2",
		}
		c.JSON(200, a)
	})
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{ //struct
			Title:   "json2",
			Desc:    "json2",
			Content: "json2",
		}
		c.JSONP(200, a) //jsonp	http://127.0.0.1:8080/jsonp?callback=xxxx 	处理跨域问题
	})
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{ //xml
			"success": true,
			"msg":     "i am a xml",
		})
	})
	r.GET("/html", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"msg": "hello html",
		})
	})
	r.Run()
}
