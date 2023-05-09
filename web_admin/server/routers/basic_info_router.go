package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/middleware"
    "github.com/gin-gonic/gin"
)

func InitBasicInfoRouter(RouterGroup *gin.RouterGroup) {
    UserRouterGroup := RouterGroup.Group("user")
    // no cookie needed
    UserNormalRouterGroup := UserRouterGroup.Group("normal")
    UserNormalRouterGroup.GET("/getWebBasicInfo", controller.GetBasicInfo)
}
