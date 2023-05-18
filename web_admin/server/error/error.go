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
    ParseParamError = NewResponseError(1000, "参数解析失败", http.StatusOK)
    EmptyFileParamError = NewResponseError(1001, "没有上传文件", http.StatusOK)
    UploadFileError = NewResponseError(1002, "上传文件失败", http.StatusOK)
    CreateFileError = NewResponseError(1003, "创建文件失败", http.StatusOK)
    OpenFileError = NewResponseError(1004, "打开文件失败", http.StatusOK)
    WriteFileError = NewResponseError(1005, "写文件失败", http.StatusOK)
    ReadFileError = NewResponseError(1006, "读文件失败", http.StatusOK)

    //DB Error
    DBInsertError = NewResponseError(2000, "数据库插入失败", http.StatusOK)
    DBUpdateError = NewResponseError(2001, "数据库更新失败", http.StatusOK)
    DBDeleteError = NewResponseError(2002, "数据库删除失败", http.StatusOK)
    DBModifyError = NewResponseError(2003, "数据库修改失败", http.StatusOK)
    DBQueryError = NewResponseError(2004, "数据库查询失败", http.StatusOK)

    //Use related Error
    GetUserInfoError = NewResponseError(3000, "获取用户信息失败", http.StatusOK)
    RegisterUserError = NewResponseError(3001, "用户注册失败", http.StatusOK)
    ModifyUserError = NewResponseError(3002, "修改用户信息失败", http.StatusOK)
    DeleteUserError = NewResponseError(3003, "删除用户失败", http.StatusOK)
    ChangePasswordError = NewResponseError(3004, "修改密码失败", http.StatusOK)
    PasswordNotCorrectError = NewResponseError(3005, "用户密码错误", http.StatusOK)

    //Redis Error
    RedisCommandError = NewResponseError(4000, "Redis命令失败", http.StatusOK)
    RedisParseStructError = NewResponseError(4001, "Redis解析Struct失败", http.StatusOK)
    RedisKeyExpireError = NewResponseError(4002, "Redis Key超时", http.StatusOK)

    //cookie Error
    GetCookieError = NewResponseError(5000, "获取Cookie失败,请先登录", http.StatusOK)
    CookieExpireError = NewResponseError(5001, "Cookie 超时,请重新登录", http.StatusOK)
    CheckAdminError = NewResponseError(5002, "检查Admin权限失败", http.StatusOK)
    CookieIsCleanedError = NewResponseError(5003, "Cookie已被清空,请重新登录", http.StatusOK)

    //article Error
    InsertArticleError = NewResponseError(6000, "新增文章失败", http.StatusOK)
    GetArticleError = NewResponseError(6001, "获取文章失败", http.StatusOK)
    ModifyArticleDetailError = NewResponseError(6002, "修改文章详情失败", http.StatusOK)
    DeleteArticleDetailError = NewResponseError(6003, "删除文章失败", http.StatusOK)
    SearchArticleError = NewResponseError(6004, "搜索文章失败", http.StatusOK)

    //label Error
    GetLabelError = NewResponseError(7000, "获取标签失败", http.StatusOK)
    AddLabelError = NewResponseError(7001, "添加标签失败", http.StatusOK)
    DeleteLabelError = NewResponseError(7002, "删除标签失败", http.StatusOK)
    ModifyLabelError = NewResponseError(7003, "修改标签失败", http.StatusOK)

    //activity Error
    InsertActivityError = NewResponseError(8000, "新增活动失败", http.StatusOK)
    GetActivityError = NewResponseError(8001, "获取活动失败", http.StatusOK)
    ModifyActivityDetailError = NewResponseError(8002, "修改活动详情失败", http.StatusOK)
    DeleteActivityDetailError = NewResponseError(8003, "删除活动失败", http.StatusOK)
    SearchActivityError = NewResponseError(8004, "搜索活动失败", http.StatusOK)
)
