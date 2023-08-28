package dao

type Group struct {
    Base
    Name string `json:"name";gorm:""`
}
