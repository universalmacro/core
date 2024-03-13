package singleton

import (
	"github.com/universalmacro/common/config"
	"github.com/universalmacro/common/dao"

	single "github.com/universalmacro/common/singleton"
	"gorm.io/gorm"
)

var GetDBInstance = single.EagerSingleton(CreateDBInstance)

func CreateDBInstance() *gorm.DB {
	db, err := dao.NewConnection(
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.database"),
	)
	if err != nil {
		panic(err)
	}
	return db
}
