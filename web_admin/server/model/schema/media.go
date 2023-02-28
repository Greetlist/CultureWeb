package schema

type Media struct {
    ID uint `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
    Path string `gorm:"column:path;not null" json:"path"`
    Size uint `gorm:"column:size;default:0" json:"size"`
    CategoryName string `gorm:"category_name;not null" json:"category_name"`
    Category int `gorm:"category;not null" json:"category"`
}

func (uv Media) TableName() string {
    return "media"
}
