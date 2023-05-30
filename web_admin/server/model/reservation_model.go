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

    LOG.Logger.Infof("args: %v", req)
    var resList []schema.Reservation
    //check time range conflict
    query_res := reservation.DB.Debug().Where("start_time >= ? AND end_time <= ?", startTime, endTime).Find(&resList)
    if query_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", query_res.Error)
        return ErrorCode.SubmitReservationError
    }
    if query_res.RowsAffected != 0 {
        var siteList []schema.Site
        for _, item := range(resList) {
            reservation.DB.Debug().Model(&item).Association("Sites").Find(&siteList)
            for _, s := range(siteList) {
                if req.Site == s.SiteID {
                    LOG.Logger.Errorf("Duplicate Time Range: %v", resList)
                    return ErrorCode.ReservationDupError
                }
            }
        }
    }

    var siteList []*schema.Site
    var site schema.Site
    site.SiteID = req.Site
    siteList = append(siteList, &site)

    create_res := reservation.DB.Create(
        &schema.Reservation{
            Usage: strings.Trim(req.Usage, " "),
            Sites: siteList,
            StartTime: startTime,
            EndTime: endTime})
    if create_res.Error != nil {
        LOG.Logger.Errorf("DB Error: %v", create_res.Error)
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
