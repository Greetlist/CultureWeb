package model

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
)

type SubmitArticleRequest struct {
    Title string `json:"title"`
    Rank uint `json:"rank"`
    Summary string `json:"summary"`
    Labels []int `json:"labels"`
    IsTop bool `json:"is_top"`
    Content string `json:"content"`
    Author string `json:"author"`
}
type SubmitArticleResponse struct {
    Result RequestResult `json:"request_result"`
}

type ModifyArticleItem struct {
    ArticleID uint `json:"article_id"`
    Title string `json:"title"`
    Summary string `json:"summary"`
    Rank uint `json:"rank"`
    Labels []int `json:"labels"`
    IsTop bool `json:"is_top"`
    IsEnable bool `json:"is_enable"`
}
type BatchModifyArticleRequest struct {
    ModifyList []ModifyArticleItem `json:"modify_list"`
}
type BatchModifyArticleResponse struct {
    Result RequestResult `json:"request_result"`
}

type BatchDeleteArticleRequest struct {
    DeleteList []int `json:"delete_list"`
}
type BatchDeleteArticleResponse struct {
    Result RequestResult `json:"request_result"`
}

type GetTotalArticleResponse struct {
    ArticleList []schema.ArticleDetail `json:"article_list"`
    Result RequestResult `json:"request_result"`
}
