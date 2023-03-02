package model

import (
    "gorm.io/gorm"
)

var (
    UserModel UserModelStruct
    StatisticModel StatisticModelStruct
)

func InitModel(db *gorm.DB) {
    UserModel = UserModelStruct{DB: db}
    StatisticModel = StatisticModelStruct{DB: db}
}
