package model

import (
    "gorm.io/gorm"
)

var (
    UserModel UserModelStruct
    MediaModel MediaModelStruct
    ArticleModel ArticleModelStruct
    StatisticModel StatisticModelStruct
    LabelModel LabelModelStruct
)

func InitModel(db *gorm.DB) {
    UserModel = UserModelStruct{DB: db}
    MediaModel = MediaModelStruct{DB: db}
    ArticleModel = ArticleModelStruct{DB: db}
    StatisticModel = StatisticModelStruct{DB: db}
    LabelModel = LabelModelStruct{DB: db}
}
