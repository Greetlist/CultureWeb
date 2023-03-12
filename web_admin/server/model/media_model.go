package model

import (
    "gorm.io/gorm"
    "path"
    "os"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    "github.com/Greetlist/CultureWeb/web_admin/server/util"
    uuid "github.com/nu7hatch/gouuid"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

type MediaModelStruct struct {
    DB *gorm.DB
}

func (media *MediaModelStruct) SaveMedia(size int, category, extension string) (string, string, *ErrorCode.ResponseError) {
    todayStr := util.GetTodayStr()
    baseDir := path.Join(config.GlobalConfig.MediaSaveDir, todayStr)
    os.Mkdir(baseDir, os.FileMode(0744))
    uid, _ := uuid.NewV4()
    fileName := uid.String() + "." + extension
    var m schema.Media
    m.Path = path.Join(todayStr, fileName)
    m.Size = size
    m.CategoryName = category
    if category == "image" {
        m.Category = 1
    } else if category == "video" {
        m.Category = 2
    }
    insert_res := media.DB.Create(&m)
    if insert_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", insert_res.Error)
        return "", "", ErrorCode.RegisterUserError
    }
    return path.Join(baseDir, fileName), genFetchUrl(m.Path), nil
}

func genFetchUrl(filePath string) string {
    url := config.GlobalConfig.BaseUrl + "/" + config.GlobalConfig.MediaBaseUrl + filePath
    return url
}
