package orm

import (
	"time"
)

type VehicleCsvImport struct {
	IDate                   time.Time `json:"i_date" query:"i_date" `
	ID                      int       `json:"id" query:"id" gorm:"primaryKey"`
	SourceID                string    `json:"source_id" query:"source_id"  `
	SaleGroup               string    `json:"sale_group" query:"sale_group" `
	VehicleGradeID          string    `json:"vehicle_grade_id" query:"vehicle_grade_id" `
	YearRegister            string    `json:"year_register" query:"year_register" `
	CarMonthReg             string    `json:"car_month_reg" query:"car_month_reg" `
	YearManufacturing       string    `json:"year_manufacturing" query:"year_manufacturing" `
	VehicleBrandName        string    `json:"vehicle_brand_name" query:"vehicle_brand_name" `
	VehicleModelName        string    `json:"vehicle_model_name" query:"vehicle_model_name" `
	VehicleSubModelName     string    `json:"vehicle_sub_model_name" query:"vehicle_sub_model_name" `
	VehicleDriveTypeName    string    `json:"vehicle_drive_type_name" query:"vehicle_drive_type_name" `
	VehicleGearID           string    `json:"vehicle_gear_id" query:"vehicle_gear_id" `
	Option                  string    `json:"option" query:"option" `
	VehicleTypeName         string    `json:"vehicle_type_name" query:"vehicle_type_name" `
	VehicleColorName        string    `json:"vehicle_color_name" query:"vehicle_color_name" `
	EngineCapacity          string    `json:"engine_capacity" query:"engine_capacity" `
	VehicleFuelTypeID       string    `json:"vehicle_fuel_type_id" query:"vehicle_fuel_type_id" `
	Seat                    string    `json:"seat" query:"seat" `
	Mile                    string    `json:"mile" query:"mile" `
	License                 string    `json:"license" query:"license" `
	LicenseProvinceName     string    `json:"license_province_name" query:"license_province_name" `
	LicenseReceiveDate      string    `json:"license_receive_date" query:"license_receive_date" `
	LicenseExpireDate       string    `json:"license_expire_date" query:"license_expire_date" `
	VehicleRegisterTypeName string    `json:"vehicle_register_type_name" query:"vehicle_register_type_name" `
	VehicleRegisterTypeID   string    `json:"vehicle_register_type_id" query:"vehicle_register_type_id" `
	InsuranceExpired        string    `json:"insurance_expired" query:"insurance_expired" `
	ChassisNo               string    `json:"chassis_no" query:"chassis_no" gorm:"index;column:chassis_no"`
	EngineNo                string    `json:"engine_no" query:"engine_no" `
	EngineSizeActual        string    `json:"engine_size_actual" query:"engine_size_actual" `
	EngineSize              string    `json:"engine_size" query:"engine_size" `
	Price                   float64   `json:"price" query:"price" gorm:"type:decimal(14,2)" `
	Carpark                 string    `json:"carpark" query:"carpark" `
	CustomerLast            string    `json:"customer_last" query:"customer_last" `
	CarStatus               string    `json:"car_status" query:"car_status" gorm:"index"`
	BranchID                string    `json:"branch_id" query:"branch_id" `
	VehicleTypeID           string    `json:"vehicle_type_id" query:"vehicle_type_id" `
	VehicleBrandID          string    `json:"vehicle_brand_id" query:"vehicle_brand_id" `
	VehicleModelID          string    `json:"vehicle_model_id" query:"vehicle_model_id" `
	VehicleSubModelID       string    `json:"vehicle_sub_model_id" query:"vehicle_sub_model_id" `
	VehicleColorID          string    `json:"vehicle_color_id" query:"vehicle_color_id" `
	VehicleGradeIDRef       string    `json:"vehicle_grade_id_ref" query:"vehicle_grade_id_ref" `
	VehicleDriveTypeID      string    `json:"vehicle_drive_type_id" query:"vehicle_drive_type_id" `
	VehicleFuelTypeIDRef    string    `json:"vehicle_fuel_type_id_ref" query:"vehicle_fuel_type_id_ref" `
	CarID                   string    `json:"car_id" query:"car_id" `
	TypeName                string    `json:"type_name" query:"type_name" `
}
