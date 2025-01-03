package itying

import (
	"GoDemo/gin07/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	println(models.UnixToTime(1629788418))
	c.HTML(http.StatusOK, "default/user.html", gin.H{
		"msg": "我是一个msg",
		"t":   1629788418,
	})
}
func (con DefaultController) News(c *gin.Context) {
	c.String(200, "News")
}
