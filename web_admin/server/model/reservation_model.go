package model

import (
    "gorm.io/gorm"
    "time"
    "strings"
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

type ReservationModelStruct struct {
    DB *gorm.DB
}

func (reservation *ReservationModelStruct) GetReservationList(response *GetTotalReservationResponse) *ErrorCode.ResponseError {
    query_res := reservation.DB.Model(schema.Reservation{}).Preload("Sites").Find(&response.ReservationList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.GetReservationError
    }
    return nil
}

func (reservation *ReservationModelStruct) SubmitReservation(req *SubmitReservationRequest) *ErrorCode.ResponseError {
    layout := "2006-01-02T15:04:05.000Z"
    startTime, _ := time.Parse(layout, req.StartTime)
    endTime, _ := time.Parse(layout, req.EndTime)
    var siteList []*schema.Site
    for _, v := range(req.SiteList) {
        var site schema.Site
        site.SiteID = v
        siteList = append(siteList, &site)
    }
    query_res := reservation.DB.Create(
        &schema.Reservation{
            Usage: strings.Trim(req.Usage, " "),
            Sites: siteList,
            StartTime: startTime,
            EndTime: endTime})
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.SubmitReservationError
    }
    return nil
}

func (reservation *ReservationModelStruct) CancelReservation(reservations *[]int) *ErrorCode.ResponseError {
    return nil
}

func (reservation *ReservationModelStruct) ModifyReservation(mList *[]ModifyReservationItem) *ErrorCode.ResponseError {
    tx := reservation.DB.Begin()
    for _, item := range(*mList) {
        layout := "2006-01-02T15:04:05.000Z"
        startTime, _ := time.Parse(layout, item.StartTime)
        endTime, _ := time.Parse(layout, item.EndTime)
        var reservation schema.Reservation
        reservation.ReservationID = item.ReservationID
        update_res := tx.Model(&reservation).Updates(
            schema.Reservation{Usage: strings.Trim(item.Usage, " "),  StartTime: startTime, EndTime: endTime})
        if update_res.Error != nil {
            LOG.Logger.Errorf("DB Error: %v", update_res.Error)
            tx.Rollback()
            return ErrorCode.ModifyReservationError
        }
    }
    if e := tx.Commit().Error; e != nil {
        LOG.Logger.Errorf("DB Error: %v", e)
        tx.Rollback()
        return ErrorCode.ModifyReservationError
    }
    return nil
}
