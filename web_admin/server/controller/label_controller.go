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

    LOG.Logger.Infof("Req is : %v", req)
    if e := model.LabelModel.AddSingleLabel(&req); e != nil {
        model.GenErrorReturn(ErrorCode.AddLabelError, &res.Result)
        c.JSON(ErrorCode.AddLabelError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
