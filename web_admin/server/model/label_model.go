package model

import (
    "gorm.io/gorm"
    "strings"
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
    query_res := label.DB.Create(&schema.Label{LabelName: strings.Trim(req.LabelName, " ")})
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.AddLabelError
    }
    return nil
}

func (label *LabelModelStruct) DeleteLabel(labels *[]int) *ErrorCode.ResponseError {
    tx := label.DB.Begin()
    for _, labelID := range(*labels) {
        var label schema.Label
        label.LabelID = labelID
        tx.Model(&label).Association("Articles").Clear()
        tx.Delete(&label)
    }

    if e := tx.Commit().Error; e != nil {
        LOG.Logger.Errorf("DB Error: %v", e)
        tx.Rollback()
        return ErrorCode.DeleteLabelError
    }
    return nil
}

func (label *LabelModelStruct) ModifyLabel(mList *[]ModifyLabelItem) *ErrorCode.ResponseError {
    tx := label.DB.Begin()
    for _, item := range(*mList) {
        var label schema.Label
        label.LabelID = item.LabelID
        update_res := tx.Model(&label).Update("label_name", strings.Trim(item.LabelName, " "))
        if update_res.Error != nil {
            LOG.Logger.Errorf("DB Error: %v", update_res.Error)
            tx.Rollback()
            return ErrorCode.ModifyLabelError
        }
    }
    if e := tx.Commit().Error; e != nil {
        LOG.Logger.Errorf("DB Error: %v", e)
        tx.Rollback()
        return ErrorCode.ModifyLabelError
    }
    return nil
}
