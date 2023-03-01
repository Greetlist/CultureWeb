package server

import (
    initModule "github.com/Greetlist/CultureWeb/web_admin/server/init"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    "strconv"
)

func RunServer(config_file string) {
    initModule.InitAllModule(config_file)
    router := initModule.InitRouterAndMiddleware()
    router.Run(config.GlobalConfig.BindAddr + ":" + strconv.FormatInt(config.GlobalConfig.BindPort, 10))
}
