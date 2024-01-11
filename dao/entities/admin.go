package entities

import (
	"github.com/universalmacro/common/auth"
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/snowflake"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Account string `gorm:"type:varchar(64);uniqueIndex"`
	dao.PhoneNumber
	auth.Password
	Role string `gorm:"type:varchar(20)"`
}

var adminIdGenerator = snowflake.NewIdGenertor(0)

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = adminIdGenerator.Uint()
	return err
}
