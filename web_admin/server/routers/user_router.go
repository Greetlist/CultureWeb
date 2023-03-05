package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    "github.com/Greetlist/CultureWeb/web_admin/server/middleware"
    "github.com/gin-gonic/gin"
)

func InitUserApiRouter(RouterGroup *gin.RouterGroup) {
    UserRouterGroup := RouterGroup.Group("user")
    UserRouterGroup.GET("/getUserInfo", middleware.AdminCookieChecker(), controller.GetUserInfo)
    UserRouterGroup.GET("/getTotalUserInfo", middleware.AdminCookieChecker(), controller.GetTotalUserInfo)
    UserRouterGroup.POST("/userRegister", controller.UserRegister)
    UserRouterGroup.POST("/login", controller.UserLogin)
    UserRouterGroup.POST("/logout", controller.UserLogout)
}
