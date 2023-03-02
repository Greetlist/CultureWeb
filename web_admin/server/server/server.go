package server

import (
    initModule "github.com/Greetlist/CultureWeb/web_admin/server/init"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    "strconv"
)

func RunServer(config_file string) {
    initModule.InitAllModule(config_file)
    router := initModule.InitRouterAndMiddleware()
    //router.Run(config.GlobalConfig.BindAddr + ":" + strconv.FormatInt(config.GlobalConfig.BindPort, 10))
    addr := config.GlobalConfig.BindAddr + ":" + strconv.FormatInt(config.GlobalConfig.BindPort, 10)
    LOG.Logger.Infof("%v %v %v", addr, config.GlobalConfig.TLSConfig.ServerCrt, config.GlobalConfig.TLSConfig.ServerKey)
    router.RunTLS(addr, config.GlobalConfig.TLSConfig.ServerCrt, config.GlobalConfig.TLSConfig.ServerKey)
}
