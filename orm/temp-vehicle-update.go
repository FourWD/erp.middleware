package orm

type VehicleUpdateTemp struct {
	ChassisNo string  `gorm:"column:chassis_no;type:varchar(256)"`
	CarID     string  `gorm:"column:car_id;type:varchar(256)"`
	Status    string  `gorm:"column:status;type:varchar(256)"`
	Mile      string  `gorm:"column:mile;type:varchar(256)"`
	Price     float64 `gorm:"column:price"`
}
