package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/Greetlist/CultureWeb/web_admin/server/model"
    LOG "github.com/Greetlist/CultureWeb/web_admin/server/logger"
    ErrorCode "github.com/Greetlist/CultureWeb/web_admin/server/error"
)

// GetTotalReservation godoc
// @Summary Query Total Reservation
// @Description Return Total Reservation
// @ID GetTotalReservation
// @Produce json
// @Success 200 {object} model.GetTotalReservationResponse
// @Router /api/admin/getTotalReservation [get]
func GetTotalReservation(c *gin.Context) {
    var res model.GetTotalReservationResponse
    if e := model.ReservationModel.GetReservationList(&res); e != nil {
        model.GenErrorReturn(ErrorCode.GetReservationError, &res.Result)
        c.JSON(ErrorCode.GetReservationError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// SubmitReservation godoc
// @Summary Add Reservation
// @Description Add New Reservation
// @ID SubmitReservation
// @Produce json
// @Param request_json body model.SubmitReservationRequest true "reservation Form"
// @Success 200 {object} model.SubmitReservationResponse
// @Router /api/admin/submitReservation [post]
func SubmitReservation(c *gin.Context) {
    var req model.SubmitReservationRequest
    var res model.SubmitReservationResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    if e := model.ReservationModel.SubmitReservation(&req); e != nil {
        model.GenErrorReturn(ErrorCode.SubmitReservationError, &res.Result)
        c.JSON(ErrorCode.SubmitReservationError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// CancelReservation godoc
// @Summary Cancel Reservation
// @Description Cancel Reservation
// @ID CancelReservation
// @Produce json
// @Param request_json body model.CancelReservationRequest true "delete reservation list"
// @Success 200 {object} model.CancelReservationResponse
// @Router /api/admin/deleteReservation [post]
func CancelReservation(c *gin.Context) {
    var req model.CancelReservationRequest
    var res model.CancelReservationResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    LOG.Logger.Infof("%v", req)
    if e := model.ReservationModel.CancelReservation(&req.Reservations); e != nil {
        model.GenErrorReturn(ErrorCode.CancelReservationError, &res.Result)
        c.JSON(ErrorCode.CancelReservationError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}

// ModifyReservation godoc
// @Summary Modify Reservation
// @Description Modify Reservation
// @ID ModifyReservation
// @Produce json
// @Param request_json body model.ModifyReservationRequest true "modify reservation list"
// @Success 200 {object} model.ModifyReservationResponse
// @Router /api/admin/batchModifyReservation [post]
func BatchModifyReservation(c *gin.Context) {
    var req model.ModifyReservationRequest
    var res model.ModifyReservationResponse
    if e := c.ShouldBindJSON(&req); e != nil {
        LOG.Logger.Errorf("Parse Param Error: %v", ErrorCode.ParseParamError)
        model.GenErrorReturn(ErrorCode.ParseParamError, &res.Result)
        c.JSON(ErrorCode.ParseParamError.HttpStatusCode, res)
        return
    }

    if e := model.ReservationModel.ModifyReservation(&req.ModifyList, &res); e != nil {
        model.GenErrorReturn(ErrorCode.ModifyReservationError, &res.Result)
        c.JSON(ErrorCode.ModifyReservationError.HttpStatusCode, res)
        return
    }
    model.GenSuccessReturn(&res.Result)
    c.JSON(http.StatusOK, res)
}
