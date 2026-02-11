package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

// midOrm "github.com/FourWD/middleware/orm"

type BookingVehicle struct { //
	ID string `json:"id" query:"id" gorm:"type:varchar(2); primary_key"`
	model.GormModel
	BookingID                string     `json:"booking_id" query:"booking_id" gorm:"type:varchar(36)"`
	ChassisNo                string     `json:"chassis_no" query:"chassis_no" gorm:"type:varchar(36)"`
	EngineNo                 string     `json:"engine_no" query:"engine_no" gorm:"type:varchar(36)"`
	VehicleModelID           string     `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36)"`
	VehicleSubModelID        string     `json:"vehicle_sub_model_id" query:"vehicle_sub_model_id" gorm:"type:varchar(36)"`
	VehicleColorID           string     `json:"vehicle_color_id" query:"vehicle_color_id" gorm:"type:varchar(36)"`
	VehicleYearManufacturing string     `json:"vehicle_year_manufacturing" query:"vehicle_year_manufacturing" gorm:"type:varchar(4)"`
	ExpectDeliveryDate       *time.Time `json:"expect_delivery_date" query:"expect_delivery_date"` // วันที่ต้องการจัดส่ง
	DeliveryDate             *time.Time `json:"delivery_date" query:"delivery_date" gorm:"type:date"`
}
