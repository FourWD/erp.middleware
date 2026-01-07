package orm

import (
	midOrm "github.com/FourWD/middleware/orm"
)

type Employee struct {
	midOrm.Employee
	RoleDetail string `json:"role_detail" query:"role_detail" gorm:"type:text"`
	CompanyID  string `json:"company_id" query:"company_id" gorm:"type:varchar(36)"`
	AcceptWEB  bool   `json:"accept_web" query:"accept_web" `
	AcceptBO   bool   `json:"accept_bo" query:"accept_bo" `
	SourceList string `json:"source_list" query:"source_list" gorm:"type:varchar(100)"`
}
