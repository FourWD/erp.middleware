package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type VehicleSubModel struct {
	midOrm.VehicleSubModel
	VehicleGradeID    string `json:"vehicle_grade_id" query:"vehicle_grade_id" gorm:"type:varchar(2)"`
	YearManufacturing string    `json:"year_manufacturing" query:"year_manufacturing" gorm:"type:varchar(4)"`
	OcpbName          string `json:"ocpb_name" query:"ocpb_name" gorm:"type:varchar(50)"`
}
