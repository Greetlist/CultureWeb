package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    "github.com/gin-gonic/gin"
)

func InitUserApiRouter(RouterGroup *gin.RouterGroup) {
    UserRouterGroup := RouterGroup.Group("user")
    UserRouterGroup.GET("/getUserInfo", controller.GetUserInfo)
    UserRouterGroup.GET("/getTotalUserInfo", controller.GetTotalUserInfo)
    UserRouterGroup.POST("/userRegister", controller.UserRegister)
    UserRouterGroup.GET("/userLogin", controller.UserLogin)
    UserRouterGroup.GET("/userLogout", controller.UserLogout)
}
