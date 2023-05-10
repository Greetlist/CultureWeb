package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
)

// GetUserInfo godoc
// @Summary Query User Info
// @Description Return User Info
// @ID GetUserInfo
// @Produce json
// @Param user_id query string false "user id"
// @Success 200 {object} model.GetUserInfoResponse
// @Router /api/admin/getUserInfo [get]
func GetUserInfo(c *gin.Context) {
    var req model.GetUserInfoRequest
    var res model.GetUserInfoResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    userInfo, err := model.UserModel.GetUserInfo(req.UserID)
    res.UserInfo = userInfo
    if err != nil {
        LOG.Logger.Errorf("Get User Info Error: %v", err)
        model.GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// GetTotalUserInfo godoc
// @Summary Query Total User Info
// @Description Return Total User Info
// @ID GetTotalUserInfo
// @Produce json
// @Success 200 {object} model.GetTotalUserInfoResponse
// @Router /api/admin/getTotalUserInfo [get]
func GetTotalUserInfo(c *gin.Context) {
    var res model.GetTotalUserInfoResponse
    userInfos, err := model.UserModel.GetTotalUserInfo()
    if err != nil {
        LOG.Logger.Errorf("Get User Info Error: %v", err)
        model.GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    res.UserInfos = *userInfos
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// AddUser godoc
// @Summary Add User by admin
// @Description Add User by admin
// @ID AddUser
// @Accept json
// @Produce json
// @Param request_json body model.AddUserRequest true "User Basic Info"
// @Success 200 {object} model.AddUserResponse
// @Router /api/admin/addUser [post]
func AddUser(c *gin.Context) {
    var req model.AddUserRequest
    var res model.AddUserResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    err := model.UserModel.CreateUser(&req)
    if err != nil {
        LOG.Logger.Errorf("Create User Error: %v", err)
        model.GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// UserRegister godoc
// @Summary User Register
// @Description User Register
// @ID UserRegister
// @Accept json
// @Produce json
// @Param request_json body model.AddUserRequest true "User Basic Info"
// @Success 200 {object} model.AddUserResponse
// @Router /api/user/normal/userRegister [post]
func UserRegister(c *gin.Context) {
    var req model.AddUserRequest
    var res model.AddUserResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    err := model.UserModel.CreateUser(&req)
    if err != nil {
        LOG.Logger.Errorf("Create User Error: %v", err)
        model.GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary User Login
// @Description User Login
// @ID UserLogin
// @Accept json
// @Produce json
// @Param request_json body model.UserLoginRequest true "User Login"
// @Success 200 {object} model.UserLogoutRequest
// @Router /api/user/normal/login [post]
func UserLogin(c *gin.Context) {
    var req model.UserLoginRequest
    var res model.UserLoginResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    user, err := model.UserModel.FetchUserInfo(req.Account)
    if err != nil {
        LOG.Logger.Errorf("Fetch User Error: %v", err)
        model.GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    } else if user.Password != model.GetCryptoString(req.Password) {
        model.GenErrorReturn(ErrorCode.PasswordNotCorrectError, &res.Result)
        c.JSON(ErrorCode.PasswordNotCorrectError.HttpStatusCode, res)
        return
    }

    token := model.NewToken()
    if e := model.SetTokenToRedis(token, user); e != nil {
        LOG.Logger.Errorf("Set Token Error: %v", err)
        model.GenErrorReturn(err, &res.Result)
    }
    model.GenSuccessReturn(&res.Result)
    res.UserID = user.UserID
    c.SetCookie(token.TokenName, token.Value, int(config.GlobalConfig.TokenConfig.TokenExpireTime), "/", config.GlobalConfig.CookieDomain, false, false)
    c.JSON(http.StatusOK, res)
}

// Logout godoc
// @Summary User Logout
// @Description User Logout
// @ID UserLogout
// @Accept json
// @Produce json
// @Param request_json body model.UserLogoutRequest true "User Logout"
// @Success 200 {object} model.UserLogoutResponse
// @Router /api/user/normal/logout [post]
func UserLogout(c *gin.Context) {
    var req model.UserLogoutRequest
    var res model.UserLogoutResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    cookie, _ := c.Cookie(config.GlobalConfig.TokenConfig.TokenName)
    model.CleanRedisToken(cookie)
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// Modify User Info godoc
// @Summary Modify User Info
// @Description Modify User Info
// @ID UserModify
// @Accept json
// @Produce json
// @Param request_json body model.UserModifyRequest true "modify user info"
// @Success 200 {object} model.UserModifyResponse
// @Router /api/user/auth/userModify [post]
func UserModify(c *gin.Context) {
    var req model.UserModifyRequest
    var res model.UserModifyResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    cookie, _ := c.Cookie(config.GlobalConfig.TokenConfig.TokenName)
    userRedisInfo, err := model.GetUserInfoFromRedis(cookie, nil)
    if err != nil {
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    err = model.UserModel.ModifyUser(userRedisInfo.UserID, &req)
    if err != nil {
        LOG.Logger.Errorf("Modify User Error: %v", err)
        model.GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// Change User Password godoc
// @Summary Change User Password
// @Description Modify User Password
// @ID ChangePassword
// @Accept json
// @Produce json
// @Param request_json body model.ChangePwdRequest true "change user password"
// @Success 200 {object} model.ChangePwdResponse
// @Router /api/user/auth/changePassword [post]
func ChangePassword(c *gin.Context) {
    var req model.ChangePwdRequest
    var res model.ChangePwdResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    cookie, _ := c.Cookie(config.GlobalConfig.TokenConfig.TokenName)
    userRedisInfo, err := model.GetUserInfoFromRedis(cookie, nil)
    if err != nil {
        model.GenErrorReturn(ErrorCode.GetCookieError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    err = model.UserModel.ChangePassword(userRedisInfo.UserID, req.Password)
    if err != nil {
        LOG.Logger.Errorf("Change User Password Error: %v", err)
        model.GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
