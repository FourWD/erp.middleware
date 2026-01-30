package orm

import midOrm "github.com/FourWD/middleware/orm"

type BookingCancel struct {
	midOrm.BookingCancel
	IsWeb bool   `json:"is_web" query:"is_web" gorm:"type:bool"`
}
