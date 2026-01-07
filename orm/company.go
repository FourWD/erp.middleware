package orm

import (
	"github.com/FourWD/middleware/model"
)

// midOrm "github.com/FourWD/middleware/orm"

type Company struct { // migration
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`

	model.GormModel

	CompanyName string `json:"company_name" query:"company_name" gorm:"type:varchar(500)"`
	// BranchListCode string `json:"branch_list_code" query:"branch_list_code" gorm:"type:text"`
	Logo string `json:"logo" query:"logo" gorm:"type:varchar(255)"`

	model.GormRowOrder
}
