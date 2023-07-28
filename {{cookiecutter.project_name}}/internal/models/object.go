package models

type Object struct {
	ID         uint64 `gorm:"primaryKey;column:id"`
	Name       string `gorm:"column:name"`
	Data       string `gorm:"column:data"`
	CreateTime uint32 `gorm:"column:create_time"`
	UpdateTime uint32 `gorm:"column:update_time"`
}

func (*Object) TableName() string {
	return "object_tab"
}
