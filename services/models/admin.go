package models

import "github.com/universalmacro/core/dao/entities"

type Admin struct {
	entity *entities.Admin
}

func NewAdmin(entity *entities.Admin) *Admin {
	return &Admin{entity: entity}
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
