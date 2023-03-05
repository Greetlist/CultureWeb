package model

import (
    "gorm.io/gorm"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    err "github.com/Greetlist/CultureWeb/web_admin/server/error"
    "crypto/sha256"
    "encoding/hex"
)

type UserModelStruct struct {
    DB *gorm.DB
}

func (user *UserModelStruct) GetUserInfo(userID uint) (schema.User, *err.ResponseError) {
    var res schema.User
    query_res := user.DB.Where(&schema.User{UserID: userID}).First(&res)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return res, err.GetUserInfoError
    }
    return res, nil
}

func (user *UserModelStruct) GetTotalUserInfo() (*[]schema.User, *err.ResponseError) {
    var res []schema.User
    query_res := user.DB.Find(&res)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return &res, err.GetUserInfoError
    }
    return &res, nil
}

func (user *UserModelStruct) FetchUserInfo(account, password string) (*schema.User, *err.ResponseError) {
    var res schema.User
    query_res := user.DB.Where(&schema.User{Account: account, Password: getCryptoString(password)}).First(&res)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return &res, err.GetUserInfoError
    }
    return &res, nil
}

func (user *UserModelStruct) CreateUser(u *schema.User) *err.ResponseError {
    u.Password = getCryptoString(u.Password)
    insert_res := user.DB.Create(u)
    if insert_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", insert_res.Error)
        return err.RegisterUserError
    }
    return nil
}

func (user *UserModelStruct) ModifyUser(userID uint, u *schema.User) *err.ResponseError {
    u.UserID = userID
    newPassword := getCryptoString(u.Password)
    update_res := user.DB.Model(u).Updates(schema.User{Name: u.Name, Password: newPassword, Sex: u.Sex, Age: u.Age})
    if update_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", update_res.Error)
        return err.ModifyUserError
    }
    return nil
}

func getCryptoString(target string) string {
    h := sha256.New()
    h.Write([]byte(target))
    cryptoBytes := h.Sum(nil)
    return string(hex.EncodeToString(cryptoBytes))
}
