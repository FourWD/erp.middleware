package orm

import "github.com/FourWD/middleware/model"

type BookingAccessory struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key"`
	model.GormModel

	BookingID   string  `json:"booking_id" query:"booking_id" gorm:"type:varchar(36)"`
	AccessoryID string  `json:"accessory_id" query:"accessory_id" gorm:"type:varchar(36)"`
	UnitPreVat  float64 `json:"unit_price_pre_vat" query:"unit_price_pre_vat" gorm:"type:decimal(14,2)"`
	UnitVat     float64 `json:"unit_vat" query:"unit_vat" gorm:"type:decimal(14,2)"`
	UnitPrice   float64 `json:"unit_price" query:"unit_price" gorm:"type:decimal(14,2)"`
	Qty         int     `json:"qty" query:"qty" gorm:"type:int"`
	PricePreVat float64 `json:"price_pre_vat" query:"price_pre_vat" gorm:"type:decimal(14,2)"`
	Vat         float64 `json:"vat" query:"vat" gorm:"type:decimal(14,2)"`
	Price       float64 `json:"price" query:"price" gorm:"type:decimal(14,2)"`
	model.GormRowOrder
	//orm.BookingAccessory
	GroupID           string  `json:"group_id" query:"group_id" gorm:"type:varchar(4)"`
	ShopID            string  `json:"shop_id" query:"shop_id" gorm:"type:varchar(36)"`
	IsTrade           bool    `json:"is_trade" query:"is_trade" gorm:"type:bool"`
	IsFree            bool    `json:"is_free" query:"is_free" gorm:"type:bool"`
	TradePrice        float64 `json:"trade_price" query:"trade_price" gorm:"type:decimal(14,2)"`
	DiscountAccessory float64 `json:"discount_accessory" query:"discount_accessory" gorm:"type:decimal(14,2)"`
	IsEffectInsurance bool    `json:"is_effect_insurance" query:"is_effect_insurance" gorm:"type:bool"`
}
