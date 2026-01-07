package orm

import "github.com/FourWD/middleware/model"

type OptionGift struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel
	Name   string  `json:"name" query:"name" gorm:"type:varchar(100)"`
	Price  float64 `json:"price" query:"price" gorm:"type:decimal(12,4)"`
	Remark string  `json:"remark" query:"remark" gorm:"type:varchar(500)"`
	model.GormRowOrder
}
