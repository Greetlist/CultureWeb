package server

import (
    initModel "github.com/Greetlist/CultureWeb/web_admin/server/init"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    "strconv"
)

func RunServer(config_file string) {
    config.InitConfig(config_file)
    logger.InitLogger()
    router := initModel.InitRouterAndMiddleware()
    router.Run(config.GlobalConfig.BindAddr + ":" + strconv.FormatInt(config.GlobalConfig.BindPort, 10))
}