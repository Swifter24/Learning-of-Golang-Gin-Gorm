package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	c.String(200, "用户列表")
	con.success(c)
}
func (con UserController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}
func (con UserController) DoUpload(c *gin.Context) {
	username := c.PostForm("username")
	file, _ := c.FormFile("face")
	dst := path.Join("./templates", file.Filename)
	c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"success":  true,
	})
	c.String(http.StatusOK, "执行上传")
}
func (con UserController) Edit(c *gin.Context) {
	c.String(200, "用户列表-edit")
	con.success(c)
}
