package request

type SendRequest struct {
    UserId int64 `json:"user_id"`
    SendUser int64 `json:"send_user"`
    SendGroup int64 `json:"send_group"`
    Data     string `json:"data"`
}
