package orm

import (
	"github.com/FourWD/middleware/model"
)

// midOrm "github.com/FourWD/middleware/orm"

type PoDetail struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`

	model.GormModel

	PoID      string `json:"po_no" query:"po_no" gorm:"type:varchar(36)"`           //หมายเลขรายการ
	VehicleID string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"` //หมายเลขรายการ

	PricePreVat float64 `json:"price_pre_vat" query:"price_pre_vat" gorm:"type:decimal(14,2)"`
	Vat         float64 `json:"vat" query:"vat" gorm:"type:decimal(14,2)"`
	Price       float64 `json:"total_price" query:"total_price" gorm:"type:decimal(14,2)"`
	NetPrice    float64 `json:"net_price" query:"net_price" gorm:"type:decimal(14,2)"`

	// PoStatusID       string `json:"po_status_id" query:"po_status_id" gorm:"type:varchar(2)"`

	// AuctionID     string    `json:"auction_id" query:"auction_id" gorm:"type:varchar(36)"`
	// UserID        string    `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	// PaymentTypeID string    `json:"payment_type_id" query:"payment_type_id" gorm:"type:varchar(2)"`
	// PaymentDate   time.Time `json:"payment_date" query:"payment_date"`
	// Amount        float64   `json:"amount" query:"amount" gorm:"type:decimal(14,2)"`
	// ImageUrl      string    `json:"image_url" query:"image_url" gorm:"type:varchar(1000)"`
	// Remark        string    `json:"remark" query:"remark" gorm:"type:varchar(50)"`
}
