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

// GetTotalSite godoc
// @Summary Query Total Site
// @Description Return Total Site
// @ID GetTotalSite
// @Produce json
// @Success 200 {object} model.GetTotalSiteResponse
// @Router /api/user/normal/getTotalSite [get]
func GetTotalSite(c *gin.Context) {
    var res model.GetTotalSiteResponse
    if e := model.SiteModel.GetSiteList(&res); e != nil {
        model.GenErrorReturn(ErrorCode.GetSiteError, &res.Result)
        c.JSON(ErrorCode.GetSiteError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// AddSingleSite godoc
// @Summary Add Site
// @Description Add New Site
// @ID AddSingleSite
// @Produce json
// @Param request_json body model.AddSingleSiteRequest true "site Form"
// @Success 200 {object} model.AddSingleSiteResponse
// @Router /api/admin/addSite [post]
func AddSingleSite(c *gin.Context) {
    var req model.AddSingleSiteRequest
    var res model.AddSingleSiteResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    if e := model.SiteModel.AddSingleSite(&req); e != nil {
        model.GenErrorReturn(ErrorCode.AddSiteError, &res.Result)
        c.JSON(ErrorCode.AddSiteError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// DeleteSite godoc
// @Summary Delete Site
// @Description Delete Site
// @ID DeleteSite
// @Produce json
// @Param request_json body model.DeleteSiteRequest true "delete site list"
// @Success 200 {object} model.DeleteSiteResponse
// @Router /api/admin/deleteSite [post]
func DeleteSite(c *gin.Context) {
    var req model.DeleteSiteRequest
    var res model.DeleteSiteResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    LOG.Logger.Infof("%v", req)
    if e := model.SiteModel.DeleteSite(&req.Sites); e != nil {
        model.GenErrorReturn(ErrorCode.DeleteSiteError, &res.Result)
        c.JSON(ErrorCode.DeleteSiteError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// ModifySite godoc
// @Summary Modify Site
// @Description Modify Site
// @ID ModifySite
// @Produce json
// @Param request_json body model.ModifySiteRequest true "modify site list"
// @Success 200 {object} model.ModifySiteResponse
// @Router /api/admin/modifySite [post]
func ModifySite(c *gin.Context) {
    var req model.ModifySiteRequest
    var res model.ModifySiteResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    LOG.Logger.Infof("%v", req)
    if e := model.SiteModel.ModifySite(&req.ModifyList); e != nil {
        model.GenErrorReturn(ErrorCode.ModifySiteError, &res.Result)
        c.JSON(ErrorCode.ModifySiteError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
