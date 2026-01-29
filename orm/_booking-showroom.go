package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type BookingShowroom struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);uniqueIndex:idx_id"`
	model.GormModel
	BookingNo       string `json:"booking_no" query:"booking_no" gorm:"type:varchar(20)"` //หมายเลขรายการ
	BookingStatusID string `json:"booking_status_id" query:"booking_status_id" gorm:"type:varchar(2)"`
	BookingCancelID string `json:"booking_cancel_id" query:"booking_cancel_id" gorm:"type:varchar(2)"`

	UserID    string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
	PrefixID  string `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2)"`
	FirstName string `json:"first_name" query:"name" gorm:"type:varchar(256)"`
	LastName  string `json:"last_name" query:"name" gorm:"type:varchar(256)"`
	Mobile    string `json:"mobile" query:"mobile" gorm:"type:varchar(20)"`
	Email     string `json:"email" query:"email" gorm:"type:varchar(50)"`

	VehicleID           string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	VehicleModelID      string `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36)"`
	VehicleSubModelID   string `json:"vehicle_sub_model_id" query:"vehicle_sub_model_id" gorm:"type:varchar(36)"`
	VehicleColorID      string `json:"vehicle_color_id" query:"vehicle_color_id" gorm:"type:varchar(36)"`
	IsPaid              bool   `json:"is_paid" query:"is_paid" gorm:"type:bool"`
	IsCancel            bool   `json:"is_cancel" query:"is_cancel" gorm:"type:bool"`
	Remark              string `json:"remark" query:"remark" gorm:"type:text"`
	BookingCancelRemark string `json:"booking_cancel_remark" query:"booking_cancel_remark" gorm:"type:text"`
	EmployeeID          string `json:"employee_id" query:"employee_id" gorm:"type:varchar(36)"`

	Address       string `json:"address" query:"address" gorm:"type:text"`
	Building      string `json:"building" query:"building" gorm:"type:varchar(100)"`
	Room          string `json:"room" query:"room" gorm:"type:varchar(20)"`
	Street        string `json:"street" query:"street" gorm:"type:varchar(200)"`
	DistrictID    string `json:"district_id" query:"district_id" gorm:"type:varchar(4)"`         //อำเภอ
	SubDistrictID string `json:"sub_district_id" query:"sub_district_id" gorm:"type:varchar(6)"` //ตำบล
	ProvinceID    string `json:"province_id" query:"province_id" gorm:"type:varchar(2)"`
	Postcode      string `json:"postcode" query:"postcode" gorm:"type:varchar(5)"`

	Facebook string `json:"facebook" query:"facebook" gorm:"type:varchar(50)"`
	Line     string `json:"line" query:"line" gorm:"type:varchar(20)"`
	Tiktok   string `json:"tiktok" query:"tiktok" gorm:"type:varchar(50)"`

	PricePreVat float64 `json:"price_pre_vat" query:"price_pre_vat" gorm:"type:decimal(14,2)"`
	Vat         float64 `json:"vat" query:"vat" gorm:"type:decimal(14,2)"`
	Price       float64 `json:"price" query:"price" gorm:"type:decimal(14,2)"`

	RunningNo int `json:"running_no" query:"running_no" gorm:"primary_key;auto_increment;not_null"`

	VehicleYear           string    `json:"vehicle_year" query:"vehicle_year" gorm:"type:varchar(4)"`
	VehicleImg            string    `json:"vehicle_img" query:"vehicle_img" gorm:"type:text"`
	IDCard                string    `json:"id_card" query:"id_card" gorm:"type:varchar(13)"`
	CompanyRegisterNo     string    `json:"company_register_no" query:"company_register_no" gorm:"type:varchar(13)"`
	BirthDate             time.Time `json:"birth_date" query:"birth_date" gorm:"type:date"`
	OccupationMainID      string    `json:"occupation_main_id" query:"occupation_main_id" gorm:"type:varchar(255)"`
	OccupationID          string    `json:"occupation_id" query:"occupation_id" gorm:"type:varchar(255)"`
	Telephone             string    `json:"telephone" query:"telephone" gorm:"type:varchar(30)"`
	CustomerSourceID      string    `json:"customer_source_id" query:"customer_source_id" gorm:"type:varchar(36)"`
	CustomerSubSourceID   string    `json:"customer_sub_source_id" query:"customer_sub_source_id" gorm:"type:varchar(36)"`
	ReferralBy            string    `json:"referral_by" query:"referral_by" gorm:"type:varchar(50)"`
	CustomerTypeID        string    `json:"customer_type_id" query:"customer_type_id" gorm:"type:varchar(2)"`
	CustomerSignature     string    `json:"customer_signature" query:"customer_signature" gorm:"type:varchar(500)"`
	CustomerSignatureDate time.Time `json:"customer_signature_date" query:"customer_signature_date"`
	EmployeeSignature     string    `json:"employee_signature" query:"employee_signature" gorm:"type:varchar(500)"`
	EmployeeSignatureDate time.Time `json:"employee_signature_date" query:"employee_signature_date"`
	ManagerSignature      string    `json:"manager_signature" query:"manager_signature" gorm:"type:varchar(500)"`
	ManagerSignatureDate  time.Time `json:"manager_signature_date" query:"manager_signature_date"`
	Moo                   string    `json:"moo" query:"moo" gorm:"type:varchar(4)"`
	Soi                   string    `json:"soi" query:"soi" gorm:"type:varchar(100)"`
	Floor                 string    `json:"floor" query:"floor" gorm:"type:varchar(3)"`

	CompanyRegisterTypeID string `json:"company_register_type_id" query:"company_register_type_id" gorm:"type:varchar(2)"`

	ManagerID                       string    `json:"manager_id" query:"manager_id" gorm:"type:varchar(36)"`
	TelephoneCompany                string    `json:"telephone_company" query:"telephone_company" gorm:"type:varchar(20)"`
	IsConsent                       bool      `json:"is_consent" query:"is_consent" gorm:"type:bool"`
	CustomerSignatureLink           string    `json:"customer_signature_link" query:"customer_signature_link" gorm:"type:varchar(500)"`
	CustomerPDFLink                 string    `json:"customer_pdf_link" query:"customer_pdf_link" gorm:"type:varchar(500)"`
	PDFLink                         string    `json:"pdf_link" query:"pdf_link" gorm:"type:varchar(500)"`
	PDFLinkSendDate                 time.Time `json:"pdf_link_send_date" query:"pdf_link_send_date" gorm:"type:date"`
	CustomerSignatureLinkSendDate   time.Time `json:"customer_signature_link_send_date" query:"customer_signature_link_send_date" gorm:"type:date"`
	CustomerSignatureLinkExpireDate time.Time `json:"customer_signature_link_expire_date" query:"customer_signature_link_expire_date" gorm:"type:date"`
	IsSendBookingPDF                bool      `json:"is_send_booking_pdf" query:"is_send_booking_pdf" gorm:"type:bool"`
	SendBookingPDFDate              time.Time `json:"send_booking_pdf_date" query:"send_booking_pdf_date" gorm:"type:date"`
	PDFLinkExpireDate               time.Time `json:"pdf_link_expire_date" query:"pdf_link_expire_date" gorm:"type:date"`
	IsQuatation                     bool      `json:"is_quatation" query:"is_quatation" gorm:"type:bool"`
	QuatationName                   string    `json:"quatation_name" query:"quatation_name" gorm:"type:text"`
}
