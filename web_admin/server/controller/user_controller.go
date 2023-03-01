package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

// GetUserInfo godoc
// @Summary Query User Info
// @Description Return User Info
// @ID GetUserInfo
// @Produce json
// @Param user_id query string false "user id"
// @Success 200 {object} GetUserInfoResponse
// @Router /api/user/getUserInfo [get]
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
// @Router /api/user/getTotalUserInfo [get]
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

// UserRegister godoc
// @Summary User Register
// @Description User Register
// @ID UserRegister
// @Accept json
// @Produce json
// @Param request_json body UserRegisterRequest true "User Basic Info"
// @Success 200 {object} UserRegisterResponse
// @Router /api/user/userRegister [post]
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
    user.Sex = req.Sex
    user.Age = req.Age
    return &user
}
func genModifyUserSchema (req *UserModifyRequest) *schema.User {
    var user schema.User
    user.Name = req.Name
    user.Sex = req.Sex
    user.Age = req.Age
    return &user
}

// Modify User Info godoc
// @Summary Modify User Info
// @Description Modify User Info
// @ID UserModify
// @Accept json
// @Produce json
// @Param request_json body UserModifyRequest true "modify user info"
// @Success 200 {object} UserModifyResponse
// @Router /api/user/userModify [post]
func UserModify(c *gin.Context) {
    var req UserModifyRequest
    var res UserModifyResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }
    err := model.UserModel.ModifyUser(req.UserID, genModifyUserSchema(&req))
    if err != nil {
        LOG.Logger.Errorf("Modify User Error: %v", err)
        GenErrorReturn(err, &res.Result)
        c.JSON(err.HttpStatusCode, res)
        return
    }
    GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// Login godoc
// @Summary User Login
// @Description User Login
// @ID UserLogin
// @Accept json
// @Produce json
// @Param request_json body UserLoginRequest true "User Login"
// @Success 200 {object} UserLogoutRequest
// @Router /api/user/userLogin [post]
func UserLogin(c *gin.Context) {
}

// Logout godoc
// @Summary User Logout
// @Description User Logout
// @ID UserLogout
// @Accept json
// @Produce json
// @Param request_json body UserLogoutRequest true "User Logout"
// @Success 200 {object} UserLogoutResponse
// @Router /api/user/userLogout [post]
func UserLogout(c *gin.Context) {
}
