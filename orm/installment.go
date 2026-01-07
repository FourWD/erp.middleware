package orm

import (
	"github.com/FourWD/middleware/model"
)

type Installment struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36); primary_key"`
	model.GormModel

	Name  string `json:"name" query:"name" gorm:"type:varchar(256)"`
	Value int    `json:"value" query:"value" gorm:"type:int(3)"`
	model.GormRowOrder
}
