package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
)

// GetQueryStockData godoc
// @Summary Query Sepcific Stock Computed Data
// @Description Return Strategy Sepcific Computed Data
// @ID getQueryStockData
// @Accept json
// @Produce json
// @Param request_json body model.GetQueryStockDataRequest true "Query Stock List"
// @Success 200 {object} model.GetQueryStockDataResponse
// @Router /api/stock/getQueryStockData [post]
func GetQueryStockData(context *gin.Context) {
    request := model.GetQueryStockDataRequest{}
    if err := context.BindJSON(&request); err != nil {
	    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    var response model.GetQueryStockDataResponse
    for _, stockCode := range(request.StockList) {
        var stockDataItem model.StockDataItem
        stockDataItem.StockCode = stockCode
    }
    context.JSON(http.StatusOK, response)
}

// GetAllStockCode godoc
// @Summary Query All Stock Code
// @Description Return All Stock Code
// @ID getAllStockCode
// @Produce json
// @Success 200 {object} model.GetAllStockCodeResponse
// @Router /api/stock/getAllStockCode [get]
func GetAllStockCode(context *gin.Context) {
    LOG.LOG.Info("Test Info")
    LOG.LOG.Error("Test ERROR")
    LOG.LOG.Alert("Test ALERT")
    LOG.LOG.Critical("Test critical")
    LOG.LOG.Warning("Test WARNING")
}

// GetAllStockCode godoc
// @Summary Query Daily Calc Stock Data
// @Description Return All Daily Calc Stock Data
// @ID getDailyCalcStockData
// @Produce json
// @Success 200 {object} model.GetDailyCalcStockDataResponse
// @Router /api/stock/getDailyCalcStockData [get]
func GetDailyCalcStockData(context *gin.Context) {
}
