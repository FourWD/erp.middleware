package orm

import "github.com/FourWD/middleware/model"

// midOrm "github.com/FourWD/middleware/orm"

type BookingCancelReason struct { //
	ID string `json:"id" query:"id" gorm:"type:varchar(3); primary_key"`
	model.GormModel

	BookingStatusID string `json:"booking_status_id" query:"booking_status_id" gorm:"type:varchar(2)"`

	Name string `json:"name" query:"name" gorm:"type:varchar(256)"`
	model.GormRowOrder
}
