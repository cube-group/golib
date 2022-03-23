package models

type Demo struct {
	Id   uint   `gorm:"column:id;primary_key"`
	Name string `gorm:""`

	No string `gorm:"-"` //不是表字段
}

func (t *Demo) TableName() string {
	return "demo"
}
