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
    res.HttpStatusCode = status
    return res
}

var (
    //General Error
    RequestSuccess = NewResponseError(0, "成功", http.StatusOK)
    ParseParamError = NewResponseError(10000, "参数解析失败", http.StatusOK)
    EmptyFileParamError = NewResponseError(10001, "没有上传文件", http.StatusOK)
    UploadFileError = NewResponseError(10002, "上传文件失败", http.StatusOK)

    //DB Error
    DBInsertError = NewResponseError(20000, "数据库插入失败", http.StatusOK)
    DBUpdateError = NewResponseError(20001, "数据库更新失败", http.StatusOK)
    DBDeleteError = NewResponseError(20002, "数据库删除失败", http.StatusOK)
    DBModifyError = NewResponseError(20003, "数据库修改失败", http.StatusOK)
    DBQueryError = NewResponseError(20004, "数据库查询失败", http.StatusOK)

    //Use related Error
    GetUserInfoError = NewResponseError(30000, "获取用户信息失败", http.StatusOK)
    RegisterUserError = NewResponseError(30001, "用户注册失败", http.StatusOK)
    ModifyUserError = NewResponseError(30002, "修改用户信息失败", http.StatusOK)
    DeleteUserError = NewResponseError(30003, "删除用户失败", http.StatusOK)
    ChangePasswordError = NewResponseError(30004, "修改密码失败", http.StatusOK)
    PasswordNotCorrectError = NewResponseError(30005, "用户密码错误", http.StatusOK)

    //Redis Error
    RedisCommandError = NewResponseError(40000, "Redis命令失败", http.StatusOK)
    RedisParseStructError = NewResponseError(40001, "Redis解析Struct失败", http.StatusOK)
    RedisKeyExpireError = NewResponseError(40002, "Redis Key超时", http.StatusOK)

    //cookie Error
    GetCookieError = NewResponseError(50000, "获取Cookie失败,请先登录", http.StatusOK)
    CookieExpireError = NewResponseError(50001, "Cookie 超时,请重新登录", http.StatusOK)
    CheckAdminError = NewResponseError(50002, "检查Admin权限失败", http.StatusOK)
    CookieIsCleanedError = NewResponseError(50003, "Cookie已被清空,请重新登录", http.StatusOK)
)
