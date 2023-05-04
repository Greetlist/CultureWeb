package model
import (
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
)

type GetTotalLabelResponse struct {
    LabelList []schema.Label `json:"label_list"`
    Result RequestResult `json:"request_result"`
}

type AddSingleLabelRequest struct {
    LabelName string `json:"label_name"`
}
type AddSingleLabelResponse struct {
    Result RequestResult `json:"request_result"`
}
