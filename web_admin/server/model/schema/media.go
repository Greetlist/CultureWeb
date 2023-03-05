package schema

type Media struct {
    MediaID uint `gorm:"column:media_id;primary_key;AUTO_INCREMENT;" json:"media_id"`
    Path string `gorm:"column:path;not null" json:"path"`
    Size uint `gorm:"column:size;default:0" json:"size"`
    CategoryName string `gorm:"category_name;not null" json:"category_name"`
    Category int `gorm:"category;not null" json:"category"`
    UserVerificationID uint `gorm:"column:user_verification_id" json:"user_verification_id"`
}

func (uv Media) TableName() string {
    return "media"
}
