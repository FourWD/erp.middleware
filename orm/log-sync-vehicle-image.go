package orm

import "time"

// midOrm "github.com/FourWD/middleware/orm"

type LogSyncVehicleImage struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`

	CreatedAt time.Time `json:"created_at" query:"created_at"  gorm:"<-:create"`
	CreatedBy string    `json:"created_by" query:"created_by"  gorm:"type:varchar(36)"`

	Success int `json:"success" query:"success" gorm:"type:int"`
	Error   int `json:"error" query:"error" gorm:"type:int"`
	Skip    int `json:"skip" query:"skip" gorm:"type:int"`
	Total   int `json:"total" query:"total" gorm:"type:int"`
	// Message string `json:"message" query:"message" gorm:"type:text;"`
}
