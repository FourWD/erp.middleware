package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Notification struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`
	model.GormModel
	BookingID        string    `json:"booking_id" query:"booking_id" gorm:"type:varchar(36)"`
	BookingStatusID  string    `json:"booking_status_id" query:"booking_status_id" gorm:"type:varchar(2)"`
	BranchID         string    `json:"branch_id" query:"branch_id" gorm:"type:varchar(36)"`
	NotificationDate time.Time `json:"notification_date" query:"notification_date"`
}
