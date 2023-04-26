package model

import (
    "bufio"
    "os"
    "gorm.io/gorm"
    "path"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    uuid "github.com/nu7hatch/gouuid"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    "github.com/Greetlist/CultureWeb/web_admin/server/util"
)

type ArticleModelStruct struct {
    DB *gorm.DB
}

func (article *ArticleModelStruct) SaveArticle(req *SubmitArticleRequest) *ErrorCode.ResponseError {
    //save on local machine, manual RAID1
    local_save_name, e := saveArticleLocal(req)
    if e != nil {
        return e
    }
    a := &schema.Article {
        ArticleID: 0,
        Title: req.Title,
        Rank: req.Rank,
        Summary: req.Summary,
        IsTop: req.IsTop,
        Content: req.Content,
        LocalSaveName: local_save_name,
    }
    insert_res := article.DB.Create(&a)
    if insert_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", insert_res.Error)
        return ErrorCode.InsertArticleError
    }
    return nil
}

func saveArticleLocal(req *SubmitArticleRequest) (string, *ErrorCode.ResponseError) {
    todayStr := util.GetTodayStr()
    baseDir := path.Join(config.GlobalConfig.ArticleSaveDir, todayStr)
    os.MkdirAll(baseDir, os.FileMode(0744))
    uid, _ := uuid.NewV4()
    fileName := uid.String() + ".html"
    finalPath := path.Join(baseDir, fileName)
    f, err := os.Create(finalPath)
    if err != nil {
        LOG.Logger.Errorf("Open File Error: %v", err)
        return "", ErrorCode.CreateFileError
    }
    w := bufio.NewWriter(f)
    _, err = w.WriteString(req.Content)
    if err != nil {
        LOG.Logger.Errorf("Write File Error: %v", err)
        return "", ErrorCode.WriteFileError
    }
    w.Flush()
    return fileName, nil
}

func (article *ArticleModelStruct) GetTotalArticle(res *GetTotalArticleResponse) *ErrorCode.ResponseError {
    query_res := article.DB.Model(&schema.Article{}).Find(&res.ArticleList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.GetUserInfoError
    }
    return nil
}
