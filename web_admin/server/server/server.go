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

    LOG.Logger.Infof("Start No-TLS Server at: %v:%v", config.GlobalConfig.BindAddr, config.GlobalConfig.BindPort)
    router.Run(config.GlobalConfig.BindAddr + ":" + strconv.FormatInt(config.GlobalConfig.BindPort, 10))
}

func RunTLSServer(config_file string) {
    initModule.InitAllModule(config_file)
    router := initModule.InitRouterAndMiddleware()

    LOG.Logger.Infof("Start TLS Server at: %v:%v", config.GlobalConfig.BindAddr, config.GlobalConfig.BindPort)
    addr := config.GlobalConfig.BindAddr + ":" + strconv.FormatInt(config.GlobalConfig.BindPort, 10)
    router.RunTLS(addr, config.GlobalConfig.TLSConfig.ServerCrt, config.GlobalConfig.TLSConfig.ServerKey)
}
