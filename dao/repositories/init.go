package repositories

import (
	"github.com/universalmacro/common/dao"
	single "github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/singleton"
)

func init() {
	singleton.GetDBInstance().AutoMigrate(&entities.Admin{}, &entities.Node{}, &entities.NodeConfig{})
}

var GetAdminRepository = single.EagerSingleton(func() *AdminRepository {
	return &AdminRepository{
		dao.NewRepository[entities.Admin](singleton.CreateDBInstance()),
	}
})

type AdminRepository struct {
	*dao.Repository[entities.Admin]
}
