package controllers

import (
	container "login_jwt/container"
	"github.com/gin-gonic/gin"
	"net/http"
	"login_jwt/models"
	"login_jwt/constants"
	"fmt"
)

var jwtSecretKey = []byte("")
var config = container.NewContainer()

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	// jwt.StandardClaims
}

func Singin(c *gin.Context){
	var loginData Login
	err := c.BindJSON(& loginData)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"Some time when error",
		})
		return
	}
	user,err := models.FindByUserName(loginData.Username)
	if err != nil || models.CheckPassword(loginData.Password,user.Password) == false {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"Sai rồi nhé đéo nói mật khẩu đâu",
		})
		return
	}
	c.JSON(200,gin.H{
		"msg":"pass login",
	})
	return
}
func Register (c * gin.Context){
	var postData models.UserForm
	err := c.BindJSON(&postData)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":constants.SERVER_ERROR_500,
		})
		return
	}
	_,err = models.FindByUserName(postData.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"user readly exist",
		})
		return
	}
	models.Create(postData)
	c.JSON(http.StatusCreated,gin.H{
		"msg":"User create success",
	})
}