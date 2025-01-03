package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (con ApiController) Index(c *gin.Context) {
	username, _ := c.Get("username") //中间件与控制器进行传值
	fmt.Println(username)
	v, ok := username.(string)
	if ok {
		c.String(200, "api Index"+v)
	} else {
		c.String(200, "获取用户失败")
	}

}
func (con ApiController) UserList(c *gin.Context) {
	c.String(200, "api Userlist")
}
func (con ApiController) Plist(c *gin.Context) {
	c.SetCookie("username", "张三", 3600, "/", "127.0.0.1", false, false) //设置cookie
	//c.SetCookie("username", "张三", -1, "/", "127.0.0.1", false, false)	//删除cookie
	//c.SetCookie("username", "张三", 3600, "/", ".itying.com", false, false)	//共享cookie
	c.String(200, "api Plist")
}
func (con ApiController) Cart(c *gin.Context) {
	username, _ := c.Cookie("username") //获取cookie
	c.String(200, "api Cart"+username)
}