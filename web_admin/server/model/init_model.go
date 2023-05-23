package model

import (
    "gorm.io/gorm"
)

var (
    UserModel UserModelStruct
    MediaModel MediaModelStruct
    ArticleModel ArticleModelStruct
    ActivityModel ActivityModelStruct
    StatisticModel StatisticModelStruct
    LabelModel LabelModelStruct
    SiteModel SiteModelStruct
)

func InitModel(db *gorm.DB) {
    UserModel = UserModelStruct{DB: db}
    MediaModel = MediaModelStruct{DB: db}
    ArticleModel = ArticleModelStruct{DB: db}
    ActivityModel = ActivityModelStruct{DB: db}
    StatisticModel = StatisticModelStruct{DB: db}
    LabelModel = LabelModelStruct{DB: db}
    SiteModel = SiteModelStruct{DB: db}
}
