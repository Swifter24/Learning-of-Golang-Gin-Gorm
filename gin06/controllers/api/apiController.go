package api

import (
	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "api Index")
}
func (con ApiController) UserList(c *gin.Context) {
	c.String(200, "api Userlist")
}
func (con ApiController) Plist(c *gin.Context) {
	c.String(200, "api Plist")
}
func (con ApiController) Cart(c *gin.Context) {
	c.String(200, "api Cart")
}
