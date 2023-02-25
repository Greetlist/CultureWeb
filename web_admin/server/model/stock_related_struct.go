package model

type GetQueryStockDataRequest struct {
    StockList []string `json:"stock_list" example:['000001', '002142', ...] binding:"required"`
}

type SingleRecord struct {
    Date string `json:"date" csv:"date"`
    Open float64 `json:"open" csv:"open"`
    Close float64 `json:"close" csv:"close"`
    High float64 `json:"high" csv:"high"`
    Low float64 `json:"low" csv:"low"`
    Volume float64 `json:"volume" csv:"volume"`
    Money float64 `json:"money" csv:"money"`
    Ma13 float64 `json:"ma13" csv:"MA13"`
    Ma34 float64 `json:"ma34" csv:"MA34"`
    Ma55 float64 `json:"ma55" csv:"MA55"`
}

type StockDataItem struct {
    StockCode string `json:"stock_code" example:"002142"`
    Records []SingleRecord `json:"records"`
}

type GetQueryStockDataResponse struct {
    StockDatas []StockDataItem `json:"stock_datas"`
}

type GetAllStockCodeResponse struct {
    StockList []string `json:"stock_list" example:"['000001', '002142', ...]"`
}

type GetDailyCalcStockDataResponse struct {
    StockDatas []StockDataItem `json:"stock_datas"`
}
