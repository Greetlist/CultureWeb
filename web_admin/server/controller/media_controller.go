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

// GetTotalMedia godoc
// @Summary Get Total Media
// @Description Return Media List
// @ID GetTotalMedia
// @Produce json
// @Success 200 {object} GetTotalMediaResponse
// @Router /api/admin/getTotalMedia [get]
func getTotalMedia(c *gin.Context) {
    var res GetTotalMediaResponse
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
