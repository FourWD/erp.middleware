package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type InterestMonth struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`

	model.GormModel

	FinanceID               string    `json:"finance_id" query:"finance_id" gorm:"type:varchar(36)"`
	VehicleTypeID           string    `json:"vehicle_type_id" query:"vehicle_type_id" gorm:"type:varchar(36)"` // รย3
	StartYearManufacturing  int       `json:"start_year_manufacturing" query:"start_year_manufacturing" gorm:"type:int(4)"`
	EndYearManufacfacturing int       `json:"end_year_manufacturing" query:"end_year_manufacturing" gorm:"type:int(4)"`
	InterestMonth           int       `json:"interest_month" query:"interest_month" gorm:"type:int(2)"`
	StartDate               time.Time `json:"start_date" query:"start_date" `
	EndDate                 time.Time `json:"end_date" query:"end_date" `
	Interest                float64   `json:"interest" query:"interest" gorm:"type:float(4,2)"`
	IsActive                int       `json:"is_active" query:"is_active" gorm:"type:tinyint(1) ; column:is_active"`
	StartYear               string    `json:"start_year" query:"start_year" `
	EndYear               string    `json:"end_year" query:"end_year" `
}
