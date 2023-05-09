package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/error"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
)

// GetBasicInfo godoc
// @Summary Query Total Label
// @Description Return Total Label
// @ID GetBasicInfo
// @Produce json
// @Success 200 {object} model.GetWebBasicInfoResponse
// @Router /api/user/normal/getWebBasicInfo [get]
func GetBasicInfo(c *gin.Context) {
    var res model.GetWebBasicInfoResponse
    res.WebBasicInfo = config.GlobalConfig.WebInfo
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
