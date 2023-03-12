package schema

type UserVerification struct {
    UserVerificationID uint `gorm:"column:user_verification_id;primary_key;AUTO_INCREMENT;" json:"user_verification_id"`
    UserID uint `gorm:"column:user_id;not null" json:"user_id"`
    IsVerified bool `gorm:"column:user_id;default:0;" json:"user_verified"`
    IsReject bool `gorm:"column:is_reject;default:0;" json:"is_reject"`
    IsActive bool `gorm:"column:is_active;default:1;" json:"is_active"`
    FirstMedia uint `gorm:"column:first_media" json:"first_media"`
    SecondMedia uint `gorm:"column:second_media" json:"second_media"`
}

func (uv UserVerification) TableName() string {
    return "user_verification"
}
