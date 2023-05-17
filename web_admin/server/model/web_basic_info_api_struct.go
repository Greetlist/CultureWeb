package model
import (
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
)

type GetWebBasicInfoResponse struct {
    WebBasicInfo config.WebBasicInfo `json:"web_basic_info"`
    Result RequestResult `json:"request_result"`
}
