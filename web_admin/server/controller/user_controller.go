package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

// GetUserInfo godoc
// @Summary Query User Info
// @Description Return User Info
// @ID GetUserInfo
// @Produce json
// @Param user_id query string false "user id"
// @Success 200 {object} model.GetUserInfoResponse
// @Router /api/user/getUserInfo [get]
func GetUserInfo(context *gin.Context) {
}

// GetTotalUserInfo godoc
// @Summary Query Total User Info
// @Description Return Total User Info
// @ID GetTotalUserInfo
// @Produce json
// @Success 200 {object} model.GetTotalUserInfoResponse
// @Router /api/user/getTotalUserInfo [get]
func GetTotalUserInfo(context *gin.Context) {
    err := ErrorCode.ParseParamError
    LOG.Logger.Infof("%v", err.Message)
    var response model.GetTotalUserInfoResponse
    context.JSON(http.StatusOK, response)
}

// UserRegister godoc
// @Summary User Register
// @Description User Register
// @ID UserRegister
// @Accept json
// @Produce json
// @Param request_json body model.UserRegisterRequest true "User Basic Info"
// @Success 200 {object} model.UserRegisterResponse
// @Router /api/user/userRegister [post]
func UserRegister(context *gin.Context) {
    request := model.UserRegisterRequest{}
    if err := context.BindJSON(&request); err != nil {
        LOG.Logger.Errorf("Register Params error: %v", err)
	    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    var response model.RequestResult
    context.JSON(http.StatusOK, response)
}

// Login godoc
// @Summary User Login
// @Description User Login
// @ID UserLogin
// @Accept json
// @Produce json
// @Param request_json body model.UserLoginRequest true "User Login"
// @Success 200 {object} model.UserLogoutRequest
// @Router /api/user/userLogin [post]
func UserLogin(context *gin.Context) {
}

// Logout godoc
// @Summary User Logout
// @Description User Logout
// @ID UserLogout
// @Accept json
// @Produce json
// @Param request_json body model.UserLogoutRequest true "User Logout"
// @Success 200 {object} model.UserLogoutResponse
// @Router /api/user/userLogout [post]
func UserLogout(context *gin.Context) {
}
