package services

import (
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/repositories"
)

var adminService = singleton.NewSingleton(newAdminService, singleton.Eager)

func GetAdminService() *AdminService {
	return adminService.Get()
}

type AdminService struct {
	adminRepository *repositories.AdminRepository
}

func newAdminService() *AdminService {
	return &AdminService{adminRepository: repositories.GetAccounrRepository()}
}
