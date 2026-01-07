package orm

import (
	"github.com/FourWD/middleware/model"
)

type Tag struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(2);primary_key;"`

	model.GormModel

	Name     string `json:"name" query:"name" gorm:"type:varchar(100)"`
	Color    string `json:"color" query:"color" gorm:"type:varchar(100)"`
	ColorBg  string `json:"color_bg" query:"name" color_bg:"type:varchar(100)"`
	IconPath string `json:"icon_path" query:"icon_path" gorm:"type:varchar(100)"`
	model.GormRowOrder
}
