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

type ActivityModelStruct struct {
    DB *gorm.DB
}

var activity_regex string = "activity_visit_num#*"
var activity_key_template string = "activity_visit_num#%d"

func (activity *ActivityModelStruct) SaveActivity(req *SubmitActivityRequest) *ErrorCode.ResponseError {
    //save on local machine, manual RAID1
    local_save_name, e := saveActivityLocal(req.Content, true, "", "")
    if e != nil {
        return e
    }
    var labelList []*schema.Label
    for _, labelID := range(req.Labels) {
        var item schema.Label
        item.LabelID = labelID
        labelList = append(labelList, &item)
    }
    a := &schema.Activity {
        ActivityDetail: schema.ActivityDetail {
            ActivityID: 0,
            Title: strings.Trim(req.Title, " "),
            Summary: req.Summary,
            Labels: labelList,
            LocalSaveName: local_save_name,
        },
        Content: req.Content,
    }
    insert_res := activity.DB.Create(&a)
    if insert_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", insert_res.Error)
        return ErrorCode.InsertActivityError
    }
    return nil
}

func saveActivityLocal(content string, isNew bool, createTime, localSaveName string) (string, *ErrorCode.ResponseError) {
    var fileName, finalPath string
    if isNew {
        todayStr := util.GetTodayStr()
        baseDir := path.Join(config.GlobalConfig.ActivitySaveDir, todayStr)
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
        baseDir := path.Join(config.GlobalConfig.ActivitySaveDir, strList[0])
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

func (activity *ActivityModelStruct) GetTotalActivity(res *GetTotalActivityResponse) *ErrorCode.ResponseError {
    var activityList []schema.Activity
    query_res := activity.DB.Model(schema.Activity{}).Preload("Labels").Find(&activityList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.GetActivityError
    }
    for _, activity := range(activityList) {
        res.ActivityList = append(res.ActivityList, activity.ActivityDetail)
    }
    return nil
}

func (activity *ActivityModelStruct) BatchModifyActivity(activityList *[]ModifyActivityItem) *ErrorCode.ResponseError {
    tx := activity.DB.Begin()
    for _, item := range(*activityList) {
        var activity schema.Activity
        activity.ActivityDetail.ActivityID = item.ActivityID

        if item.IsModifyContent {
            saveActivityLocal(item.Content, false, item.CreateTime, item.LocalSaveName)
            if e := tx.Model(&activity).Updates(
                schema.Activity{
                    ActivityDetail: schema.ActivityDetail{
                        Title: item.Title,
                        Summary: item.Summary,
                        IsEnable: item.IsEnable,
                    },
                    Content: item.Content,
                }).Error; e != nil {
                LOG.Logger.Errorf("DB Error: %v", e)
                tx.Rollback()
                return ErrorCode.ModifyActivityDetailError
            }
        } else {
            if e := tx.Model(&activity).Updates(
                schema.Activity{
                    ActivityDetail: schema.ActivityDetail{
                        Title: item.Title,
                        Summary: item.Summary,
                        IsEnable: item.IsEnable,
                    },
                }).Error; e != nil {
                LOG.Logger.Errorf("DB Error: %v", e)
                tx.Rollback()
                return ErrorCode.ModifyActivityDetailError
            }
        }

        var labelList []*schema.Label
        for _, labelID := range(item.Labels) {
            labelList = append(labelList, &schema.Label{LabelID: labelID})
        }
        LOG.Logger.Infof("labelList: %v", labelList)
        tx.Model(&activity).Association("Labels").Replace(labelList)
    }
    if e := tx.Commit().Error; e != nil {
        LOG.Logger.Errorf("DB Error: %v", e)
        tx.Rollback()
        return ErrorCode.ModifyActivityDetailError
    }
    return nil
}

func (activity *ActivityModelStruct) BatchDeleteActivity(deleteList *[]uint) *ErrorCode.ResponseError {
    tx := activity.DB.Begin()
    for _, activityID := range(*deleteList) {
        var activity schema.Activity
        activity.ActivityDetail.ActivityID = activityID
        tx.Model(&activity).Association("Labels").Clear()
        tx.Delete(&activity)
    }
    if e := tx.Commit().Error; e != nil {
        LOG.Logger.Errorf("DB Error: %v", e)
        tx.Rollback()
        return ErrorCode.DeleteActivityDetailError
    }
    return nil
}

func (activity *ActivityModelStruct) SearchActivity(keyWord string, res *SearchActivityResponse) *ErrorCode.ResponseError {
    var idList []uint
    activity.DB.Raw("SELECT activity_id FROM activity WHERE MATCH(`content`) AGAINST(?)", keyWord).Scan(&idList)

    var activityList []schema.Activity
    query_res := activity.DB.Where("activity_id IN ?", idList).Model(schema.Activity{}).Preload("Labels").Find(&activityList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.SearchActivityError
    }
    for _, activity := range(activityList) {
        res.ActivityList = append(res.ActivityList, activity.ActivityDetail)
    }
    return nil
}

func (activity *ActivityModelStruct) GetLocalActivityContent(req *GetActivityContentRequest) (string, *ErrorCode.ResponseError) {
    strList := strings.Split(req.CreateTime, "T")
    if len(strList) < 1 {
        LOG.Logger.Errorf("CreateTime Split Error.")
        return "", ErrorCode.ParseParamError
    }
    baseDir := path.Join(config.GlobalConfig.ActivitySaveDir, strList[0])
    finalPath := path.Join(baseDir, req.LocalSaveName)
    b, err := os.ReadFile(finalPath)
    if err != nil {
        LOG.Logger.Errorf("Read File Error: %v", err)
        return "", ErrorCode.ReadFileError
    }
    content := string(b)
    return content, nil
}

func (activity *ActivityModelStruct) IncrActivityVisitNumber(activityID uint) *ErrorCode.ResponseError {
    redisClient, _ := <- redisPool.RedisPool
    defer redisPool.ReturnRedisClient(redisClient)

    activity_visit_key := fmt.Sprintf(activity_key_template, activityID)
    if _, e := redisClient.Do("GET", activity_visit_key); e != nil {
        LOG.Logger.Infof("Failed to get key from redis, fetch from db")
        activityDetail, queryError := activity.GetActivityDetail(activityID);
        if queryError != nil {
            LOG.Logger.Infof("Failed to get key from DB, skip this activity")
            return nil
        }
        if _, setError := redisClient.Do("SET", activity_visit_key, activityDetail.VisitNumber); setError != nil {
            LOG.Logger.Infof("Failed to set key, skip this activity")
            return nil
        }
    }
    if _, e := redisClient.Do("INCR", activity_visit_key); e != nil {
        LOG.Logger.Errorf("INCR error is: %v", e)
        return ErrorCode.RedisCommandError
    }
    return nil
}

func (activity *ActivityModelStruct) GetActivityDetail(activityID uint) (*schema.ActivityDetail, *ErrorCode.ResponseError) {
    var res schema.Activity
    if query_res := activity.DB.Find(&res, activityID); query_res.Error != nil {
        return nil, ErrorCode.SearchActivityError
    }
    return &res.ActivityDetail, nil
}

func (activity *ActivityModelStruct) SaveActivityVisitNumToDB() {
    LOG.Logger.Infof("Start save visit number to DB")
    redisClient, _ := <- redisPool.RedisPool
    defer redisPool.ReturnRedisClient(redisClient)

    r, e := redisClient.Do("KEYS", activity_regex)
    if e != nil {
        LOG.Logger.Warningf("Fetch from redis error: %v", e)
        LOG.Logger.Infof("Failed to get key from redis, quit cron function at this round")
        return
    }

    keyList := reflect.ValueOf(r)
    for i := 0; i < keyList.Len(); i++ {
        activityKey := fmt.Sprintf("%s", keyList.Index(i))
        activityID, _ := strconv.ParseUint(strings.Split(activityKey, "#")[1], 10, 64)

        redisRes, _ := redisClient.Do("GET", activityKey)
        visitNum, _ := strconv.ParseUint(fmt.Sprintf("%s", redisRes), 10, 64)

        var curActivity schema.Activity
        curActivity.ActivityDetail.ActivityID = uint(activityID)
        if updateRes := activity.DB.Model(&curActivity).Update("visit_number", uint(visitNum)); updateRes.Error != nil {
            LOG.Logger.Errorf("Update Activity VisitNumber error: %v", updateRes.Error)
        }
    }
}
