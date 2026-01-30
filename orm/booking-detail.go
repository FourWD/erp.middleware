package orm

import (
	"time"
)

type BookingDetail struct {
	BookingID string `json:"booking_id" query:"booking_id" gorm:"type:varchar(36);primary_key"`

	ContactAddress       string `json:"contact_address" query:"contact_address" gorm:"type:text"`
	ContactDistrictID    string `json:"contact_district_id" query:"contact_district_id" gorm:"type:varchar(4)"`         //อำเภอ
	ContactSubDistrictID string `json:"contact_sub_district_id" query:"contact_sub_district_id" gorm:"type:varchar(6)"` //ตำบล
	ContactProvinceID    string `json:"contact_province_id" query:"contact_province_id" gorm:"type:varchar(2)"`
	ContactPostcode      string `json:"contact_postcode" query:"contact_postcode" gorm:"type:varchar(5)"`

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
	CustomerSignature               string    `json:"customer_signature" query:"customer_signature" gorm:"type:text"`
	CustomerSourceID                string    `json:"customer_source_id" query:"customer_source_id" gorm:"type:varchar(36)"`
	CustomerSignatureDate           time.Time `json:"customer_signature_date" query:"customer_signature_date" gorm:"type:datetime"`
	CustomerSignatureLink           string    `json:"customer_signature_link" query:"customer_signature_link" gorm:"type:text"`
	CustomerSignatureLinkSendDate   time.Time `json:"customer_signature_link_send_date" query:"customer_signature_link_send_date" gorm:"type:datetime"`
	CustomerSignatureLinkExpireDate time.Time `json:"customer_signature_link_expire_date" query:"customer_signature_link_expire_date" gorm:"type:datetime"`
	SaleManagerSignatureEmployeeID  string    `json:"sale_manager_signature_employee_id" query:"sale_manager_signature_employee_id" gorm:"type:varchar(36)"`
	SaleManagerSignatureDate        time.Time `json:"sale_manager_signature_date" query:"sale_manager_signature_date" gorm:"type:datetime"`
	Witness1Signature               string    `json:"witness_1_signature" query:"witness_1_signature" gorm:"type:text"`
	Witness1SignatureDate           time.Time `json:"witness_1_signature_date" query:"witness_1_signature_date" gorm:"type:datetime"`
	CompanySignatureEmployeeID      string    `json:"company_signature_employee_id" query:"company_signature_employee_id" gorm:"type:varchar(36)"`
	CompanySignatureDate            time.Time `json:"company_signature_date" query:"company_signature_date" gorm:"type:datetime"`
	SaleSignatureEmployeeID         string    `json:"sale_signature_employee_id" query:"sale_signature_employee_id" gorm:"type:varchar(36)"`
	SaleSignatureDate               time.Time `json:"sale_signature_date" query:"sale_signature_date" gorm:"type:datetime"`
	Witness2Signature               string    `json:"witness_2_signature" query:"witness_2_signature" gorm:"type:text"`
	Witness2SignatureDate           time.Time `json:"witness_2_signature_date" query:"witness_2_signature_date" gorm:"type:datetime"`

	IsManagerApprove   bool       `json:"is_manager_approve" query:"is_manager_approve"`
	ManagerApproveDate *time.Time `json:"manager_approve_date" query:"manager_approve_date" gorm:"type:datetime"`
}
