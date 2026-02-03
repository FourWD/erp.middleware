package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type VehicleColors struct {
	midOrm.VehicleColor
	ColorCode string `json:"color_code" query:"color_code" gorm:"type:varchar(10)"`
	Code      string `json:"code" query:"code" gorm:"type:varchar(4)"`
	ImagePath string `json:"image_path" query:"image_path" gorm:"type:varchar(1024)"`
}
