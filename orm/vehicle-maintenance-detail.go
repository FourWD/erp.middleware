package orm

import (
	"github.com/FourWD/middleware/model"
)

type VehicleMaintenanceDetail struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel
	VehicleMaintenanceID string  `json:"vehicle_maintenance_id" query:"vehicle_maintenance_id" gorm:"type:varchar(36)"`
	VehicleID            string  `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	ChassisNo            string  `json:"chassis_no" query:"chassis_no" gorm:"type:varchar(36)"`
	RepDetail            string  `json:"rep_detail" query:"rep_detail" gorm:"type:text"`
	RepairPrice          float64 `json:"repair_price" query:"repair_price" gorm:"type:decimal(14,2)"`
}
