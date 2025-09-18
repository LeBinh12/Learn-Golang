package common

import "time"

type SQLModel struct {
	Id int `json:"id" gorm:"column:id;"`

	CreateAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdateAt *time.Time `json:"update_at,omitempty" gorm:"column:update_at"`
}
