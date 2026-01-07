package orm

import "time"

// midOrm "github.com/FourWD/middleware/orm"

type LogSyncVehicleError struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`

	CreatedAt time.Time `json:"created_at" query:"created_at" gorm:"<-:create"`
	CreatedBy string    `json:"created_by" query:"created_by"  gorm:"type:varchar(36)"`

	LogSyncVehicleID string `json:"log_sync_vehicle_id" query:"log_sync_vehicle_id" gorm:"type:varchar(36)"`
	ChassisNo        string `json:"chassis_no" query:"chassis_no" gorm:"type:varchar(36)"`
	EngineNo         string `json:"engine_no" query:"engine_no" gorm:"type:varchar(36)"`
	Message          string `json:"message" query:"message" gorm:"type:text;"`
}
