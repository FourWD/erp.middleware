package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type VehicleApiOther struct {
	orm.Vehicle
	VehicleTypeID string `json:"vehicle_type_id" query:"vehicle_type_id" gorm:"type:varchar(10)"`
	// VehicleSubTypeID string `json:"vehicle_sub_type_id" query:"vehicle_sub_type_id" gorm:"type:varchar(2)"`
	VehicleBrandID        string    `json:"vehicle_brand_id" query:"vehicle_brand_id" gorm:"type:varchar(36)"`
	VehicleStatus         string    `json:"vehicle_status" query:"vehicle_status" gorm:"type:varchar(2)"`
	PricePreVat           float64   `json:"price_pre_vat" query:"price_pre_vat" gorm:"type:decimal(14,2)"`
	Vat                   float64   `json:"vat" query:"vat" gorm:"type:decimal(14,2)"`
	PriceIncVat           float64   `json:"price_inc_vat" query:"price_inc_vat" gorm:"type:decimal(14,2)"`
	VehicleRegisterTypeID string    `json:"vehicle_register_type_id" query:"vehicle_register_type_id" gorm:"type:varchar(2)"`
	TagID                 string    `json:"tag_id" query:"tag_id" gorm:"type:varchar(2)"`
	VehiclePickupDate     time.Time `json:"vehicle_pickup_date" query:"vehicle_pickup_date"`

	// Remark
	RemarkEngineGearDrivengear string `json:"remark_engine_gear_drivengear" query:"remark_engine_gear_drivengear" gorm:"type:text"`
	RemarkExternal             string `json:"remark_external" query:"remark_external" gorm:"type:text"`
	RemarkInterior             string `json:"remark_interior" query:"remark_interior" gorm:"type:text"`
	RemarkChassisColor         string `json:"remark_chassis_color" query:"remark_chassis_color" gorm:"type:text"`
	RemarkServiceList          string `json:"remark_service_list" query:"remark_service_list" gorm:"type:text"`
	RemarkEngineSystem         string `json:"remark_engine_system" query:"remark_engine_system" gorm:"type:text"`
	RemarkAirSystem            string `json:"remark_air_system" query:"remark_air_system" gorm:"type:text"`
	RemarkColor                string `json:"remark_color" query:"remark_color" gorm:"type:text"`
	RemarkEtc                  string `json:"remark_etc" query:"remark_etc" gorm:"type:text"` // อื่น ๆ

	// name from api import
	VehicleBranchCode   string `json:"vehicle_branch_code" query:"vehicle_branch_code" gorm:"type:varchar(255)"`
	VehicleTypeName     string `json:"vehicle_type_name" query:"vehicle_type_name" gorm:"type:varchar(255)"`
	VehicleBrandName    string `json:"vehicle_brand_name" query:"vehicle_brand_name" gorm:"type:varchar(255)"`
	VehicleModelName    string `json:"vehicle_model_name" query:"vehicle_model_name" gorm:"type:varchar(255)"`
	VehicleSubModelName string `json:"vehicle_sub_model_name" query:"vehicle_sub_model_name" gorm:"type:varchar(255)"`
	VehicleColorName    string `json:"vehicle_color_name" query:"vehicle_color_name" gorm:"type:varchar(255)"`
	LicenseProvinceName string `json:"license_province_name" query:"license_province_name" gorm:"type:varchar(255)"`
	CarStatus           string `json:"car_status" query:"car_status" gorm:"type:varchar(255)"`

	// CostMA int // ต้นทุนการปรับสภาพรถ
}

