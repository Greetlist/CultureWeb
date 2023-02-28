package init

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    "github.com/Greetlist/CultureWeb/web_admin/server/routers"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/docs"
    midware "github.com/Greetlist/CultureWeb/web_admin/server/middleware"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    "github.com/Greetlist/CultureWeb/web_admin/server/database"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
)

func InitAllModule(config_file string) {
    config.InitConfig(config_file)
    logger.InitLogger()
    database.InitDB()
    database.AutoMigrate()
    model.InitModel(database.DB)
}

func InitRouterAndMiddleware() *gin.Engine {
    //全局性设置
    Router := gin.New()
    Router.Use(gin.Logger())
    Router.Use(gin.Recovery())
    //跨域请求
    Router.Use(midware.SetCORSHeader())
    //API Doc
    Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    //前端展示路由,需要跟vue集成
    //WebInterfaceGroup := Router.Group("")

    /*
      功能API路由设置
      API 相关组
    */
    ApiRouterGroup := Router.Group("api")
    routers.InitUserApiRouter(ApiRouterGroup)
    return Router
}
