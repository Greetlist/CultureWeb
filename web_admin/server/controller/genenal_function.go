package controller

import (
    err "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

func GenErrorReturn(err *err.ResponseError, res *RequestResult) {
    res.ReturnCode = err.Code
    res.ErrorMsg = err.Message
}

func GenSuccessReturn(res *RequestResult) {
    res.ReturnCode = err.RequestSuccess.Code
    res.ErrorMsg = err.RequestSuccess.Message
}
