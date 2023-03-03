package schema

type Interest struct {
    InterestID uint `gorm:"column:"interest_id";primary_key;AUTO_INCREMENT;" json:"interest_id"`
    UserID uint `gorm:"column:user_id;" json:"user_id"`
    ArticleID uint `gorm:"article_id;" json:"article_id"`
}

func (i Interest) TableName() string {
    return "interest"
}
