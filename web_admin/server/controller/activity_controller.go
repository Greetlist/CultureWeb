package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/config"
)

// SubmitActivity godoc
// @Summary Submit Activity
// @Description  Submit Activity
// @ID SubmitActivity
// @Produce json
// @Param request_json body model.SubmitActivityRequest true "Ariticle Form"
// @Success 200 {object} model.SubmitActivityResponse
// @Router /api/admin/submitActivity [post]
func SubmitActivity(c *gin.Context) {
    var req model.SubmitActivityRequest
    var res model.SubmitActivityResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    if e := model.ActivityModel.SaveActivity(&req); e != nil {
        model.GenErrorReturn(ErrorCode.InsertActivityError, &res.Result)
        c.JSON(ErrorCode.InsertActivityError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// GetTotalActivity godoc
// @Summary Get Total Activity
// @Description Get Total Activity
// @ID GetTotalActivity
// @Produce json
// @Success 200 {object} model.GetTotalActivityResponse
// @Router /api/admin/getTotalActivity [get]
func GetTotalActivity(c *gin.Context) {
    var res model.GetTotalActivityResponse
    if e := model.ActivityModel.GetTotalActivity(&res); e != nil {
        model.GenErrorReturn(ErrorCode.GetActivityError, &res.Result)
        c.JSON(ErrorCode.GetActivityError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// SearchActivity godoc
// @Summary Seearch Activity
// @Description Search Activity
// @ID SearchActivity
// @Produce json
// @Param request_json body model.SearchActivityRequest true "Search Ariticle Form"
// @Success 200 {object} model.SearchActivityResponse
// @Router /api/user/normal/searchActivity [post]
func SearchActivity(c *gin.Context) {
    var req model.SearchActivityRequest
    var res model.SearchActivityResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    if e := model.ActivityModel.SearchActivity(req.KeyWord, &res); e != nil {
        model.GenErrorReturn(ErrorCode.SearchActivityError, &res.Result)
        c.JSON(ErrorCode.GetActivityError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// BatchModifyActivity godoc
// @Summary Modify Activity
// @Description Modify Activity
// @ID BatchModifyActivity
// @Produce json
// @Param request_json body model.BatchModifyActivityRequest true "Ariticle Form"
// @Success 200 {object} model.BatchModifyActivityResponse
// @Router /api/admin/batchModifyActivity [post]
func BatchModifyActivity(c *gin.Context) {
    var req model.BatchModifyActivityRequest
    var res model.BatchModifyActivityResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    if e := model.ActivityModel.BatchModifyActivity(&req.ModifyList); e != nil {
        LOG.Logger.Errorf("Req param is: %v", req)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// BatchDeleteActivity godoc
// @Summary Delete Activity
// @Description Delete Activity
// @ID BatchDeleteActivity
// @Produce json
// @Param request_json body model.BatchDeleteActivityRequest true "Ariticle Form"
// @Success 200 {object} model.BatchDeleteActivityResponse
// @Router /api/admin/batchDeleteActivity [post]
func BatchDeleteActivity(c *gin.Context) {
    var req model.BatchDeleteActivityRequest
    var res model.BatchDeleteActivityResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    LOG.Logger.Infof("Req is: %v", req)
    if e := model.ActivityModel.BatchDeleteActivity(&req.DeleteList); e != nil {
        LOG.Logger.Errorf("Req param is: %v", req)
        c.JSON(ErrorCode.DeleteActivityDetailError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// GetActivityContent godoc
// @Summary Get Activity Content
// @Description Get Activity Content
// @ID GetActivityContent
// @Produce json
// @Param request_json body model.GetActivityContentRequest true "Ariticle Info"
// @Success 200 {object} model.GetActivityContentResponse
// @Router /api/user/normal/getActivityContent [post]
func GetActivityContent(c *gin.Context) {
    var req model.GetActivityContentRequest
    var res model.GetActivityContentResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    if !req.IsAdmin {
        incrErr := model.ActivityModel.IncrActivityVisitNumber(req.ActivityID)
        if incrErr != nil {
            LOG.Logger.Errorf("Redis Error, Req param is: %v", req)
        }
    }

    content, e := model.ActivityModel.GetLocalActivityContent(&req)
    if e != nil {
        LOG.Logger.Errorf("Req param is: %v", req)
        model.GenErrorReturn(ErrorCode.ReadFileError, &res.Result)
        c.JSON(ErrorCode.ReadFileError.HttpStatusCode, res)
        return
    }
    res.ActivityContent = content
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
