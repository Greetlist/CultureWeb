package schema

type UserVerification struct {
    ID uint `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
    UserID uint `gorm:"column:user_id;not null" json:"user_id"`
    IsVerified bool `gorm:"column:user_id;default:0;" json:"user_verified"`
    IsReject bool `gorm:"column:is_reject;default:0;" json:"is_reject"`
    IsActive bool `gorm:"column:is_active;default:1;" json:"is_active"`
    Medias []Media `gorm:"column:medias;foreignKey:ID" json:"medias"`
}

func (uv UserVerification) TableName() string {
    return "user_verification"
}
