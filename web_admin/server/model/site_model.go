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

func (site *SiteModelStruct) GetSiteList(response *GetTotalSiteResponse) *ErrorCode.ResponseError {
    var oriSiteList []schema.Site
    query_res := site.DB.Order("site_id").Find(&oriSiteList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.GetSiteError
    }
    for _, item := range(oriSiteList) {
        LOG.Logger.Infof("%v", item)
        response.SiteList = append(response.SiteList, QuerySiteStruct{SiteID: item.SiteID, SiteName: item.SiteName, Location: item.Location, PhoneNumber: item.PhoneNumber, ContactName: item.ContactName})
    }
    return nil
}

func (site *SiteModelStruct) AddSingleSite(req *AddSingleSiteRequest) *ErrorCode.ResponseError {
    query_res := site.DB.Create(
        &schema.Site{
            SiteName: strings.Trim(req.SiteName, " "),
            Location: strings.Trim(req.Location, " "),
            PhoneNumber: strings.Trim(req.PhoneNumber, " "),
            ContactName: strings.Trim(req.ContactName, " ")})
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.AddSiteError
    }
    return nil
}

func (site *SiteModelStruct) DeleteSite(sites *[]int) *ErrorCode.ResponseError {
    tx := site.DB.Begin()
    for _, siteID := range(*sites) {
        var site schema.Site
        site.SiteID = siteID
        tx.Model(&site).Association("Articles").Clear()
        tx.Delete(&site)
    }

    if e := tx.Commit().Error; e != nil {
        LOG.Logger.Errorf("DB Error: %v", e)
        tx.Rollback()
        return ErrorCode.DeleteSiteError
    }
    return nil
}

func (site *SiteModelStruct) ModifySite(mList *[]ModifySiteItem) *ErrorCode.ResponseError {
    tx := site.DB.Begin()
    for _, item := range(*mList) {
        var site schema.Site
        site.SiteID = item.SiteID
        update_res := tx.Model(&site).Updates(
            schema.Site{SiteName: strings.Trim(item.SiteName, " "), Location: strings.Trim(item.Location, " "), PhoneNumber: strings.Trim(item.PhoneNumber, " "), ContactName: strings.Trim(item.ContactName, " ")})
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
