package orm

import "github.com/FourWD/middleware/model"

type Branch struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	Code string `json:"code" query:"code" gorm:"type:varchar(2)"`

	Name    string `json:"name" query:"name" gorm:"type:varchar(100)"`
	NameEn  string `json:"name_en" query:"name_en" gorm:"type:varchar(100)"`
	Label   string `json:"label" query:"label" gorm:"type:varchar(100)"`
	LabelEn string `json:"label_en" query:"label_en" gorm:"type:varchar(100)"`

	TextColor       string  `json:"text_color" query:"text_color" gorm:"type:varchar(7)"`
	BackgroundColor string  `json:"background_color" query:"background_color" gorm:"type:varchar(7)"`
	Address         string  `json:"address" query:"address" gorm:"type:text"`
	ProvinceID      string  `json:"province_id" query:"province_id" gorm:"type:varchar(36)"`
	Phone1          string  `json:"phone_1" query:"phone_1" gorm:"type:varchar(20)"`
	Phone2          string  `json:"phone_2" query:"phone_2" gorm:"type:varchar(20)"`
	Line            string  `json:"line" query:"line" gorm:"type:varchar(20)"`
	Facebook        string  `json:"facebook" query:"facebook" gorm:"type:varchar(20)"`
	Tiktok          string  `json:"tiktok" query:"tiktok" gorm:"type:varchar(20)"`
	Latitude        float64 `json:"latitude" query:"latitude" gorm:"type:float"`
	Longitude       float64 `json:"longitude" query:"longitude" gorm:"type:float"`
	MapUrl          string  `json:"map_url" query:"map_url" gorm:"type:varchar(500)"`

	DistrictID    		string  `json:"district_id" query:"district_id" gorm:"type:varchar(4)"`
	SubDistrictID     	string  `json:"sub_district_id" query:"sub_district_id" gorm:"type:varchar(6)"`
	Street    			string  `json:"street" query:"street" gorm:"type:varchar(200)"`
	Moo    				string  `json:"moo" query:"moo" gorm:"type:varchar(5)"`
	Soi    				string 	`json:"soi" query:"soi" gorm:"type:varchar(200)"`
	Postcode    		string 	`json:"postcode" query:"postcode" gorm:"type:varchar(5)"`
	Building            string  `json:"building" query:"building" gorm:"type:varchar(100)"`
	Room               	string  `json:"room" query:"room" gorm:"type:varchar(20)"`
	Floor               string  `json:"floor" query:"floor" gorm:"type:varchar(3)"`
	Email               string  `json:"email" query:"email" gorm:"type:varchar(100)"`

	OptionalID1 string `json:"optional_id_1" query:"optional_id_1" gorm:"column:optional_id_1;type:varchar(20)"`
	OptionalID2 string `json:"optional_id_2" query:"optional_id_2" gorm:"column:optional_id_2;type:varchar(20)"`
	OptionalID3 string `json:"optional_id_3" query:"optional_id_3" gorm:"column:optional_id_3;type:varchar(20)"`

	model.GormRowOrder
	CompanyID string `json:"company_id" query:"company_id" gorm:"type:varchar(36)"`
}
