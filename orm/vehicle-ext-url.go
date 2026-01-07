package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type VehicleExtUrl struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`

	model.GormModel

	VehicleID   string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	GenByUserID string    `json:"gen_by_user_id" query:"gen_by_user_id" gorm:"type:varchar(36)"`
	URL         string    `json:"url" query:"url" gorm:"type:varchar(500)"` // ex. www.usedcar.fourwd.me/product/Xds2g36s82
	RandomPath  string    `json:"random_path" query:"random_path" gorm:"type:varchar(100)"`
	ExpireDate  time.Time `json:"expire_date" query:"expire_date" ` // 1 Day
	ClickCount  int       `json:"click_count" query:"click_count" gorm:"type:int"`
	SharedCount int       `json:"shared_count" query:"shared_count" gorm:"type:int"`
}
