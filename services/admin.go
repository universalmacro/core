package services

import (
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/repositories"
)

var adminService = singleton.NewSingleton(newAdminService, singleton.Eager)

func GetAdminService() *AdminService {
	return adminService.Get()
}

func newAdminService() *AdminService {
	return &AdminService{adminRepository: repositories.GetAdminRepository()}
}

type AdminService struct {
	adminRepository *repositories.AdminRepository
}

func init() {
	// adminRepository := repositories.GetAdminRepository()
	// adminRepository.GetById()
}
