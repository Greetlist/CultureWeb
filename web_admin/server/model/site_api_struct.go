package model
import (
    _ "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
)

type QuerySiteStruct struct {
    SiteID int `json:"label_id"`
    SiteName string `json:"label_name"`
    RelateArticleNumber int64 `json:"article_number"`
}

type GetTotalSiteResponse struct {
    SiteList []QuerySiteStruct `json:"labels"`
    Result RequestResult `json:"request_result"`
}

type AddSingleSiteRequest struct {
    SiteName string `json:"label_name"`
}
type AddSingleSiteResponse struct {
    Result RequestResult `json:"request_result"`
}

type DeleteSiteRequest struct {
    Sites []int `json:"labels"`
}
type DeleteSiteResponse struct {
    Result RequestResult `json:"request_result"`
}

type ModifySiteItem struct {
    SiteID int `json:"label_id"`
    SiteName string `json:"label_name"`
}
type ModifySiteRequest struct {
    ModifyList []ModifySiteItem `json:"modify_list"`
}
type ModifySiteResponse struct {
    Result RequestResult `json:"request_result"`
}
