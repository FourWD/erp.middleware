package orm

import "github.com/FourWD/middleware/orm"

type BookingGift struct {
	orm.BookingGift
	GiftSetID string `json:"set_id" query:"set_id" gorm:"type:varchar(36)"`
}
