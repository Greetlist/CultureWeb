package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    //"github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    //"github.com/Greetlist/CultureWeb/web_admin/server/config"
)

// GetTotalLabel godoc
// @Summary Query Total Label
// @Description Return Total Label
// @ID GetTotalLabel
// @Produce json
// @Success 200 {object} model.GetTotalLabelResponse
// @Router /api/user/normal/getTotalLabel [get]
func GetTotalLabel(c *gin.Context) {
    var res model.GetTotalLabelResponse
    if e := model.LabelModel.GetLabelList(&res); e != nil {
        model.GenErrorReturn(ErrorCode.GetLabelError, &res.Result)
        c.JSON(ErrorCode.GetLabelError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// AddSingleLabel godoc
// @Summary Add Label
// @Description Add New Label
// @ID AddSingleLabel
// @Produce json
// @Param request_json body model.AddSingleLabelRequest true "label Form"
// @Success 200 {object} model.AddSingleLabelResponse
// @Router /api/admin/addLabel [post]
func AddSingleLabel(c *gin.Context) {
    var req model.AddSingleLabelRequest
    var res model.AddSingleLabelResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    if e := model.LabelModel.AddSingleLabel(&req); e != nil {
        model.GenErrorReturn(ErrorCode.AddLabelError, &res.Result)
        c.JSON(ErrorCode.AddLabelError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// DeleteLabel godoc
// @Summary Delete Label
// @Description Delete Label
// @ID DeleteLabel
// @Produce json
// @Param request_json body model.DeleteLabelRequest true "delete label list"
// @Success 200 {object} model.DeleteLabelResponse
// @Router /api/admin/deleteLabel [post]
func DeleteLabel(c *gin.Context) {
    var req model.DeleteLabelRequest
    var res model.DeleteLabelResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    LOG.Logger.Infof("%v", req)
    if e := model.LabelModel.DeleteLabel(&req.Labels); e != nil {
        model.GenErrorReturn(ErrorCode.DeleteLabelError, &res.Result)
        c.JSON(ErrorCode.DeleteLabelError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// ModifyLabel godoc
// @Summary Modify Label
// @Description Modify Label
// @ID ModifyLabel
// @Produce json
// @Param request_json body model.ModifyLabelRequest true "modify label list"
// @Success 200 {object} model.ModifyLabelResponse
// @Router /api/admin/modifyLabel [post]
func ModifyLabel(c *gin.Context) {
    var req model.ModifyLabelRequest
    var res model.ModifyLabelResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    LOG.Logger.Infof("%v", req)
    if e := model.LabelModel.ModifyLabel(&req.ModifyList); e != nil {
        model.GenErrorReturn(ErrorCode.ModifyLabelError, &res.Result)
        c.JSON(ErrorCode.ModifyLabelError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
