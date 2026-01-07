package orm

import (
	"github.com/FourWD/middleware/model"
)

type GiftAwayPromotionDetail struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`
	model.GormModel
	GiftPromotionID string `json:"gift_promotion_id" query:"gift_promotion_id" gorm:"type:varchar(36)"`
	GiftID          string `json:"gift_id" query:"gift_id" gorm:"type:varchar(36)"`
	IsDefault       int    `json:"is_default" query:"is_default" gorm:"type:tinyint(1)"`
	model.GormRowOrder
}
