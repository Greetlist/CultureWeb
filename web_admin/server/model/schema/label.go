package schema

type Label struct {
    LabelID int `gorm:"column:label_id;primary_key;not null;" json:"label_id"`
    LabelName string `gorm:"column:label_name;" json:"label_name"`
    ParentID int `gorm:"column:parent_id;" json:"parent_id"`
}

func (uv Label) TableName() string {
    return "label"
}
