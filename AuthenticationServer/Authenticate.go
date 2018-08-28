package AuthenticationServer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Request Header
//Authorization: Basic cm9vdDpwd2Q=  (root:pwd)
func StartJWT() {
	r := gin.Default()

	r.GET("/", gin.BasicAuth(gin.Accounts{"root": "pwd"}), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"user": "root"})
	})
	r.Run(":80")
}
