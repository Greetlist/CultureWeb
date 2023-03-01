package controller

type RequestResult struct {
    ReturnCode int `json:"return_code"`
    ErrorMsg string `json:"error_msg"`
}
