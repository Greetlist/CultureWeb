package cronjobs

import (
    "github.com/robfig/cron/v3"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
)

func init() {
    c := cron.New()
    c.AddFunc("@every 10m", func() {
        model.ArticleModel.SaveArticleVisitNumToDB()
    })
    c.Start()
}
