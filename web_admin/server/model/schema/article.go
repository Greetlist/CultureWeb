package schema

import (
    "time"
)

type Article struct {
    ArticleID uint `gorm:"column:article_id;primary_key;AUTO_INCREMENT;" json:"article_id"`
    Rank uint `gorm:"column:rank;default:0;" json:"rank"`
    Title string `gorm:"column:title;not null;" json:"title"`
    Content string `gorm:"column:content;not null;" json:"content"`
    Lables []Label `gorm:"foreignKey:LabelID" json:"labels"`
    CreateTime time.Time `gorm:"column:create_time;autoCreateTime:milli;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime:milli;" json:"update_time"`
}

func (uv Article) TableName() string {
    return "article"
}
