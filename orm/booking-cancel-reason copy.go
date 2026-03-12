package orm

import "github.com/FourWD/middleware/model"

// midOrm "github.com/FourWD/middleware/orm"

type BookingVehicleReason struct { //
	ID string `json:"id" query:"id" gorm:"type:varchar(3); primary_key"`
	model.GormModel

	BookingVehicleStatusID string `json:"booking_vehicle_status_id" query:"booking_vehicle_status_id" gorm:"type:varchar(2)"`

	Name string `json:"name" query:"name" gorm:"type:varchar(256)"`
	model.GormRowOrder
}
