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
    UserID uint `redis:"user_id"`
    Account string `redis:"account"`
    UserName string `redis:"user_name"`
    IsAdmin bool `redis:"is_admin"`
    ExpireTime int64 `redis:"expire_time"`
}

func NewToken() *UserToken {
    uid, _ := uuid.NewV4()
    tokenValue := GetCryptoString(uid.String())
    return &UserToken {
        TokenName: config.GlobalConfig.TokenConfig.TokenName,
        Value: tokenValue,
        ExpireTime: time.Now().Unix() + config.GlobalConfig.TokenConfig.TokenExpireTime,
    }
}

func VerifyToken(token string, checkAdmin bool) (bool, *ErrorCode.ResponseError) {
    saveMap, getErr := GetUserInfoFromRedis(token, nil)
    if getErr != nil {
        return false, getErr
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
    saveMap.UserID = user.UserID
    saveMap.Account = user.Account
    saveMap.UserName = user.Name
    saveMap.IsAdmin = user.IsAdmin
    saveMap.ExpireTime = ut.ExpireTime
    if _, e := redisClient.Do("HSET", redis.Args{}.Add(ut.Value).AddFlat(&saveMap)...); e != nil {
        LOG.Logger.Errorf("HSET error is: %v", e)
        return ErrorCode.RedisCommandError
    }
    if _, e := redisClient.Do("EXPIRE", ut.Value, config.GlobalConfig.TokenConfig.TokenExpireTime); e != nil {
        LOG.Logger.Errorf("EXPIRE error is: %v, delete current key", e)
        redisClient.Do("DEL", ut.Value)
        return ErrorCode.RedisCommandError
    }
    return nil
}

func CleanAllExpireRedisKey() {
    redisClient, _ := <- redisPool.RedisPool
    defer redisPool.ReturnRedisClient(redisClient)
    keys, e := redis.Strings(redisClient.Do("KEYS", "*"))
    if e != nil {
        LOG.Logger.Errorf("Get all Keys error is: %v", e)
        return
    }
    for _, key := range keys {
        saveMap, e := GetUserInfoFromRedis(key, redisClient)
        if e != nil {
            LOG.Logger.Errorf("Get token info error is: %v", e)
            continue
        }
        currentTimestamp := time.Now().Unix()
        if currentTimestamp > saveMap.ExpireTime {
            CleanRedisToken(key)
        }
    }
}

func CleanRedisToken(token string) {
    redisClient, _ := <- redisPool.RedisPool
    defer redisPool.ReturnRedisClient(redisClient)
    if _, e := redisClient.Do("DEL", token); e != nil {
        LOG.Logger.Errorf("DEL error is: %v", e)
    }
}

func RefreshRedisToken(token string) *ErrorCode.ResponseError {
    redisClient, _ := <- redisPool.RedisPool
    defer redisPool.ReturnRedisClient(redisClient)
    newExpireTime := time.Now().Unix() + config.GlobalConfig.TokenConfig.TokenExpireTime
    if _, e := redisClient.Do("HSET", token, "expire_time", newExpireTime); e != nil {
        LOG.Logger.Errorf("HSET error is: %v", e)
        return ErrorCode.RedisCommandError
    }
    return nil
}

func GetUserInfoFromRedis(token string, rc redis.Conn) (*RedisSaveStruct, *ErrorCode.ResponseError) {
    var redisClient redis.Conn
    if rc == nil {
        redisClient, _ = <- redisPool.RedisPool
        defer redisPool.ReturnRedisClient(redisClient)
    } else {
        redisClient = rc
    }

    isExists, checkExistsError := redis.Int(redisClient.Do("EXISTS", token))
    if checkExistsError != nil {
        LOG.Logger.Errorf("EXISTS ErrorCode is: %v", checkExistsError)
        return nil, ErrorCode.RedisCommandError
    } else if isExists != 1 {
        LOG.Logger.Infof("Cookie: %v is not exists, user already logout.", token)
        return nil, ErrorCode.CookieIsCleanedError
    }
    var saveMap RedisSaveStruct
    v, e := redis.Values(redisClient.Do("HGETALL", token))
    if e != nil {
        LOG.Logger.Errorf("HGETALL ErrorCode is: %v", e)
        return nil, ErrorCode.RedisCommandError
    }
    if e := redis.ScanStruct(v, &saveMap); e != nil {
        LOG.Logger.Errorf("Parse ErrorCode is: %v", e)
        return nil, ErrorCode.RedisParseStructError
    }
    return &saveMap, nil
}
