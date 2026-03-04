package orm

import "github.com/FourWD/middleware/model"

type LogBooking struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`
	model.GormModel

	BookingID       string `json:"booking_id" query:"booking_id" gorm:"type:varchar(36)"`
	BookingStatusID string `json:"booking_status_id" query:"booking_status_id" gorm:"type:varchar(2)"`
	Remark          string `json:"remark" query:"remark" gorm:"type:text"`
	LogChange       string `json:"log_change" query:"log_change" gorm:"type:text"`
}
