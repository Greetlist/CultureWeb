package model

import (
    "gorm.io/gorm"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    err "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

type StatisticModelStruct struct {
    DB *gorm.DB
}

func (s *StatisticModelStruct) InitStatisticTable() *err.ResponseError {
    for id, name := range schema.StatisticTypeMap {
        statistic := schema.Statistic{StatisticID: id, StatisticName: name, StatisticCount: 0}
        create_result := s.DB.Create(statistic)
        if create_result.Error != nil {
            LOG.Logger.Errorf("DB Error: %v", create_result.Error)
        }
    }
    return nil
}
