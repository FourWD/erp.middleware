package orm

import "github.com/FourWD/middleware/model"

// midOrm "github.com/FourWD/middleware/orm"

type ReferralChannelSub struct { //
	ID string `json:"id" query:"id" gorm:"type:varchar(2); primary_key"`
	model.GormModel

	Name              string `json:"name" query:"name" gorm:"type:varchar(256)"` // line face
	ReferralChannelID string `json:"referral_channel_id" query:"referral_channel_id" gorm:"type:varchar(2)"`
	SubCode           string `json:"sub_code" query:"sub_code" gorm:"type:varchar(2)"`
	model.GormRowOrder
}
