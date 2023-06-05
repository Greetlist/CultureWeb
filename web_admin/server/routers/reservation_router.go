package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    "github.com/Greetlist/CultureWeb/web_admin/server/middleware"
    "github.com/gin-gonic/gin"
)

func InitAdminReservationApiRouter(RouterGroup *gin.RouterGroup) {
    //TODO API
    AdminUserRouterGroup := RouterGroup.Group("admin")
    AdminUserRouterGroup.Use(middleware.AdminCookieChecker())
    AdminUserRouterGroup.POST("/getTotalReservation", controller.GetTotalReservation)
    AdminUserRouterGroup.POST("/submitReservation", controller.SubmitReservation)
    AdminUserRouterGroup.POST("/batchModifyReservation", controller.BatchModifyReservation)
    AdminUserRouterGroup.POST("/cancelReservation", controller.CancelReservation)
}
