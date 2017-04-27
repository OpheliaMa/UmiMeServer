package model

import "github.com/jinzhu/gorm"

type EntryForm struct {
	gorm.Model

	Name string
	Sex string
	School string
	Grade string
	ParentPhone string
	Remark string `gorm:"type:varchar(2000)"`
}
