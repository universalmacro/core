package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/universalmacro/common/auth"
	"github.com/universalmacro/common/config"
	"github.com/universalmacro/common/dao"
	"github.com/universalmacro/common/fault"
	"github.com/universalmacro/common/singleton"
	"github.com/universalmacro/common/snowflake"
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

var adminService = singleton.SingletonFactory(newAdminService, singleton.Eager)

func GetAdminService() *AdminService {
	return adminService.Get()
}

func newAdminService() *AdminService {
	return &AdminService{adminRepository: repositories.GetAdminRepository()}
}

type AdminService struct {
	adminRepository *repositories.AdminRepository
}

var ErrPasswordNotMatch = errors.New("password not match")

var sessionIdGenerator = snowflake.NewIdGenertor(0)

func (a *AdminService) CreateSession(account, password string) (string, error) {
	admin, _ := a.adminRepository.FindOne("account = ?", account)
	if admin == nil {
		return "", fault.ErrNotFound
	}
	if !admin.PasswordMatching(password) {
		return "", ErrPasswordNotMatch
	}
	expired := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := Claims{ID: sessionIdGenerator.String(), AdminId: utils.UintToString(admin.ID), StandardClaims: jwt.StandardClaims{ExpiresAt: expired}}
	return auth.SignJwt(claims)
}

var ErrAccountExist = errors.New("account exist")
var ErrCanNotCreateRoot = errors.New("can not create root")
var ErrRoleNotExist = errors.New("role not exist")

func (s *AdminService) CreateAdmin(account, password, role string) (*models.Admin, error) {
	if role == "" {
		role = "ADMIN"
	}
	if role == "ROOT" {
		return nil, ErrCanNotCreateRoot
	}
	if role != "ADMIN" {
		return nil, ErrRoleNotExist
	}
	admin := &entities.Admin{
		Account: account,
		Role:    role,
	}
	admin.SetPassword(password)
	admin, ctx := s.adminRepository.Create(admin)
	if ctx.RowsAffected == 0 {
		return nil, ErrAccountExist
	}
	return models.NewAdmin(admin), nil
}

func (s *AdminService) GetAdminById(id uint) *models.Admin {
	admin, _ := s.adminRepository.GetById(id)
	if admin == nil {
		return nil
	}
	return models.NewAdmin(admin)
}

func (s *AdminService) VerifyToken(token string) (*models.Admin, error) {
	claims, err := auth.VerifyJwt(token)
	if err != nil {
		return nil, err
	}
	admin := s.GetAdminById(utils.StringToUint(claims["adminId"].(string)))
	if admin == nil {
		return nil, fault.ErrNotFound
	}
	return admin, nil
}

func (s *AdminService) ListAdmin(index, limit int64) dao.List[models.Admin] {
	if limit == 0 {
		limit = 1
	}
	adminList, _ := s.adminRepository.Pagination(index, limit)
	var admins []models.Admin
	for index := range adminList.Items {
		admins = append(admins, *models.NewAdmin(&adminList.Items[index]))
	}
	return dao.List[models.Admin]{Items: admins, Pagination: adminList.Pagination}
}

func (s *AdminService) GetAdmin(id uint) *models.Admin {
	admin, _ := s.adminRepository.GetById(id)
	if admin == nil {
		return nil
	}
	return models.NewAdmin(admin)
}

func (s *AdminService) UpdatePassword(id uint, password string) *models.Admin {
	admin := s.GetAdmin(id)
	if admin == nil {
		return admin
	}
	admin.SetPassword(password)
	s.adminRepository.Update(admin.Entity())
	return admin
}

func (s *AdminService) UpdateSelfPassword(id, oldPassword, password string) error {
	admin := s.GetAdminById(utils.StringToUint(id))
	if admin == nil {
		return fault.ErrBadRequest
	}
	if !admin.PasswordMatching(oldPassword) {
		return ErrPasswordNotMatch
	}
	admin.SetPassword(password)
	s.adminRepository.Update(admin.Entity())
	return nil
}

var ErrCanNotDeleteRoot = errors.New("can not delete root")

func (s *AdminService) DeleteAdmin(id uint) error {
	admin := s.GetAdmin(id)
	if admin.Role() == "ROOT" {
		return ErrCanNotDeleteRoot
	}
	s.adminRepository.Delete(admin.Entity())
	return nil
}
