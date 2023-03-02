package schema

type UserVerification struct {
    UserVerificationID uint `gorm:"column:user_verification_id;primary_key;AUTO_INCREMENT;" json:"user_verification_id"`
    UserID uint `gorm:"column:user_id;not null" json:"user_id"`
    IsVerified bool `gorm:"column:user_id;default:0;" json:"user_verified"`
    IsReject bool `gorm:"column:is_reject;default:0;" json:"is_reject"`
    IsActive bool `gorm:"column:is_active;default:1;" json:"is_active"`
    Medias []Media `gorm:"foreignKey:MediaID" json:"medias"`
}

func (uv UserVerification) TableName() string {
    return "user_verification"
}
