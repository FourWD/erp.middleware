package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type LogVehicleExtUrl struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`

	model.GormModel

	VehicleID     string    `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	URL           string    `json:"url" query:"url" gorm:"type:varchar(500)"` // ex. www.usedcar.fourwd.me/product/Xds2g36s82
	UrlExpireDate time.Time `json:"url_expire_date" query:"url_expire_date" `
}
