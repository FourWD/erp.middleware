package orm

type CarAuctionOma struct {
	ID                string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	VehicleBrandID    string `json:"vehicle_brand_id" query:"vehicle_brand_id" gorm:"type:varchar(36)"`
	VehicleModelID    string `json:"vehicle_model_id" query:"vehicle_model_id" gorm:"type:varchar(36)"`
	VehicleSubModelID string `json:"vehicle_sub_model_id" query:"vehicle_sub_model_id" gorm:"type:varchar(36)"`
	Year              int    `json:"year" query:"year" gorm:"type:int"`
	Gear              string `json:"gear" query:"gear" gorm:"type:varchar(10)"`
	CC                string `json:"cc" query:"cc" gorm:"type:varchar(10)"`
	MNumber           string `json:"m_number" query:"m_number" gorm:"type:varchar(50)"`
	Color             string `json:"color" query:"color" gorm:"type:varchar(50)"`
}

func (CarAuctionOma) TableName() string {
	return "car_auction_oma"
}
