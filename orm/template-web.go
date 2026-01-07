package orm

import "github.com/FourWD/middleware/model"

type TemplateWeb struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel
	SourceID       string `json:"source_id" query:"source_id" gorm:"type:varchar(8)"`
	TemplateID     string `json:"template_id" query:"template_id" gorm:"type:varchar(2)"`
	Name           string `json:"name" query:"name" gorm:"type:varchar(255)"`
	PrimaryColor   string `json:"primary_color" query:"primary_color" gorm:"type:varchar(10)"`
	SecondaryColor string `json:"secondary_color" query:"secondary_color" gorm:"type:varchar(10)"`
	Logo           string `json:"logo" query:"logo" gorm:"type:varchar(500)"`
	IsActive       int    `json:"is_active" query:"is_active" gorm:"type:tinyint(1) ; column:is_active"`
	CardStyle      string `json:"card_style" query:"card_style" gorm:"type:varchar(3)"`
}
