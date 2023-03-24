package schema

type Label struct {
    LabelID uint `gorm:"column:label_id;primary_key;AUTO_INCREMENT;" json:"label_id"`
    LabelName string `gorm:"column:label_name;" json:"label_name"`
    ArticleID uint `gorm:"column:article_id;primary_key;AUTO_INCREMENT;" json:"article_id"`
}

func (uv Label) TableName() string {
    return "label"
}
