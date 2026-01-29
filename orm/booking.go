package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Booking struct {
	// midOrm.Booking // ไม่ต้องใส่ค่าลง vehicle_model_id, vehicle_sub_model_id ใช้จอยแทน
	ID string `json:"id" query:"id" gorm:"type:varchar(36);uniqueIndex:idx_id"`

	model.GormModel

	BookingNo            string     `json:"booking_no" query:"booking_no" gorm:"type:varchar(20)"` //หมายเลขรายการ
	BookingDate          *time.Time `json:"booking_date" query:"booking_date" `
	BookingBy            string     `json:"booking_by" query:"booking_by" gorm:"type:varchar(36)"`
	BookingStatusID      string     `json:"booking_status_id" query:"booking_status_id" gorm:"type:varchar(2)"`
	BookingCancelID      string     `json:"booking_cancel_id" query:"booking_cancel_id" gorm:"type:varchar(2)"`
	BookingCancelDate    *time.Time `json:"booking_cancel_date" query:"booking_cancel_date" `
	BookingCancelRemark  string     `json:"booking_cancel_remark" query:"booking_cancel_remark" gorm:"type:text"`
	BookbankCustomerName string     `json:"bookbank_customer_name" query:"bookbank_customer_name" gorm:"type:varchar(256)"`

	SaleEmployeeID string `json:"sale_employee_id" query:"sale_employee_id" gorm:"type:varchar(36)"` // sale_contact_id
	SaleBranchID   string `json:"sale_branch_id" query:"sale_branch_id" gorm:"type:varchar(36)"`

	UserID           string     `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	IsCustomerUseCar bool       `json:"is_customer_use_car" query:"is_customer_use_car" gorm:"type:bool"`

	// show user_type_id
	EvidenceID string `json:"evidence_id" query:"evidence_id" gorm:"type:varchar(36)"`
	ReferralChannelID string `json:"referral_channel_id" query:"referral_channel_id" gorm:"type:varchar(2)"`
	ReferralName      string `json:"referral_name" query:"referral_name" gorm:"type:varchar(256)"`
	ReserveDate        *time.Time `json:"reserve_date" query:"reserve_date"`
	ReserveBy          string     `json:"reserve_by" query:"reserve_by" gorm:"type:varchar(36)"`
	ReserveDeposit     float64    `json:"reserve_deposit" query:"reserve_deposit" gorm:"type:decimal(14,2)"`
	CaseSubmissionDate *time.Time `json:"case_submission_date" query:"case_submission_date"`
	ExpectDeliveryDate *time.Time `json:"expect_delivery_date" query:"expect_delivery_date"`

	// RunningNo       int `json:"running_no" query:"running_no" gorm:"autoIncrement"`
	PaymentTypeID string `json:"payment_type_id" query:"payment_type_id" gorm:"type:varchar(2)"` // ประเภทการชําระ การมัดจำ
	VehicleID string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	IsPaid    bool   `json:"is_paid" query:"is_paid" gorm:"type:bool"`
	IsCancel  bool   `json:"is_cancel" query:"is_cancel" gorm:"type:bool"`
	Remark    string `json:"remark" query:"remark" gorm:"type:text"`
	PricePreVat float64 `json:"price_pre_vat" query:"price_pre_vat" gorm:"type:decimal(14,2)"`
	Vat         float64 `json:"vat" query:"vat" gorm:"type:decimal(14,2)"`
	Price       float64 `json:"price" query:"price" gorm:"type:decimal(14,2)"`

	// Website Section Required
	IdCardPath      string `json:"id_card_path" query:"id_card_path" gorm:"type:text"`
	IsAcceptConsent bool   `json:"is_accept_consent" query:"is_accept_consent" gorm:"type:bool"`
	ConsentID       string `json:"consent_id" query:"consent_id" gorm:"type:varchar(36)"`
	RunningNo int `json:"running_no" query:"running_no" gorm:"primary_key;auto_increment;not_null"`

	// CreditCard Section
	CreditCardNo string `json:"credit_card_no" query:"credit_card_no" gorm:"type:varchar(20)"`
	CompanyName string `json:"company_name" query:"company_name" gorm:"type:varchar(500)"`
	IsBookingPaid bool `json:"is_booking_paid" query:"is_booking_paid" gorm:"type:bool"` // จ่ายเงินดาว

}
