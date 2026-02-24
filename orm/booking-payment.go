package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type BookingPayment struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key"`
	model.GormModel

	BookingID string `json:"booking_id" query:"booking_id" gorm:"type:varchar(36)"`

	PaymentID       string     `json:"payment_id" query:"payment_id" gorm:"type:varchar(36)"`              // 00 เงินจอง 01 เงินดาว 02 จ่ายค่ารถ
	PaymentTypeID   string     `json:"payment_type_id" query:"payment_type_id" gorm:"type:varchar(2)"`     // ประเภทการชําระ การมัดจำ
	PaymentMethodID string     `json:"payment_method_id" query:"payment_method_id" gorm:"type:varchar(2)"` // วิธีชำระค่ารถ
	Amont           float64    `json:"amont" query:"amont" gorm:"type:decimal(14,2)"`
	TransectionDate *time.Time `json:"transection_date" query:"transection_date" gorm:"type:varchar(36)"`

	BankID               string `json:"bank_id" query:"bank_id" gorm:"type:varchar(2)"`
	BookbankAccountNo    string `json:"bookbank_account_no" query:"bookbank_account_no" gorm:"type:varchar(50)"`
	BookbankAccountName  string `json:"bookbank_account_name" query:"bookbank_account_name" gorm:"type:varchar(256)"`
	ChequeNo             string `json:"cheque_no" query:"cheque_no" gorm:"type:varchar(50)"`
	ChequeName           string `json:"cheque_name" query:"cheque_name" gorm:"type:varchar(256)"`
	CreditCardNo         string `json:"credit_card_no" query:"credit_card_no" gorm:"type:varchar(36)"` //mark-up
	CreditCardExpireDate string `json:"credit_card_expire_date" query:"credit_card_expire_date" gorm:"type:varchar(36)"`

	FileSlip  string `json:"file_slip"   query:"file_slip"   gorm:"type:varchar(500)"`                    // เงินจอง
	FileSlip1 string `json:"file_slip_1" query:"file_slip_1" gorm:"column:file_slip_1;type:varchar(500)"` //เงินค่ารถ 1
	FileSlip2 string `json:"file_slip_2" query:"file_slip_2" gorm:"column:file_slip_2;type:varchar(500)"` //เงินค่ารถ 2
	FileSlip3 string `json:"file_slip_3" query:"file_slip_3" gorm:"column:file_slip_3;type:varchar(500)"` //เงินค่ารถ 3
	FileSlip4 string `json:"file_slip_4" query:"file_slip_4" gorm:"column:file_slip_4;type:varchar(500)"` //เงินค่ารถ 4
	FileSlip5 string `json:"file_slip_5" query:"file_slip_5" gorm:"column:file_slip_5;type:varchar(500)"` //เงินค่ารถ 5

	IsReservePaid     bool   `json:"is_reserve_paid" query:"is_reserve_paid" gorm:"type:bool"`
	NoReserveRemarkID string `json:"no_reserve_remark_id" query:"no_reserve_remark_id" gorm:"type:varchar(500)"`

	// DownPrice     float64   `json:"down_price" query:"down_price" gorm:"type:decimal(14,2)"`
	// DownPercent   float64   `json:"down_precent" query:"down_precent" gorm:"type:decimal(14,2)"`
	// DepositPrice         float64   `json:"deposit_price" query:"deposit_pricet" gorm:"type:decimal(14,2)"`

	// PoNo string `json:"po_no" query:"po_no" gorm:"type:varchar(30)"`
	// EvidenceID             string    `json:"evidence_id" query:"evidence_id" gorm:"type:varchar(36)"` // หลักฐาน

	model.GormRowOrder
}
