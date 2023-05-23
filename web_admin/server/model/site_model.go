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

type SiteModelStruct struct {
    DB *gorm.DB
}

func (label *SiteModelStruct) GetSiteList(response *GetTotalSiteResponse) *ErrorCode.ResponseError {
    var oriSiteList []schema.Site
    query_res := label.DB.Select([]string{"label_id", "label_name"}).Order("label_id").Find(&oriSiteList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.GetSiteError
    }
    for _, item := range(oriSiteList) {
        num := label.DB.Model(&item).Association("Articles").Count()
        response.SiteList = append(response.SiteList, QuerySiteStruct{SiteID: item.SiteID, SiteName: item.SiteName, RelateArticleNumber: num})
    }
    return nil
}

func (label *SiteModelStruct) AddSingleSite(req *AddSingleSiteRequest) *ErrorCode.ResponseError {
    query_res := label.DB.Create(&schema.Site{SiteName: strings.Trim(req.SiteName, " ")})
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.AddSiteError
    }
    return nil
}

func (label *SiteModelStruct) DeleteSite(labels *[]int) *ErrorCode.ResponseError {
    tx := label.DB.Begin()
    for _, labelID := range(*labels) {
        var label schema.Site
        label.SiteID = labelID
        tx.Model(&label).Association("Articles").Clear()
        tx.Delete(&label)
    }

    if e := tx.Commit().Error; e != nil {
        LOG.Logger.Errorf("DB Error: %v", e)
        tx.Rollback()
        return ErrorCode.DeleteSiteError
    }
    return nil
}

func (label *SiteModelStruct) ModifySite(mList *[]ModifySiteItem) *ErrorCode.ResponseError {
    tx := label.DB.Begin()
    for _, item := range(*mList) {
        var label schema.Site
        label.SiteID = item.SiteID
        update_res := tx.Model(&label).Update("label_name", strings.Trim(item.SiteName, " "))
        if update_res.Error != nil {
            LOG.Logger.Errorf("DB Error: %v", update_res.Error)
            tx.Rollback()
            return ErrorCode.ModifySiteError
        }
    }
    if e := tx.Commit().Error; e != nil {
        LOG.Logger.Errorf("DB Error: %v", e)
        tx.Rollback()
        return ErrorCode.ModifySiteError
    }
    return nil
}
