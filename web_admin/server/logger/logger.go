package logger

import (
    "github.com/phachon/go-logger"
    "github.com/Greetlist/CultureWeb/web_admin/server/config"
    "path"
    "os"
)

var LOG *go_logger.Logger

func InitLogger() {
    LOG = go_logger.NewLogger()
    os.MkdirAll(config.GlobalConfig.LogDir, os.ModePerm)
    LOG.Detach("console")

    consoleConfig := &go_logger.ConsoleConfig{
        Color: true,
        JsonFormat: false,
        Format: "%timestamp_format% [%level_string%][%file%:%line%] %body%",
    }
    // 添加 console 为 logger 的一个输出
    LOG.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)

    fileConfig := &go_logger.FileConfig {
        Filename: path.Join(config.GlobalConfig.LogDir, "total.log"),
        LevelFileName : map[int]string {
            LOG.LoggerLevel("info"): path.Join(config.GlobalConfig.LogDir, "info.log"),
            LOG.LoggerLevel("alert"): path.Join(config.GlobalConfig.LogDir, "alert.log"),
            LOG.LoggerLevel("warning"): path.Join(config.GlobalConfig.LogDir, "warning.log"),
            LOG.LoggerLevel("error"): path.Join(config.GlobalConfig.LogDir, "error.log"),
            LOG.LoggerLevel("debug"): path.Join(config.GlobalConfig.LogDir, "debug.log"),
        },
        MaxSize : 0,
        MaxLine : 0,
        DateSlice : "H",
        JsonFormat: false,
        Format: "%timestamp_format% [%level_string%][%file%:%line%] %body%",
    }
    LOG.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
}
