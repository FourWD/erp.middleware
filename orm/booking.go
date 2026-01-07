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
	UserTypeID       string     `json:"user_type_id" query:"user_type_id" gorm:"type:varchar(2)"`
	PrefixID         string     `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2)"`
	FirstName        string     `json:"first_name" query:"name" gorm:"type:varchar(256)"`
	LastName         string     `json:"last_name" query:"name" gorm:"type:varchar(256)"`
	Mobile           string     `json:"mobile" query:"mobile" gorm:"type:varchar(20)"`
	Email            string     `json:"email" query:"email" gorm:"type:varchar(50)"`
	GenderID         string     `json:"gender_id" query:"gender_id" gorm:"type:varchar(2)"`
	BirthDate        *time.Time `json:"birth_date" query:"birth_date" `
	IdCardNo         string     `json:"id_card_no" query:"id_card_no" gorm:"type:varchar(20)"`
	IsCustomerUseCar bool       `json:"is_customer_use_car" query:"is_customer_use_car" gorm:"type:bool"`

	// Job Description Section
	CareerID    string `json:"career_id" query:"career_id" gorm:"type:varchar(36)"`
	CareerSubID string `json:"career_sub_id" query:"career_sub_id" gorm:"type:varchar(36)"`

	CareerSalary int `json:"career_salary" query:"career_salary" gorm:"type:int"`
	// show user_type_id
	EvidenceID string `json:"evidence_id" query:"evidence_id" gorm:"type:varchar(36)"`

	ReferralChannelID string `json:"referral_channel_id" query:"referral_channel_id" gorm:"type:varchar(2)"`
	ReferralName      string `json:"referral_name" query:"referral_name" gorm:"type:varchar(256)"`

	// Address Section
	Address       string `json:"address" query:"address" gorm:"type:text"`
	DistrictID    string `json:"district_id" query:"district_id" gorm:"type:varchar(4)"`         //อำเภอ
	SubDistrictID string `json:"sub_district_id" query:"sub_district_id" gorm:"type:varchar(6)"` //ตำบล
	ProvinceID    string `json:"province_id" query:"province_id" gorm:"type:varchar(2)"`
	Postcode      string `json:"postcode" query:"postcode" gorm:"type:varchar(5)"`

	ContactAddress       string `json:"contact_address" query:"contact_address" gorm:"type:text"`
	ContactDistrictID    string `json:"contact_district_id" query:"contact_district_id" gorm:"type:varchar(4)"`         //อำเภอ
	ContactSubDistrictID string `json:"contact_sub_district_id" query:"contact_sub_district_id" gorm:"type:varchar(6)"` //ตำบล
	ContactProvinceID    string `json:"contact_province_id" query:"contact_province_id" gorm:"type:varchar(2)"`
	ContactPostcode      string `json:"contact_postcode" query:"contact_postcode" gorm:"type:varchar(5)"`

	InsuranceName        string  `json:"insurance_name" query:"insurance_name" gorm:"type:varchar(256)"`
	InsuranceType        string  `json:"insurance_type" query:"insurance_type" gorm:"type:varchar(256)"`
	InsuranceAmount      float64 `json:"insurance_amount" query:"insurance_amount" gorm:"type:decimal(14,2)"`
	InsuranceBasicAmount float64 `json:"insurance_basic_amount" query:"insurance_basic_amount" gorm:"type:decimal(14,2)"`

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

	// Loan Section
	PaymentMethodID   string  `json:"payment_method_id" query:"payment_method_id" gorm:"type:varchar(100)"`
	FinanceID         string  `json:"finance_id" query:"finance_id" gorm:"type:varchar(36)"`
	DownAmount        float64 `json:"down_amount" query:"down_amount" gorm:"type:decimal(14,2)"`
	LoanAmount        float64 `json:"loan_amount" query:"loan_amount" gorm:"type:decimal(14,2)"`
	Interest          float64 `json:"interest" query:"interest" gorm:"type:decimal(14,2)"`
	FinanceMonth      string  `json:"finance_month" query:"finance_month" gorm:"type:varchar(2)"`
	InstallmentAmount float64 `json:"installment_amount" query:"installment_amount" gorm:"type:decimal(14,2)"`

	// เพิ่ม 1. เซ้นใบสมัคร 2.ไฟแน้นมารับใบสมัคร 3.เอกสารครบ 4.อนุมัติ
	IsRegisterFinanceOnline bool       `json:"is_register_finance_online" query:"is_register_finance_online" gorm:"type:bool"`
	FinanceRegisterDate     *time.Time `json:"finance_register_date" query:"finance_register_date"`
	FinanceRegisterBy       string     `json:"finance_register_by" query:"finance_register_by" gorm:"type:varchar(100)"`
	FinanceReceiveDate      *time.Time `json:"finance_receive_date" query:"finance_receive_date"`
	FinanceReceiveBy        string     `json:"finance_receive_by" query:"finance_receive_by" gorm:"type:varchar(100)"`
	FinanceDocDoneDate      *time.Time `json:"finance_doc_done_date" query:"finance_doc_done_date"`
	FinanceDocDoneBy        string     `json:"finance_doc_done_by" query:"finance_doc_done_by" gorm:"type:varchar(100)"`
	FinanceApproveDate      *time.Time `json:"finance_approve_date" query:"finance_approve_date"`
	FinanceApproveBy        string     `json:"finance_approve_by" query:"finance_approve_by" gorm:"type:varchar(100)"`
	FinanceApproveNo        string     `json:"finance_approve_no" query:"finance_approve_no" gorm:"type:varchar(100)"`
	// FinanceAgentID      string     `json:"finance_agent_id" query:"finance_agent_id" gorm:"type:varchar(36)"`
	FinanceAgentName   string `json:"finance_agent_name" query:"finance_agent_name" gorm:"type:varchar(100)"`
	FinanceAgentMobile string `json:"finance_agent_mobile" query:"finance_agent_mobile" gorm:"type:varchar(20)"`

	// For Booking Report
	SentCaseBy           string     `json:"sent_case_by" query:"sent_case_by" gorm:"type:varchar(100)"`
	SentCaseDate         *time.Time `json:"sent_case_date" query:"sent_case_date"`
	SignContractBy       string     `json:"sign_contract_by" query:"sign_contract_by" gorm:"type:varchar(100)"`
	SignContractDate     *time.Time `json:"sign_contract_date" query:"sign_contract_date"`
	DocCompleteBy        string     `json:"doc_complete_by" query:"doc_complete_by" gorm:"type:varchar(100)"`
	DocCompleteDate      *time.Time `json:"doc_complete_date" query:"doc_complete_date"`
	FinanceRejectBy      string     `json:"finance_reject_by" query:"finance_reject_by" gorm:"type:varchar(100)"`
	FinanceRejectDate    *time.Time `json:"finance_reject_date" query:"finance_reject_date"`
	FinanceRejectReason  string     `json:"finance_reject_reason" query:"finance_reject_reason" gorm:"type:text"`
	FinanceRejectComment string     `json:"finance_reject_comment" query:"finance_reject_comment" gorm:"type:text"`

	// For SaleList Report
	DownInvoiceDate             *time.Time `json:"down_invoice_date" query:"down_invoice_date" `                                           // วันที่ออก invoice มัดจำ
	DownInvoiceNo               string     `json:"down_invoice_no" query:"down_invoice_no" gorm:"type:varchar(100)"`                       // เลขที่ invoice มัดจำ
	DownInvoiceCustomerName     string     `json:"down_invoice_customer_name" query:"down_invoice_customer_name" gorm:"type:varchar(100)"` // ชื่อลูกค้า invoice มัดจำ
	ReceiveDownDate             *time.Time `json:"receive_down_date" query:"receive_down_date"`                                            // วันที่รับมัดจำ
	ActualDownAmount            float64    `json:"actual_down_amount" query:"actual_down_amount" gorm:"type:decimal(12,4)"`                // จำนวนเงินมัดจำจริง
	VehicleInvoiceDate          *time.Time `json:"vehicle_invoice_date" query:"vehicle_invoice_date"`                                      // วันที่ออก invoice รถ
	VehicleInvoiceNo            string     `json:"vehicle_invoice_no" query:"vehicle_invoice_no"`                                          // เลขที่ invoice รถ
	VehicleInvoiceAmount        float64    `json:"vehicle_invoice_amount" query:"vehicle_invoice_amount" gorm:"type:decimal(12,4)"`        // จำนวนเงิน invoice รถ
	VehicleInvoiceCustomerName  string     `json:"vehicle_invoice_customer_name" query:"vehicle_invoice_customer_name" `                   // ชื่อลูกค้า invoice รถ
	OwnerName                   string     `json:"owner_name"`                                                                             // ชื่อเจ้าของ
	PossessorName               string     `json:"possessor_name"`                                                                         // ผู้ครอบครอง
	VehicleReceivePaymentDate   *time.Time `json:"vehicle_receive_payment_date"`                                                           // วันที่รับชำระเงินรถ
	VehicleReceivePaymentAmount float64    `json:"vehicle_receive_payment_amount"`                                                         // จำนวนเงินรับชำระรถ
	CommissionInvoiceNo         string     `json:"commission_invoice_no"`                                                                  // เลขที่ invoice ค่าคอม
	CommissionInvoiceDate       *time.Time `json:"commission_invoice_date"`                                                                // วันที่ออก invoice ค่าคอม
	CommissionReceiveDate       *time.Time `json:"commission_receive_date"`                                                                // วันที่รับค่าคอม
	CommissionAmount            float64    `json:"commission_amount"`

	// k'pom special discount
	Discount            float64    `json:"discount" query:"discount" gorm:"type:decimal(14,2)"`
	SpecialDiscount     float64    `json:"special_discount" query:"special_discount" gorm:"type:decimal(14,2)"`
	SpecialDiscountDate *time.Time `json:"special_discount_date" query:"special_discount_date"`
	SpecialDiscountBy   string     `json:"special_discount_by" query:"special_discount_by" gorm:"type:varchar(100)"`
	SpecialDiscountRef  string     `json:"special_discount_ref" query:"special_discount_ref" gorm:"type:varchar(1000)"`

	// CreditCard Section
	CreditCardNo string `json:"credit_card_no" query:"credit_card_no" gorm:"type:varchar(20)"`

	CompanyName string `json:"company_name" query:"company_name" gorm:"type:varchar(500)"`

	IsBookingPaid bool `json:"is_booking_paid" query:"is_booking_paid" gorm:"type:bool"` // จ่ายเงินดาว

}
