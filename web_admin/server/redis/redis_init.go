package redis

import (
    "github.com/gomodule/redigo/redis"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    "strconv"
)

var RedisPool chan redis.Conn

func InitRedisChannel() {
    RedisPool = make(chan redis.Conn, config.GlobalConfig.RedisConfig.PoolSize)
    LOG.Logger.Infof("pool size: %d", len(RedisPool))
    for i := 0; i < config.GlobalConfig.RedisConfig.PoolSize; i++ {
        server_addr := config.GlobalConfig.RedisConfig.RedisServer + ":" + strconv.FormatInt(config.GlobalConfig.RedisConfig.RedisPort, 10)
        cur_conn, err := redis.Dial("tcp", server_addr)
        if err != nil {
            LOG.Logger.Errorf("Connect to Redis fail: %v", err)
        } else {
            LOG.Logger.Infof("Push a new client")
            RedisPool <- cur_conn
        }
    }
}

func ReturnRedisClient(conn redis.Conn) {
    RedisPool <- conn
}
