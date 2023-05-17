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
    if e := model.ArticleModel.SaveArticle(&req); e != nil {
        model.GenErrorReturn(ErrorCode.InsertArticleError, &res.Result)
        c.JSON(ErrorCode.InsertArticleError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// GetTotalArticle godoc
// @Summary Get Total Article
// @Description Get Total Article
// @ID GetTotalArticle
// @Produce json
// @Success 200 {object} model.GetTotalArticleResponse
// @Router /api/admin/getTotalArticle [get]
func GetTotalArticle(c *gin.Context) {
    var res model.GetTotalArticleResponse
    if e := model.ArticleModel.GetTotalArticle(&res); e != nil {
        model.GenErrorReturn(ErrorCode.GetArticleError, &res.Result)
        c.JSON(ErrorCode.GetArticleError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// SearchArticle godoc
// @Summary Seearch Article
// @Description Search Article
// @ID SearchArticle
// @Produce json
// @Param request_json body model.SearchArticleRequest true "Search Ariticle Form"
// @Success 200 {object} model.SearchArticleResponse
// @Router /api/user/normal/searchArticle [post]
func SearchArticle(c *gin.Context) {
    var req model.SearchArticleRequest
    var res model.SearchArticleResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    if e := model.ArticleModel.SearchArticle(req.KeyWord, &res); e != nil {
        model.GenErrorReturn(ErrorCode.SearchArticleError, &res.Result)
        c.JSON(ErrorCode.GetArticleError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// BatchModifyArticle godoc
// @Summary Modify Article
// @Description Modify Article
// @ID BatchModifyArticle
// @Produce json
// @Param request_json body model.BatchModifyArticleRequest true "Ariticle Form"
// @Success 200 {object} model.BatchModifyArticleResponse
// @Router /api/admin/batchModifyArticle [post]
func BatchModifyArticle(c *gin.Context) {
    var req model.BatchModifyArticleRequest
    var res model.BatchModifyArticleResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    if e := model.ArticleModel.BatchModifyArticle(&req.ModifyList); e != nil {
        LOG.Logger.Errorf("Req param is: %v", req)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// BatchDeleteArticle godoc
// @Summary Delete Article
// @Description Delete Article
// @ID BatchDeleteArticle
// @Produce json
// @Param request_json body model.BatchDeleteArticleRequest true "Ariticle Form"
// @Success 200 {object} model.BatchDeleteArticleResponse
// @Router /api/admin/batchDeleteArticle [post]
func BatchDeleteArticle(c *gin.Context) {
    var req model.BatchDeleteArticleRequest
    var res model.BatchDeleteArticleResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    LOG.Logger.Infof("Req is: %v", req)
    if e := model.ArticleModel.BatchDeleteArticle(&req.DeleteList); e != nil {
        LOG.Logger.Errorf("Req param is: %v", req)
        c.JSON(ErrorCode.DeleteArticleDetailError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// GetArticleContent godoc
// @Summary Get Article Content
// @Description Get Article Content
// @ID GetArticleContent
// @Produce json
// @Param request_json body model.GetArticleContentRequest true "Ariticle Info"
// @Success 200 {object} model.GetArticleContentResponse
// @Router /api/user/normal/getArticleContent [post]
func GetArticleContent(c *gin.Context) {
    var req model.GetArticleContentRequest
    var res model.GetArticleContentResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    if !req.IsAdmin {
        incrErr := model.ArticleModel.IncrArticleVisitNumber(req.ArticleID)
        if incrErr != nil {
            LOG.Logger.Errorf("Redis Error, Req param is: %v", req)
        }
    }

    content, e := model.ArticleModel.GetLocalArticleContent(&req)
    if e != nil {
        LOG.Logger.Errorf("Req param is: %v", req)
        model.GenErrorReturn(ErrorCode.ReadFileError, &res.Result)
        c.JSON(ErrorCode.ReadFileError.HttpStatusCode, res)
        return
    }
    res.ArticleContent = content
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
