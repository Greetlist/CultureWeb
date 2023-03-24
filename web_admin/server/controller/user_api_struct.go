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

type AddUserRequest struct {
    Account string `json:"account"`
    Password string `json:"password"`
    Name string `json:"name"`
    PhoneNumber string `json:"phone_number"`
    Sex string `json:"sex"`
    Age int `json:"age"`
    Address string `json:"address"`
    IdentifyID string `json:"identify_id"`
    IsActive bool `json:"is_active"`
    IsAdmin bool `json:"is_admin"`
    State string `json:"state"`
}
type AddUserResponse struct {
    Result RequestResult `json:"request_result"`
}

type UserRegisterRequest struct {
    Account string `json:"account"`
    Password string `json:"password"`
    Name string `json:"name"`
    PhoneNumber string `json:"phone_number"`
    Sex string `json:"sex"`
    Age int `json:"age"`
    Address string `json:"address"`
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
    Name string `json:"name"`
    Password string `json:"password"`
    Sex string `json:"sex"`
    Age int `json:"age"`
}
type UserModifyResponse struct {
    Result RequestResult `json:"request_result"`
}

type ChangePwdRequest struct {
    Password string `json:"password"`
}
type ChangePwdResponse struct {
    Result RequestResult `json:"request_result"`
}
