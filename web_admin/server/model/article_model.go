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
    var labelList []*schema.Label
    for _, labelID := range(req.Labels) {
        var item schema.Label
        item.LabelID = labelID
        labelList = append(labelList, &item)
    }
    a := &schema.Article {
        ArticleDetail: schema.ArticleDetail {
            ArticleID: 0,
            Title: req.Title,
            Rank: req.Rank,
            Summary: req.Summary,
            Labels: labelList,
            IsTop: req.IsTop,
            LocalSaveName: local_save_name,
        },
        Content: req.Content,
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
    var articleList []schema.Article
    query_res := article.DB.Model(schema.Article{}).Preload("Labels").Find(&articleList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.GetArticleError
    }
    for _, article := range(articleList) {
        res.ArticleList = append(res.ArticleList, article.ArticleDetail)
    }
    return nil
}

func (article *ArticleModelStruct) BatchModifyArticle(articleList *[]ModifyArticleItem) *ErrorCode.ResponseError {
    tx := article.DB.Begin()
    for _, item := range(*articleList) {
        var article schema.Article
        article.ArticleDetail.ArticleID = item.ArticleID
        if e := tx.Model(&article).Updates(
            schema.Article{
                ArticleDetail: schema.ArticleDetail{
                    Title: item.Title,
                    Summary: item.Summary,
                    Rank: item.Rank,
                    IsTop: item.IsTop,
                    IsEnable: item.IsEnable,
                },
            }).Error; e != nil {
            LOG.Logger.Errorf("DB Error: %v", e)
            tx.Rollback()
            return ErrorCode.ModifyArticleDetailError
        }
        var labelList []*schema.Label
        for _, labelID := range(item.Labels) {
            labelList = append(labelList, &schema.Label{LabelID: labelID})
        }
        LOG.Logger.Infof("labelList: %v", labelList)
        tx.Model(&article).Association("Labels").Replace(labelList)
    }
    if e := tx.Commit().Error; e != nil {
        LOG.Logger.Errorf("DB Error: %v", e)
        tx.Rollback()
        return ErrorCode.ModifyArticleDetailError
    }
    return nil
}

func (article *ArticleModelStruct) BatchDeleteArticle(deleteList *[]uint) *ErrorCode.ResponseError {
    tx := article.DB.Begin()
    for _, articleID := range(*deleteList) {
        var article schema.Article
        article.ArticleDetail.ArticleID = articleID
        tx.Model(&article).Association("Labels").Clear()
        tx.Delete(&article)
    }
    if e := tx.Commit().Error; e != nil {
        LOG.Logger.Errorf("DB Error: %v", e)
        tx.Rollback()
        return ErrorCode.DeleteArticleDetailError
    }
    return nil
}

func (article *ArticleModelStruct) SearchArticle(keyWord string, res *SearchArticleResponse) *ErrorCode.ResponseError {
    var idList []uint
    article.DB.Raw("SELECT article_id FROM article WHERE MATCH(`content`) AGAINST(?)", keyWord).Scan(&idList)

    var articleList []schema.Article
    query_res := article.DB.Where("article_id IN ?", idList).Model(schema.Article{}).Preload("Labels").Find(&articleList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.SearchArticleError
    }
    for _, article := range(articleList) {
        res.ArticleList = append(res.ArticleList, article.ArticleDetail)
    }
    return nil
}
