package model

type GetTotalMediaResponse struct {
    Result RequestResult `json:"request_result"`
}

type SaveMediaResponse struct {
    Result RequestResult `json:"request_result"`
    FetchUrl string `json:"fetch_url"`
}
