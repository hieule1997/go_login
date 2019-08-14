package controllers

import (
	container "login_jwt/container"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"login_jwt/models"
	"login_jwt/constants"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
	jwt.StandardClaims
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
	expirationTime := time.Now().Add(config.Config.EXPIRAION_TIME * time.Minute)
	claims := &Claims{
		Id : user.ID,
		Fullname: user.Username,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt :expirationTime.Unix(),
			},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512,claims)
	tokenString ,err := token.SignedString(jwtSecretKey)
	if err !=  nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"msg" : constants.SERVER_ERROR_500,
			})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"msg":"pass login",
		"Token":tokenString,
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
	if err == nil {

		fmt.Println(err)
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
func UpdateUser(c *gin.Context) {
	var postData models.UpdateUserForm
	err := c.BindJSON(&postData)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":constants.SERVER_ERROR_500,
		})
		return
	}
	userID := c.MustGet("userID").(string)
	fmt.Println(postData)
	var user models.UserForm
	user,err = models.UpdateUserById(userID, postData)
	if err  !=  nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"msg":constants.SERVER_ERROR_500,
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"msg":"Update success",
		"result":user,
	})
}

func init() {
	config.Read()
	jwtSecretKey = []byte(config.Config.JWT_SECRET_KEY)
}