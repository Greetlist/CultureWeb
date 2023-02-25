package api

import (
    "net/http"
    "io/ioutil"
    "strings"
    "fmt"
    "os"
    "path"
    "github.com/gin-gonic/gin"
    "github.com/gocarina/gocsv"
    "greetlist/CultureWeb/server/model"
    "greetlist/CultureWeb/server/conf"
    "greetlist/CultureWeb/server/util"
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
        stockFilePath := path.Join(conf.StockResultBaseDir, stockCode, "result_ma.csv")
        stockDataItem.Records = extractSingleStockDataFromCsv(stockFilePath)
        response.StockDatas = append(response.StockDatas, stockDataItem)
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
    var stockList []string
    fileInfoList, _ := ioutil.ReadDir(conf.StockStorageBaseDir)
    for _, fileItem := range(fileInfoList) {
        if fileItem.IsDir() {
            stockList = append(stockList, fileItem.Name())
        }
    }
    response := model.GetAllStockCodeResponse{StockList: stockList}
    context.JSON(http.StatusOK, response)
}

// GetAllStockCode godoc
// @Summary Query Daily Calc Stock Data
// @Description Return All Daily Calc Stock Data
// @ID getDailyCalcStockData
// @Produce json
// @Success 200 {object} model.GetDailyCalcStockDataResponse
// @Router /api/stock/getDailyCalcStockData [get]
func GetDailyCalcStockData(context *gin.Context) {
    todayStr := strings.ReplaceAll(util.GetTodayStr(), "-", "/")
    dailyFilePath := path.Join(conf.DailyResultBaseDir, todayStr, "/ma_statistic_result.csv")

    type result struct {
        StockCode string `csv:"StockCode"`
        CheckFlag bool `csv:"CheckFlag"`
    }
    f, err := os.OpenFile(dailyFilePath, os.O_RDWR, os.ModePerm)
    if err != nil {
        fmt.Printf("Open File %s Error : %s.\n", dailyFilePath, err)
    }
    defer f.Close()
    var resultList []*result
    if err := gocsv.UnmarshalFile(f, &resultList); err != nil {
        fmt.Printf("Unmarshal csv file error : %s.\n", err)
	}
    var dailyStockCodeList []string
    for _, value := range(resultList) {
        if (*value).CheckFlag {
            dailyStockCodeList = append(dailyStockCodeList, (*value).StockCode)
        }
    }

    var stockDatas []model.StockDataItem
    for _, stockCode := range(dailyStockCodeList) {
        var curItem model.StockDataItem
        curItem.StockCode = stockCode
        stockCsvFilePath := path.Join(conf.StockResultBaseDir, stockCode, "/result_ma.csv")
        curItem.Records = extractSingleStockDataFromCsv(stockCsvFilePath)
        stockDatas = append(stockDatas, curItem)
    }

    response := model.GetDailyCalcStockDataResponse{stockDatas}
    context.JSON(http.StatusOK, response)
}

func extractSingleStockDataFromCsv(stockFilePath string) []model.SingleRecord {
    var res []model.SingleRecord
    f, err := os.OpenFile(stockFilePath, os.O_RDWR, os.ModePerm)
    if err != nil {
        fmt.Printf("Open File %s Error : %s.\n", stockFilePath, err)
    }
    defer f.Close()

    var curStockRecords []*model.SingleRecord
    if err := gocsv.UnmarshalFile(f, &curStockRecords); err != nil {
        fmt.Printf("Unmarshal csv file error : %s.\n", err)
    }
    dataLen := len(curStockRecords)
    startIndex := dataLen - dataLen / 7
    curStockRecords = curStockRecords[startIndex:]
    for _, value := range(curStockRecords) {
        res = append(res, *value)
    }
    return res
}

