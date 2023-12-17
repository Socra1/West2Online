package api

import (
	"net/http"
	"todo_list/service"

	"github.com/gin-gonic/gin"
)

func UserRegisrter(c *gin.Context) {
	var userregister service.UserService
	//绑定服务对象
	c.Bind(&userregister)
	res := userregister.Register()
	c.JSON(http.StatusOK, gin.H{
		"Status":  res.Status,
		"Message": res.Message,
		"Data":    res.Data,
		"Error":   res.Error,
	})
}

func UserLogin(c *gin.Context) {
	var userlogin service.UserService
	c.Bind(&userlogin)
	res := userlogin.Login()
	c.JSON(http.StatusOK, gin.H{
		"Status":  res.Status,
		"Message": res.Message,
		"Data":    res.Data,
		"Error":   res.Error,
	})
}
