package model

import (
    redisPool "github.com/Greetlist/CultureWeb/web_admin/server/redis"
    "github.com/gomodule/redigo/redis"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    err "github.com/Greetlist/CultureWeb/web_admin/server/error"
    uuid "github.com/nu7hatch/gouuid"
    "time"
)

const TOKEN_NAME = "culture_web"
const TOKEN_EXPIRE_TIME = 8 * 3600

type UserToken struct {
    TokenName string
    Value string
    ExpireTime int64
}

type RedisSaveStruct struct {
    Account string `redis:"account"`
    UserName string `redis:"user_name"`
    IsAdmin bool `redis:"is_admin"`
    ExpireTime int64 `redis:"expire_time"`
}

func NewToken() *UserToken {
    uid, _ := uuid.NewV4()
    tokenValue := getCryptoString(uid.String())
    return &UserToken {
        TokenName: TOKEN_NAME,
        Value: tokenValue,
        ExpireTime: time.Now().Unix() + TOKEN_EXPIRE_TIME,
    }
}

func VerifyToken(token string) (bool, *err.ResponseError) {
    redisClient, _ := <- redisPool.RedisPool
    var saveMap RedisSaveStruct
    v, e := redis.Values(redisClient.Do("HGETALL", token))
    if e != nil {
        LOG.Logger.Errorf("HGETALL error is: %v", e)
        return false, err.RedisCommandError
    }
    if e := redis.ScanStruct(v, &saveMap); e != nil {
        LOG.Logger.Errorf("Parse error is: %v", e)
        return false, err.RedisParseStructError
    }

    currentTimestamp := time.Now().Unix()
    if currentTimestamp > saveMap.ExpireTime {
        return false, err.RedisKeyExpireError
    }
    return true, nil
}

func SetTokenToRedis(ut *UserToken, user *schema.User) *err.ResponseError {
    redisClient, _ := <- redisPool.RedisPool
    var saveMap RedisSaveStruct
    saveMap.Account = user.Account
    saveMap.UserName = user.Name
    saveMap.IsAdmin = user.IsAdmin
    saveMap.ExpireTime = ut.ExpireTime
    if _, e := redisClient.Do("HSET", redis.Args{}.Add(ut.Value).AddFlat(&saveMap)...); e != nil {
        return err.RedisCommandError
    }
    redisPool.RedisPool <- redisClient
    return nil
}
