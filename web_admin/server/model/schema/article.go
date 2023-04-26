package schema

import (
    "time"
)

type Comment struct {
    CommentID uint `gorm:"column:comment_id;primary_key;AUTO_INCREMENT;" json:"comment_id"`
    ArticleID uint `gorm:"column:article_id;not null;" json:"ariticle_id"`
    UserID uint `gorm:"column:user_id;not null" json:"user_id"`
    Content string `gorm:"column:content;not null;" json:"content"`
    CreateTime time.Time `gorm:"column:create_time;autoCreateTime:milli;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime:milli;" json:"update_time"`
}

type Article struct {
    ArticleID uint `gorm:"column:article_id;primary_key;AUTO_INCREMENT;" json:"article_id"`
    Title string `gorm:"column:title;not null;" json:"title"`
    Rank uint `gorm:"column:rank;default:0;" json:"rank"`
    Summary string `gorm:"column:summary;not null;" json:"summary"`
    Content string `gorm:"column:content;not null;" json:"content"`
    Labels []Label `gorm:"many2many:labels;" json:"labels"`
    Comments []Comment `gorm:"foreignKey:CommentID" json:"comments"`
    State string `gorm:"column:state;default:normal;" json:"state"`
    Author string `gorm:"column:author;not null;" json:"author"`
    IsEnable bool `gorm:"column:is_enable;default:1;" json:"is_enable"`
    IsTop bool `gorm:"column:is_top;default:0;" json:"is_top"`
    IsDelete bool `gorm:"column:is_delete;not null;" json:"is_delete"`
    VisitNumber uint `gorm:"column:visit_number;default:0;" json:"visit_number"`
    LocalSaveName string `gorm:"column:local_save_name;not null;" json:"local_save_name"`
    CreateTime time.Time `gorm:"column:create_time;autoCreateTime:milli;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime:milli;" json:"update_time"`
}

func (uv Article) TableName() string {
    return "article"
}

func (c Comment) TableName() string {
    return "comment"
}
