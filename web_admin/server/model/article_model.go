package model

import (
    "bufio"
    "os"
    "gorm.io/gorm"
    "path"
    "strings"
    "strconv"
    "fmt"
    "reflect"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    uuid "github.com/nu7hatch/gouuid"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    "github.com/Greetlist/CultureWeb/web_admin/server/util"
    redisPool "github.com/Greetlist/CultureWeb/web_admin/server/redis"
)

type ArticleModelStruct struct {
    DB *gorm.DB
}

var article_regex string = "article_visit_num#*"
var article_key_template string = "article_visit_num#%d"

func (article *ArticleModelStruct) SaveArticle(req *SubmitArticleRequest) *ErrorCode.ResponseError {
    //save on local machine, manual RAID1
    local_save_name, e := saveArticleLocal(req.Content, true, "", "")
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

func saveArticleLocal(content string, isNew bool, createTime, localSaveName string) (string, *ErrorCode.ResponseError) {
    var fileName, finalPath string
    if isNew {
        todayStr := util.GetTodayStr()
        baseDir := path.Join(config.GlobalConfig.ArticleSaveDir, todayStr)
        os.MkdirAll(baseDir, os.FileMode(0744))
        uid, _ := uuid.NewV4()
        fileName = uid.String() + ".html"
        finalPath = path.Join(baseDir, fileName)
    } else {
        strList := strings.Split(createTime, "T")
        if len(strList) < 1 {
            LOG.Logger.Errorf("CreateTime Split Error.")
            return "", ErrorCode.ParseParamError
        }
        baseDir := path.Join(config.GlobalConfig.ArticleSaveDir, strList[0])
        fileName = localSaveName
        finalPath = path.Join(baseDir, localSaveName)
    }
    f, err := os.Create(finalPath)
    if err != nil {
        LOG.Logger.Errorf("Open File Error: %v", err)
        return "", ErrorCode.CreateFileError
    }
    w := bufio.NewWriter(f)
    _, err = w.WriteString(content)
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

        if item.IsModifyContent {
            saveArticleLocal(item.Content, false, item.CreateTime, item.LocalSaveName)
            if e := tx.Model(&article).Updates(
                schema.Article{
                    ArticleDetail: schema.ArticleDetail{
                        Title: item.Title,
                        Summary: item.Summary,
                        Rank: item.Rank,
                        IsTop: item.IsTop,
                        IsEnable: item.IsEnable,
                    },
                    Content: item.Content,
                }).Error; e != nil {
                LOG.Logger.Errorf("DB Error: %v", e)
                tx.Rollback()
                return ErrorCode.ModifyArticleDetailError
            }
        } else {
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

func (article *ArticleModelStruct) GetLocalArticleContent(req *GetArticleContentRequest) (string, *ErrorCode.ResponseError) {
    strList := strings.Split(req.CreateTime, "T")
    if len(strList) < 1 {
        LOG.Logger.Errorf("CreateTime Split Error.")
        return "", ErrorCode.ParseParamError
    }
    baseDir := path.Join(config.GlobalConfig.ArticleSaveDir, strList[0])
    finalPath := path.Join(baseDir, req.LocalSaveName)
    b, err := os.ReadFile(finalPath)
    if err != nil {
        LOG.Logger.Errorf("Read File Error: %v", err)
        return "", ErrorCode.ReadFileError
    }
    content := string(b)
    return content, nil
}

func (article *ArticleModelStruct) IncrArticleVisitNumber(articleID uint) *ErrorCode.ResponseError {
    redisClient, _ := <- redisPool.RedisPool
    defer redisPool.ReturnRedisClient(redisClient)

    article_visit_key := fmt.Sprintf(article_key_template, articleID)
    if _, e := redisClient.Do("GET", article_visit_key); e != nil {
        LOG.Logger.Infof("Failed to get key from redis, fetch from db")
        articleDetail, queryError := article.GetArticleDetail(articleID);
        if queryError != nil {
            LOG.Logger.Infof("Failed to get key from DB, skip this article")
            return nil
        }
        if _, setError := redisClient.Do("SET", article_visit_key, articleDetail.VisitNumber); setError != nil {
            LOG.Logger.Infof("Failed to set key, skip this article")
            return nil
        }
    }
    if _, e := redisClient.Do("INCR", article_visit_key); e != nil {
        LOG.Logger.Errorf("INCR error is: %v", e)
        return ErrorCode.RedisCommandError
    }
    return nil
}

func (article *ArticleModelStruct) GetArticleDetail(articleID uint) (*schema.ArticleDetail, *ErrorCode.ResponseError) {
    var res schema.Article
    if query_res := article.DB.Find(&res, articleID); query_res.Error != nil {
        return nil, ErrorCode.SearchArticleError
    }
    return &res.ArticleDetail, nil
}

func (article *ArticleModelStruct) SaveArticleVisitNumToDB() {
    LOG.Logger.Infof("Start save visit number to DB")
    redisClient, _ := <- redisPool.RedisPool
    defer redisPool.ReturnRedisClient(redisClient)

    r, e := redisClient.Do("KEYS", article_regex)
    if e != nil {
        LOG.Logger.Warningf("Fetch from redis error: %v", e)
        LOG.Logger.Infof("Failed to get key from redis, quit cron function at this round")
        return
    }

    keyList := reflect.ValueOf(r)
    for i := 0; i < keyList.Len(); i++ {
        articleKey := fmt.Sprintf("%s", keyList.Index(i))
        articleID, _ := strconv.ParseUint(strings.Split(articleKey, "#")[1], 10, 64)

        redisRes, _ := redisClient.Do("GET", articleKey)
        visitNum, _ := strconv.ParseUint(fmt.Sprintf("%s", redisRes), 10, 64)

        var curArticle schema.Article
        curArticle.ArticleDetail.ArticleID = uint(articleID)
        if updateRes := article.DB.Model(&curArticle).Update("visit_number", uint(visitNum)); updateRes.Error != nil {
            LOG.Logger.Errorf("Update Article VisitNumber error: %v", updateRes.Error)
        }
    }
}