/*
CREATE VIEW `vehicles` AS
SELECT vehicles.* ,
vm.vehicle_brand_id,
vst.vehicle_type_id,
s.name AS source_name,
b.name AS branch_name,
vt.name AS vehicle_type_name,
vst.name AS vehicle_sub_type_name,
vb.name AS vehicle_brand_name,
vm.name AS vehicle_model_name,
vsm.name AS vehicle_sub_model_name,
vdt.name AS vehicle_drive_type_name,
vg.name AS vehicle_gear_name,
vft.name AS vehicle_fuel_type_name,
vc.name AS vehicle_color_name,
vgr.row_order AS vehicle_grade_value,
vgr.name AS vehicle_grade_name,
vrt.code AS vehicle_register_type_code,
vrt.name AS vehicle_register_type_name,
p.name AS license_province_name ,
vgrr.grade_remark AS grade_remark
FROM (
select * from `auction-vehicle-0001`.vehicles
UNION
select * from `auction-vehicle-0002`.vehicles
UNION
select * from `auction-vehicle-0003`.vehicles
UNION
select * from `auction-vehicle-0004`.vehicles
UNION
select * from `auction-vehicle-0005`.vehicles
UNION
select * from `auction-vehicle-0006`.vehicles
UNION
select * from `auction-vehicle-0007`.vehicles
UNION
select * from `auction-vehicle-0008`.vehicles
UNION
select * from `auction-vehicle-0009`.vehicles
UNION
select * from `auction-vehicle-0010`.vehicles
UNION
select * from `auction-vehicle-0011`.vehicles
UNION
select * from `auction-vehicle-0012`.vehicles
UNION
select * from `auction-vehicle-0013`.vehicles
UNION
select * from `auction-vehicle-0014`.vehicles
UNION
select * from `auction-vehicle-0015`.vehicles
UNION
select * from `auction-vehicle-0016`.vehicles
UNION
select * from `auction-vehicle-0017`.vehicles
UNION
select * from `auction-vehicle-0018`.vehicles
UNION
select * from `auction-vehicle-0019`.vehicles
UNION
select * from `auction-vehicle-0020`.vehicles
UNION
select * from `auction-vehicle-0021`.vehicles
UNION
select * from `auction-vehicle-0022`.vehicles
UNION
select * from `auction-vehicle-0023`.vehicles
UNION
select * from `auction-vehicle-0024`.vehicles
UNION
select * from `auction-vehicle-0025`.vehicles
UNION
select * from `auction-vehicle-0026`.vehicles
UNION
select * from `auction-vehicle-0027`.vehicles
UNION
select * from `auction-vehicle-0028`.vehicles
UNION
select * from `auction-vehicle-0029`.vehicles
UNION
select * from `auction-vehicle-0030`.vehicles
) AS vehicles
LEFT JOIN vehicle_sub_models vsm ON vsm.id = vehicles.vehicle_sub_model_id
LEFT JOIN vehicle_models vm ON vm.id = vsm.vehicle_model_id
LEFT JOIN vehicle_sub_types vst ON vst.id = vehicles.vehicle_sub_type_id
LEFT JOIN vehicle_types vt ON vt.id = vst.vehicle_type_id
LEFT JOIN vehicle_brands vb ON vb.id = vm.vehicle_brand_id
LEFT JOIN vehicle_drive_types vdt ON vdt.id = vehicles.vehicle_drive_type_id
LEFT JOIN vehicle_gears vg ON vg.id = vehicles.vehicle_gear_id
LEFT JOIN vehicle_fuel_types vft ON vft.id = vehicles.vehicle_fuel_type_id
LEFT JOIN vehicle_colors vc ON vc.id = vehicles.vehicle_color_id
LEFT JOIN vehicle_grades vgr ON vgr.id = vehicles.vehicle_grade_id
LEFT JOIN sources s ON s.id = vehicles.source_id
LEFT JOIN branches b ON b.id = vehicles.branch_id
LEFT JOIN provinces p ON p.id = vehicles.license_province_id
LEFT JOIN vehicle_register_types vrt ON vehicles.vehicle_register_type_id = vrt.id
LEFT JOIN vehicle_grade_remarks vgrr ON vehicles.chassis_no = vgrr.chassis_number;
*/
