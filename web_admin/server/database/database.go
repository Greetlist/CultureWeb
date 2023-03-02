package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    "time"
    "fmt"
    "os"
)

var DB *gorm.DB

func InitDB() {
    connStr := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?%s",
        config.GlobalConfig.MysqlConfig.User, config.GlobalConfig.MysqlConfig.Password,
        config.GlobalConfig.MysqlConfig.Host, config.GlobalConfig.MysqlConfig.Port,
        config.GlobalConfig.MysqlConfig.DBName, config.GlobalConfig.MysqlConfig.Params,
    )
    fmt.Printf("%s", connStr)
    var err error
    DB, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})
    if err != nil {
        fmt.Printf("Open DB Error: [%v]", err)
        os.Exit(-1)
    }

    mysql, _ := DB.DB()
    mysql.SetMaxIdleConns(config.GlobalConfig.MysqlConfig.MaxIdleConns)
    mysql.SetMaxOpenConns(config.GlobalConfig.MysqlConfig.MaxOpenConns)
    mysql.SetConnMaxLifetime(time.Hour)
}

func AutoMigrate() {
    DB.AutoMigrate(
        &schema.User{}, &schema.UserVerification{},
        &schema.Media{}, &schema.Interest{},
        &schema.Article{},
    )
}
