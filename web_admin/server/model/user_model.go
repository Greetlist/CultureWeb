package model

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
)

type UserRegisterRequest struct {
    Account string `json:"account"`
    Password string `json:"password"`
    Name string `json:"name"`
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
}

type UserLogoutRequest struct {
    Account string `json:"account"`
}
type UserLogoutResponse struct {
    Result RequestResult `json:"request_result"`
}

type GetUserInfoResponse struct {
    Result RequestResult `json:"request_result"`
}

type GetTotalUserInfoResponse struct {
    Result RequestResult `json:"request_result"`
    Users []schema.User `json:"user_infos"`
}
