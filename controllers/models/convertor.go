package models

import (
	"github.com/universalmacro/common/utils"
	"github.com/universalmacro/core/services/models"
)

func AdminConvertor(admin models.Admin) Admin {
	return Admin{
		ID:          utils.UintToString(admin.ID()),
		Account:     admin.Account(),
		Role:        admin.Role(),
		PhoneNumber: PhoneNumberConvertor(admin.PhoneNumber()),
	}
}

func PhoneNumberConvertor(phoneNumber *models.PhoneNumber) *PhoneNumber {
	if phoneNumber == nil {
		return nil
	}
	return &PhoneNumber{
		CountryCode: phoneNumber.CountryCode,
		Number:      phoneNumber.Number,
	}
}
