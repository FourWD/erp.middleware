package orm

import (
	"github.com/FourWD/middleware/model"
)

type PoDetailGift struct { // gift

	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`

	PoDetailID string `json:"po_detail_id" query:"po_detail_id" gorm:"type:varchar(36)"`

	model.GormModel

	Name   string  `json:"name" query:"name" gorm:"type:varchar(100)"`
	Price  float64 `json:"price" query:"price" gorm:"type:decimal(12,4)"`
	Remark string  `json:"remark" query:"remark" gorm:"type:varchar(500)"`
	model.GormRowOrder
}
