package model
import (
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
)

type GetTotalLabelResponse struct {
    Labels []schema.Label `json:"labels"`
    Result RequestResult `json:"request_result"`
}
