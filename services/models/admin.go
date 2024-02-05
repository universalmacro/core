package models

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/universalmacro/core/dao/entities"
	"github.com/universalmacro/core/dao/repositories"
)

func NewAdmin(entity *entities.Admin) *Admin {
	return &Admin{entity: entity}
}

type Admin struct {
	entity *entities.Admin
}

func (a *Admin) ID() uint {
	return a.entity.ID
}

func (a *Admin) Account() string {
	return a.entity.Account
}

func (a *Admin) PhoneNumber() *PhoneNumber {
	if a.entity.PhoneNumber == nil {
		return nil
	}
	return &PhoneNumber{
		CountryCode: a.entity.CountryCode,
		Number:      a.entity.Number,
	}
}

func (a *Admin) Entity() *entities.Admin {
	return a.entity
}

func (a *Admin) PasswordMatching(password string) bool {
	return a.entity.PasswordMatching(password)
}

func (a *Admin) SetPassword(password string) (string, []byte) {
	return a.entity.SetPassword(password)
}

func (a *Admin) Role() string {
	return a.entity.Role
}

func (a *Admin) CreatedAt() time.Time {
	return a.entity.CreatedAt
}
func (a *Admin) UpdatedAt() time.Time {
	return a.entity.UpdatedAt
}

func (a *Admin) UpdatePassword(password string) {
	a.entity.SetPassword(password)
}

func (a *Admin) UpdateTotp(url, code string) bool {
	key, err := otp.NewKeyFromURL(url)
	if err != nil {
		return false
	}
	if !totp.Validate(code, key.Secret()) {
		return false
	}
	a.entity.Totp = url
	a.Submit()
	return true
}

func (a *Admin) Submit() {
	repositories.GetAdminRepository().Update(a.entity)
}

type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}
