package util

import (
    _ "os"
    "testing"
    _ "fmt"
)

func TestCsvReader(t *testing.T) {
    //csvReader := CsvReader{}
    //fileName := "/home/greetlist/macd/data_storage/603788.XSHG/stock_daily_data/kline_daily.csv"
    //if err := csvReader.ReadCsv(fileName); err != nil {
        //os.Exit(1)
    //}

    //recordNum := len(csvReader.Data)
    //fmt.Printf("Total Data is : %v.\n", recordNum)
    //for i := recordNum - 10; i < recordNum; i++ {
        //fmt.Printf("%f.\n", csvReader.Data[i].Open)
    //}
    csvReader := CsvReader{}
    type Stock struct {
        Open float64 `json:"open" csv:"open"`
        Close float64 `json:"close" csv:"close"`
        High float64 `json:"high" csv:"high"`
        Low float64 `json:"low" csv:"low"`
        Money float64 `json:"money" csv:"money"`
        Volume float64 `json:"volume" csv:"volume"`
    }
    var datas Stock
    csvReader.UnmarshalCsv(datas)
}
