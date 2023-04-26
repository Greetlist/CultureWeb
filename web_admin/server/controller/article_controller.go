package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/config"
)

// SubmitArticle godoc
// @Summary Submit Article
// @Description  Submit Article
// @ID SubmitArticle
// @Produce json
// @Param request_json body model.SubmitArticleRequest true "Ariticle Form"
// @Success 200 {object} model.SubmitArticleResponse
// @Router /api/admin/submitArticle [post]
func SubmitArticle(c *gin.Context) {
    var req model.SubmitArticleRequest
    var res model.SubmitArticleResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    LOG.Logger.Infof("Req is: %v", req)
    if e := model.ArticleModel.SaveArticle(&req); e != nil {
        model.GenErrorReturn(ErrorCode.InsertArticleError, &res.Result)
        c.JSON(ErrorCode.InsertArticleError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
