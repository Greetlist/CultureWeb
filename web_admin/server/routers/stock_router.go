package routers

import (
    "greetlist/CultureWeb/server/api"
    "github.com/gin-gonic/gin"
)

func InitStockApiRouter(RouterGroup *gin.RouterGroup) {
    StockRouterGroup := RouterGroup.Group("stock")
    StockRouterGroup.GET("/getDailyCalcStockData", api.GetDailyCalcStockData)
    StockRouterGroup.POST("/getQueryStockData", api.GetQueryStockData)
    StockRouterGroup.GET("/getAllStockCode", api.GetAllStockCode)
}
