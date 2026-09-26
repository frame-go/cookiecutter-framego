package models

import (
	"time"

	"github.com/frame-go/framego/uniqueid"
)

type Object struct {
	Id          uniqueid.ID `gorm:"primaryKey;autoIncrement=false;column:id" json:"id"`
	Name        string      `gorm:"column:name" json:"name"`
	Description string      `gorm:"column:description" json:"description"`
	Status      int16       `gorm:"column:status" json:"status"`
	CreateTime  time.Time   `gorm:"column:create_time" json:"create_time"`
	UpdateTime  time.Time   `gorm:"column:update_time" json:"update_time"`
}

func (*Object) TableName() string {
	return "object_tab"
}
