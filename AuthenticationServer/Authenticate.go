package AuthenticationServer

import (
	"fmt"

	"net/http"

	"strings"

	"encoding/json"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//Request Header
//Authorization: Basic cm9vdDpwd2Q=  (root:pwd)

type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func CreateTokenEndpoint(c *gin.Context) {
	var user_json User
	if err := c.ShouldBind(&user_json); err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user_json.Username,
			"password": user_json.Password,
		})
		tokenString, _ := token.SignedString([]byte("Secret"))
		fmt.Println(tokenString)

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}

}
func ProtectedEndpoint(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if auth != "" {
		bearerToken := strings.Split(auth, " ")
		if len(bearerToken) == 2 {
			token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method:%v", token.Header["alg"])
				}
				return []byte("Secret"), nil
			})
			if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				var user User
				user.Username = claim["username"].(string)
				user.Password = claim["password"].(string)
				data, _ := json.Marshal(user)
				fmt.Println(string(data))
				c.JSON(http.StatusOK, gin.H{"context": string(data)})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"detail": "Invaild authorization token."})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"detail": "An Authorization Header is required."})

		}
	}
	fmt.Println(auth)

}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6IjEyMyIsInVzZXJuYW1lIjoiemwifQ.v-wceAG654IBYxBka1D_mNyDHeYmP4l6_4IekQ_Bq9I
func StartJWT() {
	r := gin.Default()

	r.POST("/authenticate", CreateTokenEndpoint)
	r.GET("/protected", ProtectedEndpoint)
	r.Run(":80")
}
