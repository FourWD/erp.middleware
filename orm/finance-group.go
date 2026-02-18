package orm

import (
	"github.com/FourWD/middleware/model"
)

type FinanceGroup struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`
	model.GormModel
	Name string `json:"name" query:"name" gorm:"type:varchar(256)"`
}
