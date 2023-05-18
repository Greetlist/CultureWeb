package schema

import (
    "time"
)

type ActivityDetail struct {
    ActivityID uint `gorm:"column:activity_id;primary_key;AUTO_INCREMENT;" json:"activity_id"`
    Title string `gorm:"column:title;unique;not null;" json:"title"`
    Summary string `gorm:"column:summary;not null;" json:"summary"`
    Labels []*Label `gorm:"many2many:activity_labels;" json:"labels"`
    Author string `gorm:"column:author;not null;" json:"author"`
    IsEnable bool `gorm:"column:is_enable;default:1;" json:"is_enable"`
    VisitNumber uint `gorm:"column:visit_number;default:0;" json:"visit_number"`
    StartTime time.Time `gorm:"column:start_time;not null;" json:"start_time"`
    EndTime time.Time `gorm:"column:end_time;not null;" json:"end_time"`
    LocalSaveName string `gorm:"column:local_save_name;not null;" json:"local_save_name"`
    CreateTime time.Time `gorm:"column:create_time;autoCreateTime:milli;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime:milli;" json:"update_time"`
}

type Activity struct {
    ActivityDetail ActivityDetail `gorm:"embedded;"`
    Content string `gorm:"column:content;not null;type:longtext;index:,class:FULLTEXT,option:WITH PARSER ngram" json:"content"`
}

func (uv Activity) TableName() string {
    return "activity"
}
