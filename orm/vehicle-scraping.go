package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type VehicleScraping struct {
	orm.Vehicle
	Name              string     `json:"name" query:"name" gorm:"type:varchar(500)"`
	Web               string     `json:"web" query:"web" gorm:"type:varchar(256)"`
	WebID             string     `json:"web_id" query:"web_id" gorm:"type:varchar(36)"`
	SKU               string     `json:"sku" query:"sku" gorm:"type:varchar(256)"`
	VehicleTypeID     string     `json:"vehicle_type_id" query:"vehicle_type_id" gorm:"type:varchar(10)"`
	VehicleBrandID    string     `json:"vehicle_brand_id" query:"vehicle_brand_id" gorm:"type:varchar(36)"`
	Brand             string     `json:"brand" query:"brand" gorm:"type:varchar(256)"`
	Model             string     `json:"model" query:"model" gorm:"type:varchar(256)"`
	SubModel          string     `json:"sub_model" query:"sub_model" gorm:"type:varchar(256)"`
	Gear              string     `json:"gear" query:"gear" gorm:"type:varchar(50)"`
	Province          string     `json:"province" query:"province" gorm:"type:varchar(256)"`
	Dealer            string     `json:"dealer" query:"dealer" gorm:"type:longtext"`
	Remark            string     `json:"remark" query:"remark" gorm:"type:varchar(500)"`
	Mile              string     `json:"mile" query:"mile" gorm:"type:varchar(256)"`
	ModelYear         string     `json:"model_year" query:"model_year" gorm:"type:varchar(256)"`
	YearManufacturing string     `json:"year_manufacturing" query:"year_manufacturing" gorm:"type:varchar(4)"`
	PricePreVat       string     `json:"price_pre_vat" query:"price_pre_vat" gorm:"type:varchar(256)"`
	Price             string     `json:"price" query:"price" gorm:"type:varchar(256)"`
	BodyType          string     `json:"body_type" query:"body_type" gorm:"type:varchar(256)"`
	Fuel              string     `json:"fuel" query:"fuel" gorm:"type:varchar(256)"`
	Color             string     `json:"color" query:"color" gorm:"type:varchar(256)"`
	PublishDate       *time.Time `json:"publish_date" query:"publish_date" firestore:"publish_date" gorm:"default:null"`
	Link              string     `json:"link" query:"link" gorm:"type:varchar(256)"`
}
