package model
import (
    _ "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
)

type QuerySiteStruct struct {
    SiteID int `json:"site_id"`
    SiteName string `json:"site_name"`
    Location string `json:"location"`
    PhoneNumber string `json:"phone_number"`
    ContactName string `json:"contact_name"`
}

type GetTotalSiteResponse struct {
    SiteList []QuerySiteStruct `json:"sites"`
    Result RequestResult `json:"request_result"`
}

type AddSingleSiteRequest struct {
    SiteName string `json:"site_name"`
    Location string `json:"location"`
    PhoneNumber string `json:"phone_number"`
    ContactName string `json:"contact_name"`
}
type AddSingleSiteResponse struct {
    Result RequestResult `json:"request_result"`
}

type DeleteSiteRequest struct {
    Sites []int `json:"sites"`
}
type DeleteSiteResponse struct {
    Result RequestResult `json:"request_result"`
}

type ModifySiteItem struct {
    SiteID int `json:"site_id"`
    SiteName string `json:"site_name"`
    Location string `json:"location"`
    PhoneNumber string `json:"phone_number"`
    ContactName string `json:"contact_name"`
}
type ModifySiteRequest struct {
    ModifyList []ModifySiteItem `json:"modify_list"`
}
type ModifySiteResponse struct {
    Result RequestResult `json:"request_result"`
}
