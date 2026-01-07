package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type Branch struct {
	midOrm.Branch
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
}
