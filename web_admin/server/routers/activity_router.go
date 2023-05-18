package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    "github.com/Greetlist/CultureWeb/web_admin/server/middleware"
    "github.com/gin-gonic/gin"
)

func InitAdminActivityApiRouter(RouterGroup *gin.RouterGroup) {
    AdminUserRouterGroup := RouterGroup.Group("admin")
    AdminUserRouterGroup.Use(middleware.AdminCookieChecker())
    AdminUserRouterGroup.POST("/submitActivity", controller.SubmitActivity)
    AdminUserRouterGroup.GET("/getTotalActivity", controller.GetTotalActivity)
    AdminUserRouterGroup.POST("/batchModifyActivity", controller.BatchModifyActivity)
    AdminUserRouterGroup.POST("/batchDeleteActivity", controller.BatchDeleteActivity)

    UserRouterGroup := RouterGroup.Group("user")
    UserNormalRouterGroup := UserRouterGroup.Group("normal")
    UserNormalRouterGroup.POST("/searchActivity", controller.SearchActivity)
    UserNormalRouterGroup.POST("/getActivityContent", controller.GetActivityContent)
}
