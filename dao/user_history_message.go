package dao

import "time"

type UserHistoryMessage struct {
    UserId int64 `json:"user_id"`
    ToUser int64 `json:"to_user"`
    Message string `json:"message"`
    Created time.Time `json:"created"`
}
