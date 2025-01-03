package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println(time.Now())

	c.Set("username", "张三")

	//中间件中使用goroutine，不能使用上下文，必须使用其只读副本（c.Copy()）
	cPc := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Path" + cPc.Request.URL.Path)
	}()
}
