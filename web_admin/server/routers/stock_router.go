package routers

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/controller"
    "github.com/gin-gonic/gin"
)

func InitStockApiRouter(RouterGroup *gin.RouterGroup) {
    StockRouterGroup := RouterGroup.Group("stock")
    StockRouterGroup.GET("/getDailyCalcStockData", controller.GetDailyCalcStockData)
    StockRouterGroup.POST("/getQueryStockData", controller.GetQueryStockData)
    StockRouterGroup.GET("/getAllStockCode", controller.GetAllStockCode)
}
