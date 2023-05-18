package model

import (
    "github.com/Greetlist/CultureWeb/web_admin/server/model/schema"
)

type SubmitActivityRequest struct {
    Title string `json:"title"`
    Rank uint `json:"rank"`
    Summary string `json:"summary"`
    Labels []int `json:"labels"`
    IsTop bool `json:"is_top"`
    Content string `json:"content"`
    Author string `json:"author"`
}
type SubmitActivityResponse struct {
    Result RequestResult `json:"request_result"`
}

type ModifyActivityItem struct {
    ActivityID uint `json:"activity_id"`
    Title string `json:"title"`
    Summary string `json:"summary"`
    Rank uint `json:"rank"`
    Labels []int `json:"labels"`
    IsTop bool `json:"is_top"`
    IsEnable bool `json:"is_enable"`

    //for modify activity content
    Content string `json:"content"`
    CreateTime string `json:"create_time"`
    LocalSaveName string `json:"local_save_name"`
    IsModifyContent bool `json:"is_modify_content"`
}
type BatchModifyActivityRequest struct {
    ModifyList []ModifyActivityItem `json:"modify_list"`
}
type BatchModifyActivityResponse struct {
    Result RequestResult `json:"request_result"`
}

type BatchDeleteActivityRequest struct {
    DeleteList []uint `json:"delete_list"`
}
type BatchDeleteActivityResponse struct {
    Result RequestResult `json:"request_result"`
}

type GetTotalActivityResponse struct {
    ActivityList []schema.ActivityDetail `json:"activity_list"`
    Result RequestResult `json:"request_result"`
}

type SearchActivityRequest struct {
    KeyWord string `json:"key_word"`
}
type SearchActivityResponse struct {
    ActivityList []schema.ActivityDetail `json:"activity_list"`
    Result RequestResult `json:"request_result"`
}

type GetActivityContentRequest struct {
    CreateTime string `json:"create_time"`
    LocalSaveName string `json:"local_save_name"`
    ActivityID uint `json:"activity_id"`
    IsAdmin bool `json:"is_admin"`
}
type GetActivityContentResponse struct {
    ActivityContent string `json:"activity_content"`
    Result RequestResult `json:"request_result"`
}
