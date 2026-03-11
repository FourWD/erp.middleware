package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type BookingDetail struct {
	BookingID string `json:"booking_id" query:"booking_id" gorm:"type:varchar(36);primary_key"`

	model.GormModel
	EvidenceID string `json:"evidence_id" query:"evidence_id" gorm:"type:varchar(36)"` // หลักฐาน

	InsuranceID          string  `json:"insurance_id" query:"insurance_id" gorm:"type:varchar(10)"`
	InsuranceName        string  `json:"insurance_name" query:"insurance_name" gorm:"type:varchar(256)"`
	InsuranceType        string  `json:"insurance_type" query:"insurance_type" gorm:"type:varchar(256)"`
	InsuranceAmount      float64 `json:"insurance_amount" query:"insurance_amount" gorm:"type:decimal(14,2)"`
	InsuranceBasicAmount float64 `json:"insurance_basic_amount" query:"insurance_basic_amount" gorm:"type:decimal(14,2)"`
	InsuranceFund        float64 `json:"insurance_fund" query:"insurance_fund" gorm:"type:decimal(14,2)"` //

	// For SaleList Report
	DownInvoiceDate         *time.Time `json:"down_invoice_date" query:"down_invoice_date" `                                           // วันที่ออก invoice มัดจำ
	DownInvoiceNo           string     `json:"down_invoice_no" query:"down_invoice_no" gorm:"type:varchar(100)"`                       // เลขที่ invoice มัดจำ
	DownInvoiceCustomerName string     `json:"down_invoice_customer_name" query:"down_invoice_customer_name" gorm:"type:varchar(100)"` // ชื่อลูกค้า invoice มัดจำ
	ReceiveDownDate         *time.Time `json:"receive_down_date" query:"receive_down_date"`                                            // วันที่รับมัดจำ
	ActualDownAmount        float64    `json:"actual_down_amount" query:"actual_down_amount" gorm:"type:decimal(12,4)"`                // จำนวนเงินมัดจำจริง
	InvoiceDate             *time.Time `json:"invoice_date" query:"invoice_date"`                                                      // วันที่ออก invoice รถ
	InvoiceNo               string     `json:"invoice_no" query:"invoice_no"`                                                          // เลขที่ invoice รถ
	InvoiceAmount           float64    `json:"invoice_amount" query:"invoice_amount" gorm:"type:decimal(12,4)"`                        // จำนวนเงิน invoice รถ
	InvoiceCustomerName     string     `json:"invoice_customer_name" query:"invoice_customer_name" `                                   // ชื่อลูกค้า invoice รถ
	ReceivePaymentDate      *time.Time `json:"receive_payment_date"`                                                                   // วันที่รับชำระเงินรถ                                                        // วันที่รับชำระเงินรถ
	ReceivePaymentAmount    float64    `json:"receive_payment_amount"`                                                                 // จำนวนเงินรับชำระรถ
	CommissionInvoiceNo     string     `json:"commission_invoice_no"`                                                                  // เลขที่ invoice ค่าคอม
	CommissionInvoiceDate   *time.Time `json:"commission_invoice_date"`                                                                // วันที่ออก invoice ค่าคอม
	CommissionReceiveDate   *time.Time `json:"commission_receive_date"`                                                                // วันที่รับค่าคอม
	CommissionAmount        float64    `json:"commission_amount"`

	OwnerName     string `json:"owner_name"`     // เจ้าของรถ                                                                        // ชื่อเจ้าของ
	PossessorName string `json:"possessor_name"` // ผู้ครอบครอง เช่น เจ้าของคืออั้มผู้ครอบครองคือ WIN leasing

	// For Booking ManagerSign ของ Flook
	CustomerSignature               string     `json:"customer_signature" query:"customer_signature" gorm:"type:text"`
	CustomerSourceID                string     `json:"customer_source_id" query:"customer_source_id" gorm:"type:varchar(36)"`
	CustomerSignatureDate           *time.Time `json:"customer_signature_date" query:"customer_signature_date" gorm:"type:datetime"`
	CustomerSignatureLink           string     `json:"customer_signature_link" query:"customer_signature_link" gorm:"type:text"`
	CustomerSignatureLinkSendDate   *time.Time `json:"customer_signature_link_send_date" query:"customer_signature_link_send_date" gorm:"type:datetime"`
	CustomerSignatureLinkExpireDate *time.Time `json:"customer_signature_link_expire_date" query:"customer_signature_link_expire_date" gorm:"type:datetime"`
	SaleManagerSignatureEmployeeID  string     `json:"sale_manager_signature_employee_id" query:"sale_manager_signature_employee_id" gorm:"type:varchar(36)"`
	SaleManagerSignatureDate        *time.Time `json:"sale_manager_signature_date" query:"sale_manager_signature_date" gorm:"type:datetime"`
	Witness1Signature               string     `json:"witness_1_signature" query:"witness_1_signature" gorm:"type:text"`
	Witness1SignatureDate           *time.Time `json:"witness_1_signature_date" query:"witness_1_signature_date" gorm:"type:datetime"`
	CompanySignatureEmployeeID      string     `json:"company_signature_employee_id" query:"company_signature_employee_id" gorm:"type:varchar(36)"`
	CompanySignatureDate            *time.Time `json:"company_signature_date" query:"company_signature_date" gorm:"type:datetime"`
	SaleSignatureEmployeeID         string     `json:"sale_signature_employee_id" query:"sale_signature_employee_id" gorm:"type:varchar(36)"`
	SaleSignatureDate               *time.Time `json:"sale_signature_date" query:"sale_signature_date" gorm:"type:datetime"`
	Witness2Signature               string     `json:"witness_2_signature" query:"witness_2_signature" gorm:"type:text"`
	Witness2SignatureDate           *time.Time `json:"witness_2_signature_date" query:"witness_2_signature_date" gorm:"type:datetime"`

	IsManagerApprove bool `json:"is_manager_approve" query:"is_manager_approve"` //ผู้จัดการอนุมัติ
	IsCompanyApprove bool `json:"is_company_approve" query:"is_company_approve"` //ผู้ประกอบธุรกิจอนุมัติ

	ManagerApproveDate *time.Time `json:"manager_approve_date" query:"manager_approve_date" gorm:"type:datetime"`

	CustomerTypeID string `json:"customer_type_id" query:"customer_type_id" gorm:"type:varchar(2)"`
	FirstName      string `json:"first_name" query:"name" gorm:"type:varchar(256)"`
	LastName       string `json:"last_name" query:"name" gorm:"type:varchar(256)"`
	Mobile         string `json:"mobile" query:"mobile" gorm:"type:varchar(10)"`
	Email          string `json:"email" query:"email" gorm:"type:varchar(50)"`
	NationalityID  string `json:"nationality_id" query:"nationality_id" gorm:"type:varchar(2)"`
	ChannelID      string `json:"channel_id" query:"channel_id" gorm:"type:varchar(2)"`
	SubChannelID   string `json:"sub_channel_id" query:"sub_channel_id" gorm:"type:varchar(2)"`
	Facebook       string `json:"facebook" query:"facebook" gorm:"type:varchar(50)"`
	Line           string `json:"line" query:"line" gorm:"type:varchar(20)"`

	//มาจาก booking
	PrefixID  string     `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2)"`
	GenderID  string     `json:"gender_id" query:"gender_id" gorm:"type:varchar(2)"`
	BirthDate *time.Time `json:"birth_date" query:"birth_date"`
	IdCardNo  string     `json:"id_card_no" query:"id_card_no" gorm:"type:varchar(20)"`

	// Address Section
	Address       string `json:"address" query:"address" gorm:"type:text"`
	HouseNo       string `json:"house_no" query:"house_no" gorm:"type:varchar(20)"`
	Floor         string `json:"floor" query:"floor" gorm:"type:varchar(5)"`
	Room          string `json:"room" query:"room" gorm:"type:varchar(20)"`
	Moo           string `json:"moo" query:"moo" gorm:"type:varchar(5)"`
	Soi           string `json:"soi" query:"soi" gorm:"type:varchar(200)"`
	Road          string `json:"road" query:"road" gorm:"type:varchar(200)"`
	DistrictID    string `json:"district_id" query:"district_id" gorm:"type:varchar(4)"`         //อำเภอ
	SubDistrictID string `json:"sub_district_id" query:"sub_district_id" gorm:"type:varchar(6)"` //ตำบล
	ProvinceID    string `json:"province_id" query:"province_id" gorm:"type:varchar(2)"`
	Postcode      string `json:"postcode" query:"postcode" gorm:"type:varchar(5)"`

	// Contact or Current Address
	ContactHouseNo       string `json:"contact_house_no" query:"contact_house_no" gorm:"type:varchar(20)"`
	ContactFloor         string `json:"contact_floor" query:"contact_floor" gorm:"type:varchar(5)"`
	ContactRoom          string `json:"contact_room" query:"contact_room" gorm:"type:varchar(20)"`
	ContactMoo           string `json:"contact_moo" query:"contact_moo" gorm:"type:varchar(5)"`
	ContactSoi           string `json:"contact_soi" query:"contact_soi" gorm:"type:varchar(200)"`
	ContactRoad          string `json:"contact_road" query:"contact_road" gorm:"type:varchar(200)"`
	ContactAddress       string `json:"contact_address" query:"contact_address" gorm:"type:text"`
	ContactDistrictID    string `json:"contact_district_id" query:"contact_district_id" gorm:"type:varchar(4)"`         //อำเภอ
	ContactSubDistrictID string `json:"contact_sub_district_id" query:"contact_sub_district_id" gorm:"type:varchar(6)"` //ตำบล
	ContactProvinceID    string `json:"contact_province_id" query:"contact_province_id" gorm:"type:varchar(2)"`
	ContactPostcode      string `json:"contact_postcode" query:"contact_postcode" gorm:"type:varchar(5)"`

	// Job Description Section
	CareerID     string `json:"career_id" query:"career_id" gorm:"type:varchar(36)"`
	CareerSubID  string `json:"career_sub_id" query:"career_sub_id" gorm:"type:varchar(36)"`
	CareerSalary int    `json:"career_salary" query:"career_salary" gorm:"type:int"`

	// Company Section
	CompanyNo           string     `json:"company_no" query:"company_no" gorm:"type:varchar(20)"`                      //เลขนิติบุคคล
	CompanyRegisterDate *time.Time `json:"company_register_date" query:"company_register_date" gorm:"type:datetime"`   // วันที่จดทะเบียนบริษัท
	CompanyRegisterType string     `json:"company_register_type" query:"company_register_type" gorm:"type:varchar(2)"` //ประเภทบริษัท
	CompanyName         string     `json:"company_name" query:"company_name" gorm:"type:varchar(500)"`                 // ชื่อบริษัท
	CompanyBranchCode   string     `json:"company_branch_code" query:"company_branch_code" gorm:"type:varchar(10)"`    // รหัสสาขา
	CompanyBranchName   string     `json:"company_branch_name" query:"company_branch_name" gorm:"type:varchar(500)"`   // ชื่อสาขา
	CompanyBusinessType string     `json:"company_business_type" query:"company_business_type" gorm:"type:varchar(2)"` // ประเภทประกอบธุรกิจ

	// Contact Person Section
	ContactPersonPrefixID     string `json:"contact_person_prefix_id" query:"contact_person_prefix_id" gorm:"type:varchar(2)"`
	ContactPersonFirstName    string `json:"contact_person_first_name" query:"contact_person_first_name" gorm:"type:varchar(500)"`
	ContactPersonLastName     string `json:"contact_person_last_name" query:"contact_person_last_name" gorm:"type:varchar(500)"`
	ContactPersonMobile       string `json:"contact_person_mobile" query:"contact_person_mobile" gorm:"type:varchar(20)"`
	ContactPersonCompanyPhone string `json:"contact_person_company_phone" query:"contact_person_company_phone" gorm:"type:varchar(20)"`
	ContactPersonNickname     string `json:"contact_person_nickname" query:"contact_person_nickname" gorm:"type:varchar(500)"`
	ContactPersonPosition     string `json:"contact_person_position" query:"contact_person_position" gorm:"type:varchar(500)"`
	ContactPersonEmail        string `json:"contact_person_email" query:"contact_person_email" gorm:"type:varchar(500)"`
	ContactPersonLineID       string `json:"contact_person_line_id" query:"contact_person_line_id" gorm:"type:varchar(500)"`
	ContactPersonFacebook     string `json:"contact_person_facebook" query:"contact_person_facebook" gorm:"type:varchar(500)"`

	IdCardPath      string `json:"id_card_path" query:"id_card_path" gorm:"type:text"`
	IsAcceptConsent bool   `json:"is_accept_consent" query:"is_accept_consent" gorm:"type:bool"`
	ConsentID       string `json:"consent_id" query:"consent_id" gorm:"type:varchar(36)"`
}
