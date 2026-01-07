package orm

import (
	"time"
)

type DuckbookVehicle struct {
	ID                 string     `json:"id" query:"id" gorm:"type:varchar(36)"`
	SKU                string     `json:"sku" query:"sku" gorm:"type:varchar(50)"`
	Source             string     `json:"source" query:"source" gorm:"type:varchar(20)"`
	BrandName          string     `json:"brand_name" query:"brand_name" gorm:"type:varchar(100)"`
	ModelName          string     `json:"model_name" query:"model_name" gorm:"type:varchar(100)"`
	SubModelName       string     `json:"sub_model_name" query:"sub_model_name" gorm:"type:varchar(100)"`
	MasterBrandName    string     `json:"master_brand_name" query:"master_brand_name" gorm:"type:varchar(100)"`
	MasterModelName    string     `json:"master_model_name" query:"master_model_name" gorm:"type:varchar(100)"`
	MasterSubModelName string     `json:"master_sub_model_name" query:"master_sub_model_name" gorm:"type:varchar(100)"`
	YearManufacturing  int        `json:"year_manufacturing" query:"year_manufacturing" gorm:"type:int"`
	Price              float64    `json:"price" query:"price" gorm:"type:decimal(10,2)"`
	OpenPrice          float64    `json:"open_price" query:"open_price" gorm:"type:decimal(10,2)"`
	ClosePrice         float64    `json:"close_price" query:"close_price" gorm:"type:decimal(10,2)"`
	Name               string     `json:"name" query:"name" gorm:"type:text"`
	WebRemark          string     `json:"web_remark" query:"web_remark" gorm:"type:text"`
	UpdatedAt          *time.Time `json:"updated_at" query:"updated_at"`
}
