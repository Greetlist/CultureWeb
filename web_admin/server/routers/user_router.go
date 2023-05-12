package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    "github.com/Greetlist/CultureWeb/web_admin/server/middleware"
    "github.com/gin-gonic/gin"
)

func InitAdminUserApiRouter(RouterGroup *gin.RouterGroup) {
    AdminUserRouterGroup := RouterGroup.Group("admin")
    AdminUserRouterGroup.Use(middleware.AdminCookieChecker())
    AdminUserRouterGroup.GET("/getUserInfo", controller.GetUserInfo)
    AdminUserRouterGroup.GET("/getTotalUserInfo", controller.GetTotalUserInfo)
    AdminUserRouterGroup.POST("/addUser", controller.AddUser)
}

func InitNormalUserApiRouter(RouterGroup *gin.RouterGroup) {
    UserRouterGroup := RouterGroup.Group("user")

    // no cookie needed
    UserNormalRouterGroup := UserRouterGroup.Group("normal")
    UserNormalRouterGroup.POST("/userRegister", controller.UserRegister)
    UserNormalRouterGroup.POST("/login", controller.UserLogin)
    UserNormalRouterGroup.POST("/logout", controller.UserLogout)
    UserNormalRouterGroup.GET("/checkLogin", controller.CheckLogin)

    // need cookie
    LoginUserRouterGroup := UserRouterGroup.Group("auth")
    LoginUserRouterGroup.Use(middleware.CookieChecker())
    LoginUserRouterGroup.POST("/userModify", controller.UserModify)
    LoginUserRouterGroup.POST("/changePassword", controller.ChangePassword)
}
