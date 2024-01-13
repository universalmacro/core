package repositories

import (
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/entities"
)

func init() {
	dao.GetDBInstance().AutoMigrate(&entities.Admin{}, &entities.Node{}, &entities.NodeConfig{})
}

var adminRepository = singleton.NewSingleton[AdminRepository](func() *AdminRepository {
	return &AdminRepository{
		dao.NewRepository[entities.Admin](),
	}
}, singleton.Eager)

func GetAdminRepository() *AdminRepository {
	return adminRepository.Get()
}

type AdminRepository struct {
	*dao.Repository[entities.Admin]
}
