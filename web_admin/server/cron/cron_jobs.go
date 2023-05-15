package cronjobs

import (
    "github.com/robfig/cron/v3"
)

func init() {
    c := cron.New()
    c.AddFunc("0 30 * * * *", model.CleanAllExpireRedisKey())
    c.Start()
}
