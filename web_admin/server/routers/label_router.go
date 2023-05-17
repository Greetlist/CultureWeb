package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    "github.com/Greetlist/CultureWeb/web_admin/server/middleware"
    "github.com/gin-gonic/gin"
)

func InitAdminLabelApiRouter(RouterGroup *gin.RouterGroup) {
    //TODO API
    AdminUserRouterGroup := RouterGroup.Group("admin")
    AdminUserRouterGroup.Use(middleware.AdminCookieChecker())
    AdminUserRouterGroup.POST("/addSingleLabel", controller.AddSingleLabel)
    AdminUserRouterGroup.POST("/deleteLabel", controller.DeleteLabel)
    AdminUserRouterGroup.POST("/modifyLabel", controller.ModifyLabel)

    UserRouterGroup := RouterGroup.Group("user")
    UserNormalRouterGroup := UserRouterGroup.Group("normal")
    UserNormalRouterGroup.GET("/getTotalLabel", controller.GetTotalLabel)
}
