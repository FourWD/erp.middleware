package orm

import (
	"time"
)

type LogNotification struct {
	ID        string    `json:"id" query:"id" gorm:"type:varchar(36)"`
	UpdatedAt time.Time `json:"updated_at" query:"updated_at"`
	CreatedBy string    `json:"created_by" query:"created_by" gorm:"type:varchar(36)"`
}
