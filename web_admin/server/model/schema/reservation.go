package schema

import (
    "time"
)

type Reservation struct {
    ReservationID uint `gorm:"column:reservation_id;primary_key;AUTO_INCREMENT;" json:"reservation_id"`
    Usage string `gorm:"column:usage;not null;" json:"usage"`
    Sites []*Site `gorm:"many2many:reservation_sites;" json:"sites"`
    StartTime time.Time `gorm:"column:start_time;not null" json:"start_time"`
    EndTime time.Time `gorm:"column:end_time;not null" json:"end_time"`
    CreateTime time.Time `gorm:"column:create_time;autoCreateTime:milli;" json:"create_time"`
    UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime:milli;" json:"update_time"`
}

func (uv Reservation) TableName() string {
    return "reservation"
}
