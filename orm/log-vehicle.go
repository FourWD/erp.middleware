package orm

// midOrm "github.com/FourWD/middleware/orm"

type LogVehicle struct {
	VehicleID string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	Vehicle
	// Action string `json:"action" query:"action" gorm:"type:varchar(20)"`
	// Detail string `json:"detail" query:"detail" gorm:"type:varchar(500)"`
}
