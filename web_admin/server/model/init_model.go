package model

import (
    "gorm.io/gorm"
)

var (
    UserModel UserModelStruct
    MediaModel MediaModelStruct
    StatisticModel StatisticModelStruct
)

func InitModel(db *gorm.DB) {
    UserModel = UserModelStruct{DB: db}
    MediaModel = MediaModelStruct{DB: db}
    StatisticModel = StatisticModelStruct{DB: db}
}
