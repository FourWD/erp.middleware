package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type VehicleModel struct {
	midOrm.VehicleModel
	StartPrice float64  `json:"start_price" query:"start_price" gorm:"type:decimal(14,2)"`
	IsPickup   bool `json:"is_pickup" query:"is_pickup" gorm:"type:bool"`
}
