package restful

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var authMiddleware gin.HandlerFunc = func(c *gin.Context) {
	c.Next()
}

func init() {
	username := os.Getenv("AUTH_BASIC_USERNAME")
	password := os.Getenv("AUTH_BASIC_PASSWORD")
	if username == "" || password == "" {
		logrus.Warnln("Username and Password is not config, please set environment variable AUTH_BASIC_USERNAME and AUTH_BASIC_PASSWORD")
		return
	}
	accounts := gin.Accounts{username: password}
	authMiddleware = gin.BasicAuth(accounts)
}
