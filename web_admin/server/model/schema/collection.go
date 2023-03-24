package schema

type Collection struct {
    CollectionID uint `gorm:"column:"collection_id";primary_key;AUTO_INCREMENT;" json:"collection_id"`
    UserID uint `gorm:"column:user_id;" json:"user_id"`
    ArticleID uint `gorm:"article_id;" json:"article_id"`
}

func (i Collection) TableName() string {
    return "collection"
}
