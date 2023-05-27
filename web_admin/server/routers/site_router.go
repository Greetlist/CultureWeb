package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    "github.com/Greetlist/CultureWeb/web_admin/server/middleware"
    "github.com/gin-gonic/gin"
)

func InitAdminSiteApiRouter(RouterGroup *gin.RouterGroup) {
    //TODO API
    AdminUserRouterGroup := RouterGroup.Group("admin")
    AdminUserRouterGroup.Use(middleware.AdminCookieChecker())
    AdminUserRouterGroup.POST("/addSingleSite", controller.AddSingleSite)
    AdminUserRouterGroup.POST("/deleteSite", controller.DeleteSite)
    AdminUserRouterGroup.POST("/modifySite", controller.ModifySite)

    UserRouterGroup := RouterGroup.Group("user")
    UserNormalRouterGroup := UserRouterGroup.Group("normal")
    UserNormalRouterGroup.GET("/getTotalSite", controller.GetTotalSite)
}
