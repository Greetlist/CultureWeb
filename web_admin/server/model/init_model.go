package model

import (
    "gorm.io/gorm"
)

var (
    UserModel UserModelStruct
)

func InitModel(db *gorm.DB) {
    UserModel = UserModelStruct{DB: db}
}
