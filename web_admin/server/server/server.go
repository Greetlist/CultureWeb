package server

import (
    initModel "greetlist/CultureWeb/server/init"
    "strconv"
)

func RunServer(bindAddr string, bindPort int64) {
    router := initModel.InitRouterAndMiddleware()
    router.Run(bindAddr + ":" + strconv.FormatInt(bindPort, 10))
}
