package schema

import (
    "time"
)

type User struct {
    ID uint `gorm:"primary_key;AUTO_INCREMENT;"`
    Account string `gorm:"column:account;not null;"`
    Password string `gorm:"column:password;not null;"`
    Name string `gorm:"column:name;not null;"`
    Sex string `gorm:"column:sex;not null;"`
    Age int `gorm:"column:age;default:0;"`
    Address string `gorm:"column:address;default:NULL;"`
    IdentifyID string `gorm:"column:idnentify_id;default:NULL;"`
    IsVerify bool `gorm:"column:is_verify;default:0;"`
    IsActive bool `gorm:"column:is_active;default:1;"`
    State string `gorm:"state;default:normal;"`
    Score int `gorm:"column:score;default:0;"`
    Interests []Interest `gorm:"column:interests;foreignKey:UserID"`
    CreateTime time.Time `gorm:"column:create_time;autoCreateTime:milli;"`
    UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime:milli;"`
}

func (user User) TableName() string {
    return "user"
}
