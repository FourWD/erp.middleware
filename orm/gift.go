package orm

import (
	"github.com/FourWD/middleware/model"
	midOrm "github.com/FourWD/middleware/orm"
)

type Gift struct {
	midOrm.Gift
	model.GormRowOrder
}
