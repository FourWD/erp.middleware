package orm

import "github.com/FourWD/middleware/model"

type RoleMenu struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`

	model.GormModel

	RoleID     string `json:"role_id" query:"role_id" gorm:"type:varchar(36);"`
	RoleMenuID string `json:"role_menu_id" query:"role_menu_id" gorm:"type:varchar(36);"`

	Name     string `json:"name" query:"name" gorm:"type:varchar(100);"`
	IsActive bool   `json:"is_active" query:"is_active" gorm:"type:bool;"`
	Path     string `json:"path" query:"path" gorm:"type:varchar(100);unique"`
	Icon     string `json:"icon" query:"icon" gorm:"type:varchar(256);"`

	model.GormRowOrder
}
