package schema

import (
    "time"
)

type User struct {
    UserID uint `gorm:"column:user_id;primary_key;AUTO_INCREMENT;" json:"user_id"`
    Account string `gorm:"column:account;unique;not null;" json:"account"`
    Password string `gorm:"column:password;not null;" json:"password"`
    Name string `gorm:"column:name;not null;" json:"name"`
    Sex string `gorm:"column:sex;not null;" json:"sex"`
    Age int `gorm:"column:age;default:0;" json:"age"`
    Address string `gorm:"column:address;default:NULL;" json:"address"`
    IdentifyID string `gorm:"column:idnentify_id;default:NULL;" json:"identify_id"`
    IsVerify bool `gorm:"column:is_verify;default:0;" json:"is_verify"`
    IsActive bool `gorm:"column:is_active;default:1;" json:"is_active"`
    IsAdmin bool `gorm:"column:is_active;default:0;" json:"is_admin"`
    State string `gorm:"state;default:normal;" json:"state"`
    Score int `gorm:"column:score;default:0;" json:"score"`
    Interests []Interest `gorm:"foreignKey:InterestID" json:"interests"`
    Verifications []UserVerification `gorm:"foreignKey:UserVerificationID" json:"user_verifications"`
    CreateTime time.Time `gorm:"column:create_time;autoCreateTime:milli;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime:milli;" json:"update_time"`
}

func (user User) TableName() string {
    return "user"
}
