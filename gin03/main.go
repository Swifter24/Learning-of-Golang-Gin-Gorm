package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// 时间戳转换为日期
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func Println(str1, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}
func main() {
	r := gin.Default()
	//自定义模版函数	放在加载模版之前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"Println":    Println,
	})
	//配置模版文件路径
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static/css")

	r.GET("/html", func(c *gin.Context) {
		c.HTML(200, "itying/index.html", gin.H{
			"msg":    "hello html",
			"score":  70,
			"hobby":  []string{"吃饭", "睡觉", "写代码"},
			"hobby1": []string{},
			"news": &Article{
				Title:   "111",
				Content: "222",
			},
			"date": 212445252,
		})
	})
	r.GET("/news", func(c *gin.Context) {
		a := &Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		}
		c.HTML(200, "itying/news.html", gin.H{
			"title": "新闻页面",
			"news":  a,
		})
	})

	//后台
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin/index.html", gin.H{
			"msg":   "hello html admin",
			"title": "admin",
		})
	})
	r.GET("admin/news", func(c *gin.Context) {
		a := &Article{
			Title:   "新闻标题admin",
			Content: "新闻内容admin",
		}
		c.HTML(200, "admin/news.html", gin.H{
			"title": "新闻页面admin",
			"news":  a,
		})
	})
	r.Run()
}
