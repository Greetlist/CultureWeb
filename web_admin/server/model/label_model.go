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
    query_res := label.DB.Order("label_id").Find(&response.LabelList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.GetLabelError
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
