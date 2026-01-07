package orm

import "github.com/FourWD/middleware/model"

type Users struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);uniqueIndex:idx_id;primary_key"`
	model.GormModel

	Code         string `json:"code" query:"code" gorm:"type:varchar(10)"`
	UserTypeID   string `json:"user_type_id" query:"user_type_id" gorm:"type:varchar(2)"`
	Firstname    string `json:"firstname" query:"firstname" gorm:"type:varchar(100)"`
	Lastname     string `json:"lastname" query:"lastname" gorm:"type:varchar(100)"`
	Username     string `json:"username" query:"username" gorm:"type:varchar(20)"`
	Password     string `json:"password" query:"password" gorm:"text"`
	FileAvatarID string `json:"file_avartar_id" query:"file_avartar_id" gorm:"type:varchar(36)"`
	Mobile       string `json:"mobile" query:"mobile" gorm:"type:varchar(10)"`
	Email        string `json:"email" query:"email" gorm:"type:varchar(50)"`
	Facebook     string `json:"facebook" query:"facebook" gorm:"type:varchar(50)"`
	Line         string `json:"line" query:"line" gorm:"type:varchar(20)"`
	Token        string `json:"token" query:"token" gorm:"type:text"`
}
