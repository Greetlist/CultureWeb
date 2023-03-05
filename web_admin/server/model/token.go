package model

import (
    "time"

    uuid "github.com/nu7hatch/gouuid"
    "github.com/gomodule/redigo/redis"

    redisPool "github.com/Greetlist/CultureWeb/web_admin/server/redis"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
)

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
        TokenName: config.GlobalConfig.TokenConfig.TokenName,
        Value: tokenValue,
        ExpireTime: time.Now().Unix() + config.GlobalConfig.TokenConfig.TokenExpireTime,
    }
}

func VerifyToken(token string, checkAdmin bool) (bool, *ErrorCode.ResponseError) {
    redisClient, _ := <- redisPool.RedisPool
    defer redisPool.ReturnRedisClient(redisClient)

    isExists, checkExistsError := redis.Int(redisClient.Do("EXISTS", token))
    if checkExistsError != nil {
        LOG.Logger.Errorf("EXISTS ErrorCode is: %v", checkExistsError)
        return false, ErrorCode.RedisCommandError
    } else if isExists != 1 {
        LOG.Logger.Infof("Cookie: %v is not exists, user already logout.", token)
        return false, ErrorCode.CookieIsCleanedError
    }

    var saveMap RedisSaveStruct
    v, e := redis.Values(redisClient.Do("HGETALL", token))
    if e != nil {
        LOG.Logger.Errorf("HGETALL ErrorCode is: %v", e)
        return false, ErrorCode.RedisCommandError
    }
    if e := redis.ScanStruct(v, &saveMap); e != nil {
        LOG.Logger.Errorf("Parse ErrorCode is: %v", e)
        return false, ErrorCode.RedisParseStructError
    }

    currentTimestamp := time.Now().Unix()
    if currentTimestamp > saveMap.ExpireTime {
        CleanRedisToken(token)
        return false, ErrorCode.RedisKeyExpireError
    }
    if checkAdmin && !saveMap.IsAdmin {
        return false, ErrorCode.CheckAdminError
    }
    return true, nil
}

func SetTokenToRedis(ut *UserToken, user *schema.User) *ErrorCode.ResponseError {
    redisClient, _ := <- redisPool.RedisPool
    defer redisPool.ReturnRedisClient(redisClient)
    var saveMap RedisSaveStruct
    saveMap.Account = user.Account
    saveMap.UserName = user.Name
    saveMap.IsAdmin = user.IsAdmin
    saveMap.ExpireTime = ut.ExpireTime
    if _, e := redisClient.Do("HSET", redis.Args{}.Add(ut.Value).AddFlat(&saveMap)...); e != nil {
        LOG.Logger.Errorf("HSET error is: %v", e)
        return ErrorCode.RedisCommandError
    }
    return nil
}

func CleanRedisToken(token string) {
    redisClient, _ := <- redisPool.RedisPool
    defer redisPool.ReturnRedisClient(redisClient)
    if _, e := redisClient.Do("DEL", token); e != nil {
        LOG.Logger.Errorf("DEL error is: %v", e)
    }
}
