package services

import (
	"github.com/universalmacro/common/config"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/dao/repositories"
)

// Create Root account
func init() {
	adminRepository := repositories.GetAdminRepository()
	account := config.GetString("init.account")
	root, _ := adminRepository.FindOne("account = ?", account)
	if root == nil {
		root = &entities.Admin{
			Role: "ROOT",
		}
		root.SetPassword(config.GetString("init.password"))
		adminRepository.Create(root)
	}
}

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
