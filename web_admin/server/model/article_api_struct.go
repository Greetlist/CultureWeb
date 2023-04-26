package model

type SubmitArticleRequest struct {
    Title string `json:"title"`
    Rank uint `json:"rank"`
    Summary string `json:"summary"`
    Label string `json:"label"`
    IsTop bool `json:"is_top"`
    Content string `json:"content"`
    Author string `json:"author"`
}
type SubmitArticleResponse struct {
    Result RequestResult `json:"request_result"`
}
