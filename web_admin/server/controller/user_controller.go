package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
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
// @Success 200 {object} GetUserInfoResponse
// @Router /api/admin/getUserInfo [get]
func GetUserInfo(c *gin.Context) {
    var req GetUserInfoRequest
    var res GetUserInfoResponse
    if e := c.ShouldBind(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    userInfo, err := model.UserModel.GetUserInfo(req.UserID)
    res.UserInfo = userInfo
    if err != nil {
        LOG.Logger.Errorf("Get User Info Error: %v", err)
        GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// GetTotalUserInfo godoc
// @Summary Query Total User Info
// @Description Return Total User Info
// @ID GetTotalUserInfo
// @Produce json
// @Success 200 {object} GetTotalUserInfoResponse
// @Router /api/admin/getTotalUserInfo [get]
func GetTotalUserInfo(c *gin.Context) {
    var res GetTotalUserInfoResponse
    userInfos, err := model.UserModel.GetTotalUserInfo()
    if err != nil {
        LOG.Logger.Errorf("Get User Info Error: %v", err)
        GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    res.UserInfos = *userInfos
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// AddUser godoc
// @Summary Add User by admin
// @Description Add User by admin
// @ID AddUser
// @Accept json
// @Produce json
// @Param request_json body AddUserRequest true "User Basic Info"
// @Success 200 {object} AddUserResponse
// @Router /api/admin/addUser [post]
func AddUser(c *gin.Context) {
    var req AddUserRequest
    var res AddUserResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    err := model.UserModel.CreateUser(genAddUserSchema(&req))
    if err != nil {
        LOG.Logger.Errorf("Create User Error: %v", err)
        GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// UserRegister godoc
// @Summary User Register
// @Description User Register
// @ID UserRegister
// @Accept json
// @Produce json
// @Param request_json body UserRegisterRequest true "User Basic Info"
// @Success 200 {object} UserRegisterResponse
// @Router /api/user/normal/userRegister [post]
func UserRegister(c *gin.Context) {
    var req UserRegisterRequest
    var res UserRegisterResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    err := model.UserModel.CreateUser(genCreateUserSchema(&req))
    if err != nil {
        LOG.Logger.Errorf("Create User Error: %v", err)
        GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

func genCreateUserSchema (req *UserRegisterRequest) *schema.User {
    var user schema.User
    user.Account = req.Account
    user.Password = req.Password
    user.Name = req.Name
    user.PhoneNumber = req.PhoneNumber
    user.Sex = req.Sex
    user.Age = req.Age
    return &user
}
func genAddUserSchema (req *AddUserRequest) *schema.User {
    var user schema.User
    user.Account = req.Account
    user.Password = req.Password
    user.Name = req.Name
    user.PhoneNumber = req.PhoneNumber
    user.Sex = req.Sex
    user.Age = req.Age
    user.Address = req.Address
    user.IdentifyID = req.IdentifyID
    user.IsActive = req.IsActive
    user.IsAdmin = req.IsAdmin
    user.State = req.State
    return &user
}
func genModifyUserSchema (req *UserModifyRequest) *schema.User {
    var user schema.User
    user.Name = req.Name
    user.Password = req.Password
    user.Sex = req.Sex
    user.Age = req.Age
    return &user
}

// Login godoc
// @Summary User Login
// @Description User Login
// @ID UserLogin
// @Accept json
// @Produce json
// @Param request_json body UserLoginRequest true "User Login"
// @Success 200 {object} UserLogoutRequest
// @Router /api/user/normal/login [post]
func UserLogin(c *gin.Context) {
    var req UserLoginRequest
    var res UserLoginResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    user, err := model.UserModel.FetchUserInfo(req.Account)
    if err != nil {
        LOG.Logger.Errorf("Fetch User Error: %v", err)
        GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    } else if user.Password != model.GetCryptoString(req.Password) {
        GenErrorReturn(ErrorCode.PasswordNotCorrectError, &res.Result)
        c.JSON(ErrorCode.PasswordNotCorrectError.HttpStatusCode, res)
        return
    }

    token := model.NewToken()
    if e := model.SetTokenToRedis(token, user); e != nil {
        LOG.Logger.Errorf("Set Token Error: %v", err)
        GenErrorReturn(err, &res.Result)
    }
    GenSuccessReturn(&res.Result)
    res.UserID = user.UserID
    c.SetCookie(token.TokenName, token.Value, int(config.GlobalConfig.TokenConfig.TokenExpireTime), "/", "192.168.199.123", true, false)
    c.JSON(http.StatusOK, res)
}

// Logout godoc
// @Summary User Logout
// @Description User Logout
// @ID UserLogout
// @Accept json
// @Produce json
// @Param request_json body UserLogoutRequest true "User Logout"
// @Success 200 {object} UserLogoutResponse
// @Router /api/user/normal/logout [post]
func UserLogout(c *gin.Context) {
    var req UserLogoutRequest
    var res UserLogoutResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    cookie, _ := c.Cookie(config.GlobalConfig.TokenConfig.TokenName)
    model.CleanRedisToken(cookie)
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// Modify User Info godoc
// @Summary Modify User Info
// @Description Modify User Info
// @ID UserModify
// @Accept json
// @Produce json
// @Param request_json body UserModifyRequest true "modify user info"
// @Success 200 {object} UserModifyResponse
// @Router /api/user/auth/userModify [post]
func UserModify(c *gin.Context) {
    var req UserModifyRequest
    var res UserModifyResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    cookie, _ := c.Cookie(config.GlobalConfig.TokenConfig.TokenName)
    userRedisInfo, err := model.GetUserInfoFromRedis(cookie, nil)
    if err != nil {
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    err = model.UserModel.ModifyUser(userRedisInfo.UserID, genModifyUserSchema(&req))
    if err != nil {
        LOG.Logger.Errorf("Modify User Error: %v", err)
        GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// Change User Password godoc
// @Summary Change User Password
// @Description Modify User Password
// @ID ChangePassword
// @Accept json
// @Produce json
// @Param request_json body ChangePwdRequest true "change user password"
// @Success 200 {object} ChangePwdResponse
// @Router /api/user/auth/changePassword [post]
func ChangePassword(c *gin.Context) {
    var req ChangePwdRequest
    var res ChangePwdResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    cookie, _ := c.Cookie(config.GlobalConfig.TokenConfig.TokenName)
    userRedisInfo, err := model.GetUserInfoFromRedis(cookie, nil)
    if err != nil {
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    err = model.UserModel.ChangePassword(userRedisInfo.UserID, req.Password)
    if err != nil {
        LOG.Logger.Errorf("Change User Password Error: %v", err)
        GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
