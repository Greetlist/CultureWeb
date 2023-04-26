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

type AdminFetchArticleStruct struct {
    ArticleID uint `gorm:"column:article_id;primary_key;AUTO_INCREMENT;" json:"article_id"`
    Title string `gorm:"column:title;not null;" json:"title"`
    Rank uint `gorm:"column:rank;default:0;" json:"rank"`
    Summary string `gorm:"column:summary;not null;" json:"summary"`
    State string `gorm:"column:state;default:normal;" json:"state"`
    Author string `gorm:"column:author;not null;" json:"author"`
    IsEnable bool `gorm:"column:is_enable;default:1;" json:"is_enable"`
    IsTop bool `gorm:"column:is_top;default:0;" json:"is_top"`
    IsDelete bool `gorm:"column:is_delete;not null;" json:"is_delete"`
    VisitNumber uint `gorm:"column:visit_number;default:0;" json:"visit_number"`
    LocalSaveName string `gorm:"column:local_save_name;not null;" json:"local_save_name"`
}

type GetTotalArticleResponse struct {
    ArticleList []AdminFetchArticleStruct `json:"article_list"`
    Result RequestResult `json:"request_result"`
}
