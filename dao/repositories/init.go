package repositories

import (
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/entities"
)

func init() {
	dao.GetDBInstance().AutoMigrate(&entities.Admin{})
}

var adminRepository = singleton.NewSingleton[AdminRepository](func() *AdminRepository {
	return &AdminRepository{
		dao.NewRepository[entities.Admin](),
	}
}, singleton.Eager)

func GetAccounrRepository() *AdminRepository {
	return adminRepository.Get()
}

type AdminRepository struct {
	*dao.Repository[entities.Admin]
}
