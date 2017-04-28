package model

import "github.com/jinzhu/gorm"

type EntryForm struct {
	gorm.Model

	Name string `form:"name"`
	Sex string `form:"sex"`
	School string `form:"school"`
	Grade string `form:"grade"`
	ParentPhone string `form:"parentPhone"`
	Remark string `gorm:"type:varchar(2000)"`
}
