package model

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/database"
)

var (
    UserModel UserModelStruct
)

func InitModel() {
    UserModel = UserModelStruct{DB: database.DB}
}
