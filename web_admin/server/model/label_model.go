package model

import (
    "gorm.io/gorm"
    //"path"
    //"os"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    //"github.com/Greetlist/CultureWeb/web_admin/server/config"
    //"github.com/Greetlist/CultureWeb/web_admin/server/util"
    //uuid "github.com/nu7hatch/gouuid"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

type LabelModelStruct struct {
    DB *gorm.DB
}

func (label *LabelModelStruct) GetLabelList(response *GetTotalLabelResponse) *ErrorCode.ResponseError {
    var oriLabelList []schema.Label
    query_res := label.DB.Select([]string{"label_id", "label_name"}).Order("label_id").Find(&oriLabelList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.GetLabelError
    }
    for _, item := range(oriLabelList) {
        num := label.DB.Model(&item).Association("Articles").Count()
        response.LabelList = append(response.LabelList, QueryLabelStruct{LabelID: item.LabelID, LabelName: item.LabelName, RelateArticleNumber: num})
    }
    return nil
}

func (label *LabelModelStruct) AddSingleLabel(req *AddSingleLabelRequest) *ErrorCode.ResponseError {
    query_res := label.DB.Create(&schema.Label{LabelName: req.LabelName})
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.AddLabelError
    }
    return nil
}
