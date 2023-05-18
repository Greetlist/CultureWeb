package init

import (
    "net/http"
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
    "github.com/Greetlist/CultureWeb/web_admin/server/redis"
    _ "github.com/Greetlist/CultureWeb/web_admin/server/cron"
)

func InitAllModule(config_file string) {
    config.InitConfig(config_file)
    logger.InitLogger()
    database.InitDB()
    database.AutoMigrate()
    model.InitModel(database.DB)
    redis.InitRedisChannel()
}

func InitRouterAndMiddleware() *gin.Engine {
    Router := gin.New()
    Router.Use(gin.Logger())
    Router.Use(gin.Recovery())

    //Middleware
    Router.Use(midware.SetCORSHeader())

    //API Doc
    Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    //Static media, jpg,png,mp4 etc
    Router.StaticFS(config.GlobalConfig.MediaBaseUrl, http.Dir(config.GlobalConfig.MediaSaveDir))

    //前端展示路由,需要跟vue集成
    //WebInterfaceGroup := Router.Group("")

    ApiRouterGroup := Router.Group("api")
    routers.InitAdminUserApiRouter(ApiRouterGroup)
    routers.InitAdminArticleApiRouter(ApiRouterGroup)
    routers.InitNormalUserApiRouter(ApiRouterGroup)
    routers.InitAdminLabelApiRouter(ApiRouterGroup)
    routers.InitBasicInfoRouter(ApiRouterGroup)
    routers.InitAdminActivityApiRouter(ApiRouterGroup)
    return Router
}
