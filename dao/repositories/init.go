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

var adminRepository = single.NewSingleton[AdminRepository](func() *AdminRepository {
	return &AdminRepository{
		dao.NewRepository[entities.Admin](singleton.CreateDBInstance()),
	}
}, single.Eager)

func GetAdminRepository() *AdminRepository {
	return adminRepository.Get()
}

type AdminRepository struct {
	*dao.Repository[entities.Admin]
}
