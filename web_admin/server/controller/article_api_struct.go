package controller

type SubmitArticleRequest struct {
    Title string `json:"title"`
    Summary string `json:"summary"`
    Rank int `json:"rank"`
    Label string `json:"label"`
    IsTop bool `json:"is_top"`
    Content string `json:"content"`
}
type SubmitArticleResponse struct {
    Result RequestResult `json:"request_result"`
}
