package dao

type Friend struct {
    Base
    UserId1 int64 `json:"user_id_1"`
    UserId2 int64 `json:"user_id_2"`
    Status  int   `json:"status"`
}
