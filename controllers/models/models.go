package models

import "github.com/universalmacro/core/dao/entities"

type CreateSessionRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Session struct {
	Token string `json:"token"`
}

type CreateAdminRequest struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UpdatePasswordRequest struct {
	OldPassword *string `json:"oldPassword" binding:"required"`
	Password    string  `json:"password" binding:"required"`
}

type Admin struct {
	ID          string       `json:"id"`
	Account     string       `json:"account"`
	PhoneNumber *PhoneNumber `json:"phoneNumber"`
	Role        string       `json:"role"`
	CreatedAt   int64        `json:"createdAt"`
	UpdatedAt   int64        `json:"updatedAt"`
}

type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}

type CreateNodeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Node struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SecurityKey string `json:"securityKey"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

type NodeConfig struct {
	FrontendDomain *string                `json:"frontendDomain"`
	Api            *entities.ApiConfig    `json:"api"`
	Server         *entities.ServerConfig `json:"server"`
	Database       *entities.DBConfig     `json:"database"`
	Redis          *entities.RedisConfig  `json:"redis"`
}

type CreateMerchantRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type Merchant struct {
	ID        string `json:"id"`
	Account   string `json:"account"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
