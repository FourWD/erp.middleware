package orm

import "time"

type LogSyncCsvDetail struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`

	CreatedAt time.Time `json:"created_at" query:"created_at"  gorm:"<-:create"`

	SyncCsvID   string `json:"sync_csv_id" query:"sync_csv_id" gorm:"type:varchar(36)"`
	Description string `json:"description" query:"description" gorm:"type:text"` // ข้อมูลลายหลัก
}
