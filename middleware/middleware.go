package middleware


import (
	"github.com/dgrijalva/jwt-go"
	container "login_jwt/container"
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	"login_jwt/controllers"
	"login_jwt/models"
	"log"
)

var jwtSecretKey = []byte("")
var config = container.NewContainer()

func Jwt_required() gin.HandlerFunc{
	return func (c *gin.Context){
		header:= c.Request.Header.Get("Authorization")
		token := ""
		arr := strings.Split(header, " ")
		if len(arr) > 1 {
			token = arr[1]
		} else {
			token = header
		}
		if len(token) < 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "API token required",
			})
			c.Abort()
			return
		}
		claims := &controllers.Claims{}
		parseToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})
		if !parseToken.Valid || err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Token is valid",
			})
			c.Abort()
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Internal server error",
			})
			c.Abort()
			return
		}
		var user models.UserForm
		user, err = models.FindById(claims.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "User not exist",
			})
			c.Abort()
			return
		}
		log.Println(user.Role.Code)
		c.Set("token", token)
		c.Set("claims", claims)
		c.Set("userID", claims.Id)
		c.Set("role", user.Role.Code)
		c.Next()

	}
}


func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}


func init() {
	config.Read()
	jwtSecretKey = []byte(config.Config.JWT_SECRET_KEY)
}