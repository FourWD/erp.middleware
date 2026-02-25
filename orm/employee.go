package orm

import (
	"time"

	midOrm "github.com/FourWD/middleware/orm"
)

type Employee struct {
	midOrm.Employee
	RoleDetail string `json:"role_detail" query:"role_detail" gorm:"type:text"`
	CompanyID  string `json:"company_id" query:"company_id" gorm:"type:varchar(36)"`
	AcceptWEB  bool   `json:"accept_web" query:"accept_web" `
	AcceptBO   bool   `json:"accept_bo" query:"accept_bo" `
	SourceList string `json:"source_list" query:"source_list" gorm:"type:varchar(100)"`

	//Booking
	PowerOfAttorneyDate time.Time `json:"power_of_attorney_date" query:"power_of_attorney_date"`
	PowerOfAttorneyImg1 string    `json:"power_of_attorney_img_1" query:"power_of_attorney_img_1" gorm:"column:power_of_attorney_img_1;type:text"`
	PowerOfAttorneyImg2 string    `json:"power_of_attorney_img_2" query:"power_of_attorney_img_2" gorm:"column:power_of_attorney_img_2;type:text"`
	Signature           string    `json:"signature" query:"signature" gorm:"type:varchar(500)"`
	RefreshToken        string    `json:"refresh_token" query:"refresh_token" gorm:"type:varchar(500)"`

	ProfileImagePath string `json:"profile_image_path" query:"profile_image_path" gorm:"type:varchar(1000)"`
}
