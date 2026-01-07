package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type GiftAwayPromotion struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`
	model.GormModel
	Name           string    `json:"name" query:"name" gorm:"type:varchar(256)"`
	VehicleModelID string    `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36)"`
	StartDate      time.Time `json:"start_date" query:"start_date"`
	EndDate        time.Time `json:"end_date" query:"end_date"`
	Description    string    `json:"description" query:"description" gorm:"type:varchar(500)"`
	IsActive       int       `json:"is_active" query:"is_active" gorm:"type:tinyint(1)"`
	model.GormRowOrder
}
