package models

import "github.com/universalmacro/core/dao/entities"

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

func (a *Admin) PasswordMatching(password string) bool {
	return a.entity.PasswordMatching(password)
}

func (a *Admin) SetPassword(password string) (string, []byte) {
	return a.entity.SetPassword(password)
}

func (a *Admin) Role() string {
	return a.entity.Role
}

type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}
