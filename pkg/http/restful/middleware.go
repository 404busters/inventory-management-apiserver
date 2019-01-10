package restful

import (
	"os"

	"github.com/gin-gonic/gin"
)

var authMiddleware gin.HandlerFunc

func init() {
	accounts := gin.Accounts{os.Getenv("AUTH_BASIC_USERNAME"): os.Getenv("AUTH_BASIC_PASSWORD")}
	authMiddleware = gin.BasicAuth(accounts)
}
