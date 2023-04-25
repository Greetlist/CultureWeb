package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/model"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/config"
)

// SubmitArticle godoc
// @Summary Submit Article
// @Description  Submit Article
// @ID SubmitArticle
// @Produce json
// @Param request_json body SubmitArticleRequest true "Ariticle Form"
// @Success 200 {object} SubmitArticleResponse
// @Router /api/admin/submitArticle [post]
func SubmitArticle(c *gin.Context) {
    var req SubmitArticleRequest
    var res SubmitArticleResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("%v", e)
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    LOG.Logger.Infof("Req is: %v", req)
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}






