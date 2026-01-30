package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Users struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);uniqueIndex:idx_id;primary_key"`
	model.GormModel

	Code         string `json:"code" query:"code" gorm:"type:varchar(10)"`
	UserTypeID   string `json:"user_type_id" query:"user_type_id" gorm:"type:varchar(2)"`
	FirstName    string `json:"first_name" query:"name" gorm:"type:varchar(256)"`
	LastName     string `json:"last_name" query:"name" gorm:"type:varchar(256)"`
	Username     string `json:"username" query:"username" gorm:"type:varchar(20)"`
	Password     string `json:"password" query:"password" gorm:"text"`
	FileAvatarID string `json:"file_avartar_id" query:"file_avartar_id" gorm:"type:varchar(36)"`
	Mobile       string `json:"mobile" query:"mobile" gorm:"type:varchar(10)"`
	Email        string `json:"email" query:"email" gorm:"type:varchar(50)"`
	Facebook     string `json:"facebook" query:"facebook" gorm:"type:varchar(50)"`
	Line         string `json:"line" query:"line" gorm:"type:varchar(20)"`
	Token        string `json:"token" query:"token" gorm:"type:text"`


	//มาจาก booking
	PrefixID         string     `json:"prefix_id" query:"prefix_id" gorm:"type:varchar(2)"`
	GenderID         string     `json:"gender_id" query:"gender_id" gorm:"type:varchar(2)"`
	BirthDate        *time.Time `json:"birth_date" query:"birth_date" `
	IdCardNo         string     `json:"id_card_no" query:"id_card_no" gorm:"type:varchar(20)"`
	Address       string `json:"address" query:"address" gorm:"type:text"`
	DistrictID    string `json:"district_id" query:"district_id" gorm:"type:varchar(4)"`         //อำเภอ
	SubDistrictID string `json:"sub_district_id" query:"sub_district_id" gorm:"type:varchar(6)"` //ตำบล
	ProvinceID    string `json:"province_id" query:"province_id" gorm:"type:varchar(2)"`
	Postcode      string `json:"postcode" query:"postcode" gorm:"type:varchar(5)"`


	// Job Description Section
	CareerID    string `json:"career_id" query:"career_id" gorm:"type:varchar(36)"`
	CareerSubID string `json:"career_sub_id" query:"career_sub_id" gorm:"type:varchar(36)"`
	CareerSalary int `json:"career_salary" query:"career_salary" gorm:"type:int"`

}
