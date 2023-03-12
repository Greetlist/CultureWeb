package controller

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/config"
)

// SaveMedia godoc
// @Summary Save Media
// @Description Save Media to local storage
// @ID SaveMedia
// @Produce json
// @Success 200 {object} SaveMediaResponse
// @Router /api/admin/saveMedia [post]
func SaveMedia(c *gin.Context) {
    var res SaveMediaResponse
    file, e := c.FormFile("file")
    if e != nil {
        LOG.Logger.Errorf("Empty File Error: %v", ErrorCode.EmptyFileParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    size := c.PostForm("size")
    sizeInt, e := strconv.Atoi(size)
    if e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    category := c.PostForm("category")
    extension := c.PostForm("extension")
    savePath, fetchUrl, e := model.MediaModel.SaveMedia(sizeInt, category, extension)
    if e = c.SaveUploadedFile(file, savePath); e != nil {
        LOG.Logger.Errorf("Empty File Error: %v", ErrorCode.UploadFileError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    GenSuccessReturn(&res.Result)
    res.FetchUrl = fetchUrl
    c.JSON(http.StatusOK, res)
}

// GetTotalMedia godoc
// @Summary Get Total Media
// @Description Return Media List
// @ID GetTotalMedia
// @Produce json
// @Success 200 {object} GetTotalMediaResponse
// @Router /api/admin/getTotalMedia [get]
func GetTotalMedia(c *gin.Context) {
    var res GetTotalMediaResponse
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
