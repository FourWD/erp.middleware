package orm

type Template struct {
	ID       string `json:"id" query:"id" gorm:"type:varchar(2);primary_key"`
	Name     string `json:"name" query:"name" gorm:"type:varchar(100)"`
	IsActive int    `json:"is_active" query:"is_active" gorm:"type:tinyint(1) ; column:is_active"`
}
