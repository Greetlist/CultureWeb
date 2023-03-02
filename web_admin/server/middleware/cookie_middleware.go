package middleware

import (
	"github.com/gin-gonic/gin"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/redis"
)

func CookieChecker() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
    }
}

func AdminCookieChecker() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
    }
}
