package init

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"

    "greetlist/CultureWeb/server/routers"
    _ "greetlist/CultureWeb/server/docs"
    midware "greetlist/CultureWeb/server/middleware"
)

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
    routers.InitStockApiRouter(ApiRouterGroup)
    return Router
}
