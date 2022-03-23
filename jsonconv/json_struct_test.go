package jsonconv

//将数据库对应的json字段定义成JSON结构，输出json的时候会自动解析
type Goods struct {
	ID         int  `gorm:"AUTO_INCREMENT;column:id;type:INT;primary_key" json:"id"`
	ClassForms JSON `gorm:"column:class_forms;type:JSON;" json:"class_forms"`
	Skus       JSON `gorm:"column:skus;type:JSON;" json:"skus"`
}
