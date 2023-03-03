package controller

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
)

type GetUserInfoRequest struct {
    UserID uint `form:"user_id" binding:"required"`
}
type GetUserInfoResponse struct {
    Result RequestResult `json:"request_result"`
    UserInfo schema.User `json:"user_info"`
}

type UserRegisterRequest struct {
    Account string `json:"account"`
    Password string `json:"password"`
    Name string `json:"name"`
    Sex string `json:"sex"`
    Age int `json:"age"`
}
type UserRegisterResponse struct {
    Result RequestResult `json:"request_result"`
}

type UserLoginRequest struct {
    Account string `json:"account"`
    Password string `json:"password"`
}
type UserLoginResponse struct {
    Result RequestResult `json:"request_result"`
    UserID uint `json:"user_id"`
}

type UserLogoutRequest struct {
    Account string `json:"account"`
}
type UserLogoutResponse struct {
    Result RequestResult `json:"request_result"`
}

type GetTotalUserInfoResponse struct {
    Result RequestResult `json:"request_result"`
    UserInfos []schema.User `json:"user_infos"`
}

type UserModifyRequest struct {
    UserID uint `json:"user_id"`
    Name string `json:"name"`
    Sex string `json:"sex"`
    Age int `json:"age"`
}
type UserModifyResponse struct {
    Result RequestResult `json:"request_result"`
}
