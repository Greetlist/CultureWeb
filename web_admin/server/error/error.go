package error

import (
    "net/http"
)

type ResponseError struct {
    Code int
    Message string
    HttpStatusCode int
}

func (r *ResponseError) Error() string {
    return r.Message
}

func NewResponseError(code int, msg string, status int) *ResponseError {
    res := new(ResponseError)
    res.Code = code
    res.Message = msg
    res.HttpStatusCode = code
    return res
}

var (
    //General Error
    RequestSuccess = NewResponseError(0, "成功", http.StatusOK)
    ParseParamError = NewResponseError(10000, "参数解析失败", http.StatusOK)

    //DB Error
    DBInsertError = NewResponseError(20000, "数据库插入失败", http.StatusOK)
    DBUpdateError = NewResponseError(20001, "数据库更新失败", http.StatusOK)
    DBDeleteError = NewResponseError(20002, "数据库删除失败", http.StatusOK)
    DBModifyError = NewResponseError(20003, "数据库修改失败", http.StatusOK)

    //Use related Error
    RegisterUserError = NewResponseError(30000, "注册用户失败", http.StatusOK)
)
