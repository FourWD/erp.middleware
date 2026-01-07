package orm

type VehicleStatus struct {
	ID       string `json:"id" query:"id" gorm:"type:varchar(2);primary_key"`
	Name     string `json:"name" query:"name" gorm:"not null;type:varchar(50)"`
	Priority int    `json:"priority" query:"priority" gorm:"not null;type:int"`
}
