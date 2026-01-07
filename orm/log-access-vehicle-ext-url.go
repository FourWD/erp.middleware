package orm

import (
	"github.com/FourWD/middleware/model"
)

type LogAccessVehicleExternalUrl struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36)"`

	model.GormModel

	VehicleExtUrlID string `json:"vehicle_ext_url_id" query:"vehicle_ext_url_id" gorm:"type:varchar(36)"`
	ClientDetail    string `json:"client_detail" query:"client_detail" gorm:"type:text"`
}
