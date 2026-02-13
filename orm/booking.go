package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Booking struct {
	// midOrm.Booking // ไม่ต้องใส่ค่าลง vehicle_model_id, vehicle_sub_model_id ใช้จอยแทน
	ID string `json:"id" query:"id" gorm:"type:varchar(36);uniqueIndex:idx_id"`

	model.GormModel

	BookingNo           string     `json:"booking_no" query:"booking_no" gorm:"type:varchar(20)"` //หมายเลขรายการ
	PoNo                string     `json:"po_no" query:"po_no" gorm:"type:varchar(20)"`
	BookingDate         *time.Time `json:"booking_date" query:"booking_date" `
	BookingBy           string     `json:"booking_by" query:"booking_by" gorm:"type:varchar(36)"`
	BookingStatusID     string     `json:"booking_status_id" query:"booking_status_id" gorm:"type:varchar(2)"`
	BookingCancelTypeID string     `json:"booking_cancel_type_id" query:"booking_cancel_type_id" gorm:"type:varchar(2)"`
	BookingCancelDate   *time.Time `json:"booking_cancel_date" query:"booking_cancel_date" `
	BookingCancelRemark string     `json:"booking_cancel_remark" query:"booking_cancel_remark" gorm:"type:text"`

	EmployeeID string `json:"employee_id" query:"employee_id" gorm:"type:varchar(36)"` // sale_contact_id
	BranchID   string `json:"branch_id" query:"branch_id" gorm:"type:varchar(36)"`

	CustomerID        string `json:"customer_id" query:"customer_id" gorm:"type:varchar(36)"`
	CustomerUseTypeID string `json:"customer_use_type_id" query:"customer_use_type_id" gorm:"type:varchar(2)"` //เหตุผลการซื้อทำ master

	// show user_type_id
	ReferralChannelID string     `json:"referral_channel_id" query:"referral_channel_id" gorm:"type:varchar(2)"`
	ReferralName      string     `json:"referral_name" query:"referral_name" gorm:"type:varchar(256)"`
	ReserveDate       *time.Time `json:"reserve_date" query:"reserve_date"`                                 // วันที่รับเงินจอง
	ReserveBy         string     `json:"reserve_by" query:"reserve_by" gorm:"type:varchar(36)"`             // คนรับเงินจอง
	ReserveDeposit    float64    `json:"reserve_deposit" query:"reserve_deposit" gorm:"type:decimal(14,2)"` // จำนวนเงินจอง
	// CaseSubmissionDate *time.Time `json:"case_submission_date" query:"case_submission_date"`                 // วันที่ส่งเคส
	ExpectDeliveryDate *time.Time `json:"expect_delivery_date" query:"expect_delivery_date"` // วันที่ต้องการจัดส่ง
	DeliveryDate       *time.Time `json:"delivery_date" query:"delivery_date" gorm:"type:date"`
	DeliveryLocation   string     `json:"delivery_location" query:"delivery_location" gorm:"type:varchar(500)"`
	RemarkDelivery     string     `json:"remark_delivery" query:"remark_delivery" gorm:"type:text"`

	IsDownPaid  bool    `json:"is_down_paid" query:"is_down_paid" gorm:"type:bool"` // จ่ายเงินจองเรียบร้อย
	IsCancel    bool    `json:"is_cancel" query:"is_cancel" gorm:"type:bool"`
	Remark      string  `json:"remark" query:"remark" gorm:"type:text"`
	PricePreVat float64 `json:"price_pre_vat" query:"price_pre_vat" gorm:"type:decimal(14,2)"`
	Vat         float64 `json:"vat" query:"vat" gorm:"type:decimal(14,2)"`
	Price       float64 `json:"price" query:"price" gorm:"type:decimal(14,2)"`
	// k'pom special discount
	Discount              float64    `json:"discount" query:"discount" gorm:"type:decimal(14,2)"`
	SpecialDiscount       float64    `json:"special_discount" query:"is_special_discount" gorm:"type:decimal(14,2)"`
	SpecialDiscountBy     string     `json:"special_discount_by" query:"special_discount_by" gorm:"type:varchar(36)"`
	SpecialDiscountDate   *time.Time `json:"special_discount_date" query:"special_discount_date" gorm:"type:date"`
	SpecialDescountRef    string     `json:"special_descount_ref" query:"special_descount_ref" gorm:"type:text"`
	SpecialDiscountRemark string     `json:"special_discount_remark" query:"special_discount_remark" gorm:"type:text"`

	// discount booking
	// DiscountVehicle        float64 `json:"discount_vehicle" query:"discount_vehicle" gorm:"type:decimal(14,2)"`
	// DiscountDown           float64 `json:"discount_down" query:"discount_down" gorm:"type:decimal(14,2)"`
	// DiscountAccessory      float64 `json:"discount_accessory" query:"discount_accessory" gorm:"type:decimal(14,2)"`
	DepositRedLicensePlate float64 `json:"deposit_red_license_plate" query:"deposit_red_license_plate" gorm:"type:decimal(14,2)"`
	DepositRegistationFee  float64 `json:"deposit_registation_fee" query:"deposit_registation_fee" gorm:"type:decimal(14,2)"`
	TotalPrice             float64 `json:"total_price" query:"total_price" gorm:"type:decimal(14,2)"`

	// Website Section Required
	RunningNo         int    `json:"running_no" query:"running_no" gorm:"primary_key;auto_increment;not_null"`
	IsSelectedFinance string `json:"is_selected_finance" query:"is_selected_finance" gorm:"type:varchar(10)"`

	// Booking System
	BookingGroupID           string `json:"booking_group_id" query:"booking_group_id" gorm:"type:varchar(36)"` // เอาไว้จัดกรุปรายการจองรถหลายคัน
	IdCardPath               string `json:"id_card_path" query:"id_card_path" gorm:"type:text"`
	VehicleID                string `json:"vehicle_id" query:"vehicle_id" gorm:"type:varchar(36)"`
	ChassisNo                string `json:"chassis_no" query:"chassis_no" gorm:"type:varchar(36)"`
	EngineNo                 string `json:"engine_no" query:"engine_no" gorm:"type:varchar(36)"`
	VehicleModelID           string `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36)"`
	VehicleSubModelID        string `json:"vehicle_sub_model_id" query:"vehicle_sub_model_id" gorm:"type:varchar(36)"`
	VehicleColorID           string `json:"vehicle_color_id" query:"vehicle_color_id" gorm:"type:varchar(36)"`
	VehicleYearManufacturing string `json:"vehicle_year_manufacturing" query:"vehicle_year_manufacturing" gorm:"type:varchar(4)"`
	UploadPath1              string `json:"upload_path_1" query:"upload_path_1" gorm:"type:varchar(500)"`
	UploadPath2              string `json:"upload_path_2" query:"upload_path_2" gorm:"type:varchar(500)"`
	UploadPath3              string `json:"upload_path_3" query:"upload_path_3" gorm:"type:varchar(500)"`
	UploadPath4              string `json:"upload_path_4" query:"upload_path_4" gorm:"type:varchar(500)"`
	UploadPath5              string `json:"upload_path_5" query:"upload_path_5" gorm:"type:varchar(500)"`
}
