package services

import (
	"errors"

	"github.com/universalmacro/common/auth"
	"github.com/universalmacro/common/config"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/common/utils"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/dao/repositories"
	"github.com/universalmacro/core/services/models"
)

// Create Root account
func init() {
	adminRepository := repositories.GetAdminRepository()
	account := config.GetString("init.admin.account")
	root, _ := adminRepository.FindOne("account = ?", account)
	if root == nil {
		root = &entities.Admin{
			Account: account,
			Role:    "ROOT",
		}
		root.SetPassword(config.GetString("init.admin.password"))
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

var ErrorPasswordNotMatch = errors.New("password not match")

func (a *AdminService) CreateSession(account, password string) (string, error) {
	admin, _ := a.adminRepository.FindOne("account = ?", account)
	if admin == nil {
		return "", fault.ErrNotFound
	}
	if !admin.PasswordMatching(password) {
		return "", ErrorPasswordNotMatch
	}
	claims := Claims{ID: utils.UintToString(admin.ID)}
	return auth.SignJwt(claims)
}

func (a *AdminService) GetAdminById(id uint) *models.Admin {
	admin, _ := a.adminRepository.GetById(id)
	if admin == nil {
		return nil
	}
	return models.NewAdmin(admin)
}

func (a *AdminService) VerifyToken(token string) (*models.Admin, error) {
	claims, err := auth.VerifyJwt(token)
	if err != nil {
		return nil, err
	}
	admin := a.GetAdminById(utils.StringToUint(claims["id"].(string)))
	if admin == nil {
		return nil, fault.ErrNotFound
	}
	return admin, nil
}
