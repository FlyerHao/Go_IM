package dao

import "time"

type GroupHistoryMessage struct {
    UserId int64 `json:"user_id"`
    GroupId int64 `json:"group_id"`
    Message string `json:"message"`
    Created time.Time `json:"created"`
}
