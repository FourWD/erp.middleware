package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type VehicleModel struct {
	midOrm.VehicleModel
	StartPrice float64 `json:"start_price" query:"start_price" gorm:"type:decimal(14,2)"`
	IsPickup   bool    `json:"is_pickup" query:"is_pickup" gorm:"type:bool"`
	ImagePath1 string  `json:"image_path_1" query:"image_path_1" gorm:"type:varchar(1024)"`
	ImagePath2 string  `json:"image_path_2" query:"image_path_2" gorm:"type:varchar(1024)"`
	ImagePath3 string  `json:"image_path_3" query:"image_path_3" gorm:"type:varchar(1024)"`
	ImagePath4 string  `json:"image_path_4" query:"image_path_4" gorm:"type:varchar(1024)"`
}
