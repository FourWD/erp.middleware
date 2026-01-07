package orm

import "github.com/FourWD/middleware/model"

type Accessory struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`
	model.GormModel
	Name              string  `json:"name" query:"name" gorm:"type:varchar(256)"`
	AccessoryTypeID   string  `json:"accessory_type_id" query:"accessory_type_id" gorm:"type:varchar(4)"`
	PricePreVat       float64 `json:"price_pre_vat" query:"price_pre_vat" gorm:"type:decimal(14,2)"`
	Vat               float64 `json:"vat" query:"vat" gorm:"type:decimal(14,2)"`
	Price             float64 `json:"price" query:"price" gorm:"type:decimal(14,2)"`
	Label             string  `json:"label" query:"label" gorm:"type:varchar(100)"`
	ShopID            string  `json:"shop_id" query:"shop_id" gorm:"type:varchar(36)"`
	ImagePath         string  `json:"image_path" query:"image_path" gorm:"type:text"`
	IsEffectInsurance bool    `json:"is_effect_insurance" query:"is_effect_insurance" gorm:"type:bool"`
	IsTurn            bool    `json:"is_turn" query:"is_turn" gorm:"type:bool"`
	IsChassis         bool    `json:"is_chassis" query:"is_chassis" gorm:"type:bool"`
	model.GormRowOrder
}
