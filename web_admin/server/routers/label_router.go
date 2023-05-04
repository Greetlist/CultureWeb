package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/middleware"
    "github.com/gin-gonic/gin"
)

func InitAdminLabelApiRouter(RouterGroup *gin.RouterGroup) {
    //TODO API
    AdminUserRouterGroup := RouterGroup.Group("admin")
    //AdminUserRouterGroup.Use(middleware.AdminCookieChecker())
    AdminUserRouterGroup.POST("/addSingleLabel", controller.AddSingleLabel)

    UserRouterGroup := RouterGroup.Group("user")
    UserNormalRouterGroup := UserRouterGroup.Group("normal")
    UserNormalRouterGroup.GET("/getTotalLabel", controller.GetTotalLabel)
}
