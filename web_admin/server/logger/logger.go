package logger

import (
    "github.com/phachon/go-logger"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    "path"
    "os"
)

var Logger *go_logger.Logger

func InitLogger() {
    Logger = go_logger.NewLogger()
    os.MkdirAll(config.GlobalConfig.LogDir, os.ModePerm)
    Logger.Detach("console")

    consoleConfig := &go_logger.ConsoleConfig{
        Color: true,
        JsonFormat: false,
        Format: "%timestamp_format% [%level_string%][%file%:%line%] %body%",
    }
    // 添加 console 为 logger 的一个输出
    Logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)

    fileConfig := &go_logger.FileConfig {
        Filename: path.Join(config.GlobalConfig.LogDir, "total.log"),
        LevelFileName : map[int]string {
            Logger.LoggerLevel("info"): path.Join(config.GlobalConfig.LogDir, "info.log"),
            Logger.LoggerLevel("alert"): path.Join(config.GlobalConfig.LogDir, "alert.log"),
            Logger.LoggerLevel("warning"): path.Join(config.GlobalConfig.LogDir, "warning.log"),
            Logger.LoggerLevel("error"): path.Join(config.GlobalConfig.LogDir, "error.log"),
            Logger.LoggerLevel("debug"): path.Join(config.GlobalConfig.LogDir, "debug.log"),
        },
        MaxSize : 0,
        MaxLine : 0,
        DateSlice : "h",
        JsonFormat: false,
        Format: "%timestamp_format% [%level_string%][%file%:%line%] %body%",
    }
    Logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
}
