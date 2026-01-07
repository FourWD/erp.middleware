package orm

import (
	"github.com/FourWD/middleware/model"
)

// midOrm "github.com/FourWD/middleware/orm"

type PromotionVehicle struct {
	ID          string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	PromotionID string `json:"promotion_id" query:"promotion_id" gorm:"type:varchar(36);"`
	VehicleID   string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36);"`
	model.GormModel
}
