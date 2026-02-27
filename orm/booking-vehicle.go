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
	PoNo                     string     `json:"po_no" query:"po_no" gorm:"type:varchar(20)"`
	ChassisNo                string     `json:"chassis_no" query:"chassis_no" gorm:"type:varchar(36)"`
	EngineNo                 string     `json:"engine_no" query:"engine_no" gorm:"type:varchar(36)"`
	VehicleModelID           string     `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36)"`
	VehicleSubModelID        string     `json:"vehicle_sub_model_id" query:"vehicle_sub_model_id" gorm:"type:varchar(36)"`
	VehicleColorID           string     `json:"vehicle_color_id" query:"vehicle_color_id" gorm:"type:varchar(36)"`
	VehicleYearManufacturing string     `json:"vehicle_year_manufacturing" query:"vehicle_year_manufacturing" gorm:"type:varchar(4)"`
	ExpectDeliveryDate       *time.Time `json:"expect_delivery_date" query:"expect_delivery_date"`                    // วันที่ต้องการจัดส่ง
	PaymentMethodID          string     `json:"payment_method_id" query:"payment_method_id" gorm:"type:varchar(36)"`  // ประเภทการซื้อ ไฟแนน หรือ เงินสด
	DeliveryDate             *time.Time `json:"delivery_date" query:"delivery_date" gorm:"type:date"`                 // วันที่จัดส่ง
	DeliveryLocation         string     `json:"delivery_location" query:"delivery_location" gorm:"type:varchar(500)"` // สถานที่จัดส่ง
	PODate                   *time.Time `json:"po_date" query:"po_date" gorm:"type:date"`
	UploadPath1              string     `json:"upload_path_1" query:"upload_path_1" gorm:"column:upload_path_1;type:varchar(500)"`
	UploadPath2              string     `json:"upload_path_2" query:"upload_path_2" gorm:"column:upload_path_2;type:varchar(500)"`
	UploadPath3              string     `json:"upload_path_3" query:"upload_path_3" gorm:"column:upload_path_3;type:varchar(500)"`
	UploadPath4              string     `json:"upload_path_4" query:"upload_path_4" gorm:"column:upload_path_4;type:varchar(500)"`
	UploadPath5              string     `json:"upload_path_5" query:"upload_path_5" gorm:"column:upload_path_5;type:varchar(500)"`
	UploadPath6              string     `json:"upload_path_6" query:"upload_path_6" gorm:"column:upload_path_6;type:varchar(500)"`
	UploadPath7              string     `json:"upload_path_7" query:"upload_path_7" gorm:"column:upload_path_7;type:varchar(500)"`
	UploadPath8              string     `json:"upload_path_8" query:"upload_path_8" gorm:"column:upload_path_8;type:varchar(500)"`
	UploadPath9              string     `json:"upload_path_9" query:"upload_path_9" gorm:"column:upload_path_9;type:varchar(500)"`
	UploadPath10             string     `json:"upload_path_10" query:"upload_path_10" gorm:"column:upload_path_10;type:varchar(500)"`
	BookingStatusID          string     `json:"booking_status_id" query:"booking_status_id" gorm:"type:varchar(36)"`
}
