package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type BookingPayment struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key"`
	model.GormModel

	IsAccessory            bool      `json:"is_accessory" query:"is_accessory" gorm:"type:bool"`
	BookingID              string    `json:"booking_id" query:"booking_id" gorm:"type:varchar(36)"`
	VehiclePrice           float64   `json:"vehicle_price" query:"vehicle_price" gorm:"type:decimal(14,2)"`
	VehicleDiscountPrice   float64   `json:"vehicle_discount_price" query:"vehicle_discount_price" gorm:"type:decimal(14,2)"`
	VehicleTotalPrice      float64   `json:"vehicle_total_price" query:"vehicle_total_price" gorm:"type:decimal(14,2)"`
	AccessoryPrice         float64   `json:"accessory_price" query:"accessory_price" gorm:"type:decimal(14,2)"`
	AccessoryDiscountPrice float64   `json:"accessory_discount_price" query:"accessory_discount_price" gorm:"type:decimal(14,2)"`
	AccessoryTotalPrice    float64   `json:"accessory_total_price" query:"accessory_total_pric" gorm:"type:decimal(14,2)"`
	TotalPrice             float64   `json:"total_price" query:"total_price" gorm:"type:decimal(14,2)"`
	AddPriceVehicle        float64   `json:"add_price_vehicle" query:"add_price_vehicle" gorm:"type:decimal(14,2)"`
	VatVehiclePrice        float64   `json:"vat_vehicle_price" query:"vat_vehicle_price" gorm:"type:decimal(14,2)"`
	DownPrice              float64   `json:"down_price" query:"down_price" gorm:"type:decimal(14,2)"`
	DownPercent            float64   `json:"down_precent" query:"down_precent" gorm:"type:decimal(14,2)"`
	Discount               float64   `json:"discount" query:"discount" gorm:"type:decimal(14,2)"`
	FinancePrice           float64   `json:"finance_price" query:"finance_price" gorm:"type:decimal(14,2)"`
	FinanceMonth           int       `json:"finance_month" query:"finance_month" gorm:"type:int(3)"`
	IsFinanceMonth         bool      `json:"is_finance_month" query:"is_finance_month" gorm:"type:bool"`
	FinancePricePerMonth   float64   `json:"finance_price_per_month" query:"finance_price_per_month" gorm:"type:decimal(14,2)"`
	InterestTypeID         string    `json:"interest_type_id" query:"interest_type_id" gorm:"type:varchar(36)"`
	InsuranceFund          float64   `json:"insurance_fund" query:"insurance_fund" gorm:"type:decimal(14,2)"` //
	FinanceID              string    `json:"finance_id" query:"finance_id" gorm:"type:varchar(10)"`
	FinanceInterest        float64   `json:"finance_interest" query:"finance_interest" gorm:"type:decimal(5,3)"`
	InsuranceID            string    `json:"insurance_id" query:"insurance_id" gorm:"type:varchar(10)"`
	PaymentTypeID          string    `json:"payment_type_id" query:"payment_type_id" gorm:"type:varchar(2)"`
	DepositPrice           float64   `json:"deposit_price" query:"deposit_pricet" gorm:"type:decimal(14,2)"`
	BankID                 string    `json:"bank_id" query:"bank_id" gorm:"type:varchar(2)"`
	CreditCardNo           string    `json:"credit_card_no" query:"credit_card_no" gorm:"type:varchar(36)"` //mark-up
	CreditCardExpireDate   string    `json:"credit_card_expire_date" query:"credit_card_expire_date" gorm:"type:varchar(36)"`
	TransferDate           time.Time `json:"transfer_date" query:"transfer_date"`
	// PickupDate             time.Time `json:"pickup_date" query:"pickup_date"`
	// ExpectPickupDate       time.Time `json:"expect_pickup_date" query:"expect_pickup_date"`
	FileSlip               string    `json:"file_slip"   query:"file_slip"   gorm:"type:varchar(500)"`
	FileSlip1              string    `json:"file_slip_1" query:"file_slip_1" gorm:"column:file_slip_1;type:varchar(500)"`
	FileSlip2              string    `json:"file_slip_2" query:"file_slip_2" gorm:"column:file_slip_2;type:varchar(500)"`
	FileSlip3              string    `json:"file_slip_3" query:"file_slip_3" gorm:"column:file_slip_3;type:varchar(500)"`
	DeliveryDate           time.Time `json:"delivery_date" query:"delivery_date" gorm:"type:date"`
	DiscountVehicle        float64   `json:"discount_vehicle" query:"discount_vehicle" gorm:"type:decimal(14,2)"`
	DiscountDown           float64   `json:"discount_down" query:"discount_down" gorm:"type:decimal(14,2)"`
	DiscountAccessory      float64   `json:"discount_accessory" query:"discount_accessory" gorm:"type:decimal(14,2)"`
	DepositRedLicensePlate float64   `json:"deposit_red_license_plate" query:"deposit_red_license_plate" gorm:"type:decimal(14,2)"`
	DepositRegistationFee  float64   `json:"deposit_registation_fee" query:"deposit_registation_fee" gorm:"type:decimal(14,2)"`

	PoNo string `json:"po_no" query:"po_no" gorm:"type:varchar(30)"`

	LeasingStatusID    string    `json:"leasing_status_id" query:"leasing_status_id" gorm:"type:varchar(10)"`
	RequestLeasingDate time.Time `json:"request_leasing_date" query:"request_leasing_date" gorm:"type:date"`
	ApproveLeasingDate time.Time `json:"approve_leasing_date" query:"approve_leasing_date" gorm:"type:date"`

	model.GormRowOrder
}
