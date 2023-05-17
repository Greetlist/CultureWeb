package model

import (
    err "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

type RequestResult struct {
    ReturnCode int `json:"return_code"`
    ErrorMsg string `json:"error_msg"`
}

type GeneralResponse struct {
    Result RequestResult `json:"request_result"`
}

func GenErrorReturn(err *err.ResponseError, res *RequestResult) {
    res.ReturnCode = err.Code
    res.ErrorMsg = err.Message
}

func GenSuccessReturn(res *RequestResult) {
    res.ReturnCode = err.RequestSuccess.Code
    res.ErrorMsg = err.RequestSuccess.Message
}
