package orm

import "time"

// midOrm "github.com/FourWD/middleware/orm"

type LogSyncVehicleImageError struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`

	CreatedAt time.Time `json:"created_at" query:"created_at" gorm:"<-:create"`
	CreatedBy string    `json:"created_by" query:"created_by"  gorm:"type:varchar(36)"`

	LogSyncVehicleImageID  string `json:"log_sync_vehicle_image_id" query:"log_sync_vehicle_image_id" gorm:"type:varchar(36)"`
	VehicleID              string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	TemplateVehicleImageID string `json:"template_vehicle_image_id" query:"template_vehicle_image_id" gorm:"type:varchar(36)"`

	Message string `json:"message" query:"message" gorm:"type:text;"`
}
