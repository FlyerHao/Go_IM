package dao

type GroupUserConn struct {
    Base
    GroupId int64 `json:"group_id"`
    UserId  int64 `json:"user_id"`
}
