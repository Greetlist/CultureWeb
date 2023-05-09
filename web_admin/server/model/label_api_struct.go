package model
import (
    _ "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
)

type QueryLabelStruct struct {
    LabelID int `json:"label_id"`
    LabelName string `json:"label_name"`
    RelateArticleNumber int64 `json:"article_number"`
}

type GetTotalLabelResponse struct {
    LabelList []QueryLabelStruct `json:"labels"`
    Result RequestResult `json:"request_result"`
}

type AddSingleLabelRequest struct {
    LabelName string `json:"label_name"`
}
type AddSingleLabelResponse struct {
    Result RequestResult `json:"request_result"`
}

type DeleteLabelRequest struct {
    Labels []int `json:"labels"`
}
type DeleteLabelResponse struct {
    Result RequestResult `json:"request_result"`
}

type ModifyLabelItem struct {
    LabelID int `json:"label_id"`
    LabelName string `json:"label_name"`
}
type ModifyLabelRequest struct {
    ModifyList []ModifyLabelItem `json:"modify_list"`
}
type ModifyLabelResponse struct {
    Result RequestResult `json:"request_result"`
}
