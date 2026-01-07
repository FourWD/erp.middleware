package orm

import "github.com/FourWD/middleware/model"

type BookingAccessory struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel
	BookingID   string  `json:"booking_id" query:"booking_id" gorm:"type:varchar(36);"`
	AccessoryID string  `json:"accessory_id" query:"accessory_id" gorm:"type:varchar(36);"`
	Price       float64 `json:"price" query:"price" gorm:"type:decimal(12,4)"`
}
