package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    //"github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    //LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    //"github.com/Greetlist/CultureWeb/web_admin/server/config"
)

// GetTotalLabel godoc
// @Summary Query Total Label
// @Description Return Total Label
// @ID GetTotalLabel
// @Produce json
// @Success 200 {object} model.GetTotalLabelResponse
// @Router /user/normal/getTotalLabel [get]
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
