package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    //"github.com/Greetlist/CultureWeb/web_admin/server/model"
    //"github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    //LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    //ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    //"github.com/Greetlist/CultureWeb/web_admin/server/config"
)

// GetTotalLabel godoc
// @Summary Query Total Label
// @Description Return Total Label
// @ID GetTotalLabel
// @Produce json
// @Success 200 {object} GetTotalLabelResponse
// @Router /user/normal/getTotalLabel [get]
func GetTotalLabel(c *gin.Context) {
    var res GetTotalLabelResponse
    c.JSON(http.StatusOK, res)
}
