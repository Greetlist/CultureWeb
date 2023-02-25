package util

//import (
    //"os"
    //"bufio"
    //"glog"
    //"io"
    //"strings"
    //"errors"
//)

//type DataRow []string

//type CsvReader struct {
//    Head []string
//}
//
//func (reader *CsvReader) ReadCsv(fileName string) error {
//    reader.indexMap = make(map[string]int)
//    file, err := os.Open(fileName)
//    if err != nil {
//        glog.Errorf("Open %s Error : %s.", fileName, err)
//        return err
//    }
//    defer file.Close()
//    fileReader := bufio.NewReader(file)
//    var readHead bool = false
//    for {
//        var lineByte []byte
//        var isPrefix bool
//        var err error
//        lineByte, isPrefix, err = fileReader.ReadLine()
//        lineString := string(lineByte)
//        if err == nil {
//            if readHead == false {
//                reader.Head = strings.Split(lineString, ",")
//                if !checkHead(reader.Head) {
//                    return errors.New("Head Has Empty Columns")
//                }
//                readHead = true
//            } else {
//                rawData := strings.Split(lineString, ",")
//            }
//        }
//        if err != nil || isPrefix {
//            if err == io.EOF {
//                break
//            }
//            glog.Errorf("Read Csv File Line Error : %s.", err)
//            return err
//        }
//    }
//    return nil
//}
//
//func checkHead(head []string) bool {
//    for _, v := range(head) {
//        if v == "" {
//            return false
//        }
//    }
//    return true
//}
//
//func (reader *CsvReader) ReadCsv(fileName string, i interface{}) {
//    f, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
//    if err != nil {
//        panic(err)
//    }
//    defer f.Close()
//}



