package schema

type Label struct {
    LabelID int `gorm:"column:label_id;primary_key;not null;autoIncrement" json:"label_id"`
    LabelName string `gorm:"column:label_name;unique;not null" json:"label_name"`
    Articles []*Article `gorm:"many2many:article_labels;" json:"articles"`
}

func (l Label) TableName() string {
    return "label"
}
