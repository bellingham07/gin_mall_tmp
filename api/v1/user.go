package v1

import (
	"gin_mall_tmp/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//相当于controller

func UserRegister(c *gin.Context) {
	var UserRegister service.UserService

	if err := c.ShouldBind(&UserRegister); err == nil {
		res := UserRegister.Register(c.Request.Context()) //获取上下文
		c.JSON(http.StatusOK, res)
	}
}
