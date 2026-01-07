package orm

import "github.com/FourWD/middleware/model"

// midOrm "github.com/FourWD/middleware/orm"

type CareerSub struct { //
	ID string `json:"id" query:"id" gorm:"type:varchar(2); primary_key"`
	model.GormModel

	CareerID string `json:"career_id" query:"career_id" gorm:"type:varchar(2)"`

	Name string `json:"name" query:"name" gorm:"type:varchar(256)"`
	model.GormRowOrder
}
