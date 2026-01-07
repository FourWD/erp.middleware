package orm

import "time"

type LogSyncCsv struct {
	ID        string    `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	CreatedAt time.Time `json:"created_at" query:"created_at"  gorm:"<-:create"`
	CreatedBy string    `json:"created_by" query:"created_by"  gorm:"type:varchar(36)"`
	FilePath  string    `json:"file_path" query:"file_path" gorm:"type:text"`
	SyncName  string    `json:"sync_name" query:"sync_name" gorm:"type:varchar(500)"` // ชื่อกรุปข้อมูลว่า sync อะไร
}
