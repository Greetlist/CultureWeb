package model

import (
    "gorm.io/gorm"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    err "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

type UserModelStruct struct {
    DB *gorm.DB
}

func (user *UserModelStruct) GetUserInfo(userID uint) (schema.User, error) {
    var res schema.User
    query_res := user.DB.Where(&schema.User{ID: userID}).First(&res)
    if query_res.Error != nil {
        return res, err.GetUserInfoError
    }
    return res, nil
}

func (user *UserModelStruct) GetTotalUserInfo() (*[]schema.User, error) {
    var res []schema.User
    query_res := user.DB.Find(&res)
    if query_res.Error != nil {
        return &res, err.GetUserInfoError
    }
    return &res, nil
}

func (user *UserModelStruct) CreateUser(u *schema.User) error {
    insert_res := user.DB.Create(u)
    if insert_res.Error != nil {
        return err.RegisterUserError
    }
    return nil
}
