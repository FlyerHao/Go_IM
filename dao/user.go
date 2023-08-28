package dao

type Person struct {
    Base
    Name string `json:"name";gorm:"varchar(20)"`
    Phone string `json:"phone";gorm:""`
}
