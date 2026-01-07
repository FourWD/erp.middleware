package orm

import "github.com/FourWD/middleware/model"

type TemplateDetail struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel
	TemplateID string `json:"template_id" query:"template_id" gorm:"type:varchar(2)"`
	// TemplateSectionID    string `json:"template_section_id" query:"template_section_id" gorm:"type:varchar(2)"`
	// TemplateMasterID     string `json:"template_master_id" query:"template_master_id" gorm:"type:varchar(36)"`
	TemplateResolutionID string `json:"template_resolution_id" query:"template_resolution_id" gorm:"type:varchar(2)"`
	// Value                string `json:"value" query:"value" gorm:"type:text"`

	ContainerWidth int    `json:"container_width" gorm:"column:container_width"`
	Resolution     string `json:"resolution" gorm:"column:resolution"` // "01"=PC, "02"=Mobile, "03"=Tablet

	// Navigation
	NavBgColor    string `json:"nav_bg_color" gorm:"column:nav_bg_color"`
	NavTextColor  string `json:"nav_text_color" gorm:"column:nav_text_color"`
	NavLogoPath   string `json:"nav_logo_path" gorm:"column:nav_logo_path"`
	NavAlign      string `json:"nav_align" gorm:"column:nav_align"`
	NavMargin     int    `json:"nav_margin" gorm:"column:nav_margin"`
	NavPadding    int    `json:"nav_padding" gorm:"column:nav_padding"`
	NavFontSize   int    `json:"nav_font_size" gorm:"column:nav_font_size"`
	NavFontWeight int    `json:"nav_font_weight" gorm:"column:nav_font_weight"`
	NavGap        int    `json:"nav_gap" gorm:"column:nav_gap"`

	// Header
	HeaderWidth     int    `json:"header_width" gorm:"column:header_width"`
	HeaderHeight    int    `json:"header_height" gorm:"column:header_height"`
	HeaderObjectFit string `json:"header_object_fit" gorm:"column:header_object_fit"`

	// Topic
	TopicFontWeight int    `json:"topic_font_weight" gorm:"column:topic_font_weight"`
	TopicFontSize   int    `json:"topic_font_size" gorm:"column:topic_font_size"`
	TopicText       string `json:"topic_text" gorm:"column:topic_text"`

	// Card
	CardAspect            string `json:"card_aspect" gorm:"column:card_aspect"`
	CardHeight            int    `json:"card_height" gorm:"column:card_height"`
	CardWidth             int    `json:"card_width" gorm:"column:card_width"`
	CardGap               int    `json:"card_gap" gorm:"column:card_gap"`
	CardGridCols          int    `json:"card_grid_cols" gorm:"column:card_grid_cols"`
	CardObjectFit         string `json:"card_object_fit" gorm:"column:card_object_fit"`
	CardImageHeight       int    `json:"card_image_height" gorm:"column:card_image_height"`
	CardImageWidth        int    `json:"card_image_width" gorm:"column:card_image_width"`
	CardLayout            string `json:"card_layout" gorm:"column:card_layout"`
	CardMainColor         string `json:"card_main_color" gorm:"column:card_main_color"`
	CardSubColor          string `json:"card_sub_color" gorm:"column:card_sub_color"`
	CardIsHidePricePreVat int    `json:"card_is_hide_price_pre_vat" gorm:"column:card_is_hide_price_pre_vat"`
	CardIsShowPrice       string `json:"card_is_show_price" gorm:"column:card_is_show_price"`
	CardMargin            int    `json:"card_margin" gorm:"column:card_margin"`
	CardPadding           int    `json:"card_padding" gorm:"column:card_padding"`
	CardFontSize          int    `json:"card_font_size" gorm:"column:card_font_size"`
	CardFontWeight        int    `json:"card_font_weight" gorm:"column:card_font_weight"`
	CardRadius            int    `json:"card_radius" gorm:"column:card_radius"`
	CardInsideGrid        int    `json:"card_inside_grid" gorm:"column:card_inside_grid"`
	CardInsideGap         int    `json:"card_inside_gap" gorm:"column:card_inside_gap"`
	CardPriceAlign        string `json:"card_price_align" gorm:"column:card_price_align"`
	CardShadow            string `json:"card_shadow" gorm:"column:card_shadow"`
	CardPriceColor        string `json:"card_price_color" gorm:"column:card_price_color"`
	CardPriceFontSize     int    `json:"card_price_font_size" gorm:"column:card_price_font_size"`
	CardPriceFontWeight   int    `json:"card_price_font_weight" gorm:"column:card_price_font_weight"`

	// Footer
	FooterGrid       int    `json:"footer_grid" gorm:"column:footer_grid"`
	FooterLayout     string `json:"footer_layout" gorm:"column:footer_layout"`
	FooterTextColor  string `json:"footer_text_color" gorm:"column:footer_text_color"`
	FooterBgColor    string `json:"footer_bg_color" gorm:"column:footer_bg_color"`
	FooterMargin     int    `json:"footer_margin" gorm:"column:footer_margin"`
	FooterPadding    int    `json:"footer_padding" gorm:"column:footer_padding"`
	FooterFontSize   int    `json:"footer_font_size" gorm:"column:footer_font_size"`
	FooterFontWeight int    `json:"footer_font_weight" gorm:"column:footer_font_weight"`
	FooterGap        int    `json:"footer_gap" gorm:"column:footer_gap"`

	// Misc
	CSSUnit string `json:"css_unit" gorm:"column:css_unit"`
}
