package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type VehicleMaintenance struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel
	MaintenanceTypeID string `json:"maintenance_type_id" query:"maintenance_type_id" gorm:"type:varchar(36)"`
	VehicleID         string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	ChassisNo         string `json:"chassis_no" query:"chassis_no" gorm:"type:varchar(36)"`
	ServiceName       string `json:"service_name" query:"service_name" gorm:"type:varchar(500)"`
	ServiceDate       time.Time
	Kilo              string `json:"kilo" query:"kilo" gorm:"type:varchar(500)"`
}
