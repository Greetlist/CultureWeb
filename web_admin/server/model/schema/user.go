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
    Age int `gorm:"column:age;"`
    Address string `gorm:"column:address;"`
    IsActive bool `gorm:"column:is_active;not null;"`
    State string `gorm:"state;default:normal;"`
    Score int `gorm:"column:score;default:0;"`
    Interests []Interest `gorm:"column:interests;foreignKey:UserID"`
    CreateTime time.Time `gorm:"column:create_time;autoCreateTime:milli;"`
    UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime:milli;"`
}
