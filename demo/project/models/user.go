package models

import (
	"encoding/json"
	"time"
)

type User struct {
	Id       uint      `gorm:"column:id;primary_key" json:"id"`
	Username string    `gorm:"" json:"username"`
	CreateAt time.Time `gorm:"" json:"createAt"`
	UpdateAt time.Time `gorm:"" json:"updateAt"`
}

func (t *User) TableName() string {
	return "user"
}

func (t *User) ToBytes() []byte {
	b, _ := json.Marshal(t)
	return b
}
