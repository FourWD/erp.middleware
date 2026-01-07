package orm

import (
	"github.com/FourWD/middleware/model"
)

type BookingOption struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`

	model.GormModel

	BookingID    string `json:"booking_id" query:"booking_id" gorm:"type:varchar(36)"`
	OptionGiftID string `json:"option_gift_id" query:"option_gift_id" gorm:"type:varchar(36)"`
	Cost         int    `json:"cost" query:"cost" gorm:"type:int"`
}
