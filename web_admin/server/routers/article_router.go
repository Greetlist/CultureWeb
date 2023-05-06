package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/middleware"
    "github.com/gin-gonic/gin"
)

func InitAdminArticleApiRouter(RouterGroup *gin.RouterGroup) {
    AdminUserRouterGroup := RouterGroup.Group("admin")
    //AdminUserRouterGroup.Use(middleware.AdminCookieChecker())
    AdminUserRouterGroup.POST("/saveMedia", controller.SaveMedia)
    AdminUserRouterGroup.POST("/submitArticle", controller.SubmitArticle)
    AdminUserRouterGroup.GET("/getTotalArticle", controller.GetTotalArticle)
    AdminUserRouterGroup.POST("/batchModifyArticle", controller.BatchModifyArticle)
    AdminUserRouterGroup.POST("/batchDeleteArticle", controller.BatchDeleteArticle)

    UserRouterGroup := RouterGroup.Group("user")
    UserNormalRouterGroup := UserRouterGroup.Group("normal")
    UserNormalRouterGroup.POST("/searchArticle", controller.SearchArticle)
}
