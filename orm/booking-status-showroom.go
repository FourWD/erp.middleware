package orm

import "github.com/FourWD/middleware/model"

type BookingStatus struct { //
	ID string `json:"id" query:"id" gorm:"type:varchar(2); primary_key"`
	model.GormModel

	Name     string `json:"name" query:"name" gorm:"type:varchar(256)"`
	Priority string `json:"priority" query:"priority" gorm:"type:int"`
	model.GormRowOrder
}
