package orm

import (
	"github.com/FourWD/middleware/model"
)

// midOrm "github.com/FourWD/middleware/orm"

type PromotionMaster struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`

	model.GormModel

	Name string `json:"name" query:"name" gorm:"type:varchar(255)"`

	model.GormRowOrder
}

// Title            string    `json:"title" query:"title" gorm:"type:varchar(500)"`
// Description      string    `json:"description" query:"description" gorm:"type:text" `
// Remark           string    `json:"remark" query:"remark" gorm:"type:varchar(255)"`
// BannerPath       string    `json:"banner_path" query:"banner_path" `
// BannerMobilePath string    `json:"banner_mobile_path" query:"banner_mobile_path" `
// ShowDate         time.Time `json:"show_date" query:"show_date" `
// StartDate        time.Time `json:"start_date" query:"start_date" `
// EndDate          time.Time `json:"end_date" query:"end_date" `
// IsHome           bool      `json:"is_home"`
