package middleware

import (
	"github.com/gin-gonic/gin"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

func CookieChecker() gin.HandlerFunc {
    return func(c *gin.Context) {
        if !checkCookie(c, false) {
            return
        }
        c.Next()
    }
}

func AdminCookieChecker() gin.HandlerFunc {
    return func(c *gin.Context) {
        if !checkCookie(c, true) {
            return
        }
        c.Next()
    }
}

func sendError(c *gin.Context, err *ErrorCode.ResponseError) {
    var res controller.RequestResult
    controller.GenErrorReturn(err, &res)
    c.AbortWithStatusJSON(err.HttpStatusCode, res)
}

func checkCookie(c *gin.Context, checkAdmin bool) bool {
    cookie, e := c.Cookie(config.GlobalConfig.TokenConfig.TokenName)
    if e != nil {
        LOG.Logger.Errorf("get cookie err is: %v", e)
        sendError(c, ErrorCode.GetCookieError)
        return false
    }
    _, err := model.VerifyToken(cookie, checkAdmin)
    if err != nil {
        LOG.Logger.Errorf("verify cookie err is: %v", err)
        sendError(c, err)
        return false
    }
    return true
}
