package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Customer struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);uniqueIndex:idx_id;primary_key"`
	model.GormModel

	Code           string `json:"code" query:"code" gorm:"type:varchar(10)"`
	CustomerTypeID string `json:"customer_type_id" query:"customer_type_id" gorm:"type:varchar(2)"`
	FirstName      string `json:"first_name" query:"name" gorm:"type:varchar(256)"`
	LastName       string `json:"last_name" query:"name" gorm:"type:varchar(256)"`
	Username       string `json:"username" query:"username" gorm:"type:varchar(20)"`
	Password       string `json:"password" query:"password" gorm:"text"`
	FileAvatarID   string `json:"file_avartar_id" query:"file_avartar_id" gorm:"type:varchar(36)"`
	Mobile         string `json:"mobile" query:"mobile" gorm:"type:varchar(10)"`
	Email          string `json:"email" query:"email" gorm:"type:varchar(50)"`
	Facebook       string `json:"facebook" query:"facebook" gorm:"type:varchar(50)"`
	Line           string `json:"line" query:"line" gorm:"type:varchar(20)"`
	Token          string `json:"token" query:"token" gorm:"type:text"`

	//มาจาก booking
	PrefixID  string     `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2)"`
	GenderID  string     `json:"gender_id" query:"gender_id" gorm:"type:varchar(2)"`
	BirthDate *time.Time `json:"birth_date" query:"birth_date" `
	IdCardNo  string     `json:"id_card_no" query:"id_card_no" gorm:"type:varchar(20)"`

	// Address Section
	Address       string `json:"address" query:"address" gorm:"type:text"`
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
	ContactAddress       string `json:"contact_address" query:"contact_address" gorm:"type:text"`
	ContactFloor         string `json:"contact_floor" query:"contact_floor" gorm:"type:varchar(5)"`
	ContactRoom          string `json:"contact_room" query:"contact_room" gorm:"type:varchar(20)"`
	ContactMoo           string `json:"contact_moo" query:"contact_moo" gorm:"type:varchar(5)"`
	ContactSoi           string `json:"contact_soi" query:"contact_soi" gorm:"type:varchar(200)"`
	ContactRoad          string `json:"contact_road" query:"contact_road" gorm:"type:varchar(200)"`
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
	ContectPersonNickname     string `json:"contect_person_nickname" query:"contect_person_nickname" gorm:"type:varchar(500)"`
	ContactPersonPosition     string `json:"contact_person_position" query:"contact_person_position" gorm:"type:varchar(500)"`
	ContactPersonEmail        string `json:"contact_person_email" query:"contact_person_email" gorm:"type:varchar(500)"`
	ContactPersonLineID       string `json:"contact_person_line_id" query:"contact_person_line_id" gorm:"type:varchar(500)"`
	ContactPersonFacebook     string `json:"contact_person_facebook" query:"contact_person_facebook" gorm:"type:varchar(500)"`

	IdCardPath      string `json:"id_card_path" query:"id_card_path" gorm:"type:text"`
	IsAcceptConsent bool   `json:"is_accept_consent" query:"is_accept_consent" gorm:"type:bool"`
	ConsentID       string `json:"consent_id" query:"consent_id" gorm:"type:varchar(36)"`
}
