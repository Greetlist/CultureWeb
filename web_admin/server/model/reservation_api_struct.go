package model
import (
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
)

type GetTotalReservationResponse struct {
    ReservationList []schema.Reservation `json:"reservations"`
    Result RequestResult `json:"request_result"`
}

type SubmitReservationRequest struct {
    Usage string `json:"usage"`
    SiteList []int `json:"site_list"`
    StartTime string `json:"start_time"`
    EndTime string `json:"end_time"`
}
type SubmitReservationResponse struct {
    Result RequestResult `json:"request_result"`
}

type CancelReservationRequest struct {
    Reservations []int `json:"reservations"`
}
type CancelReservationResponse struct {
    Result RequestResult `json:"request_result"`
}

type ModifyReservationItem struct {
    ReservationID uint `json:"reservation_id"`
    Usage string `json:"usage"`
    StartTime string `json:"start_time"`
    EndTime string `json:"end_time"`
}
type ModifyReservationRequest struct {
    ModifyList []ModifyReservationItem `json:"modify_list"`
}
type ModifyReservationResponse struct {
    Result RequestResult `json:"request_result"`
}
