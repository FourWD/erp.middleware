package orm

import "time"

type BookingDetail struct {
	
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

	
	}
