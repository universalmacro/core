package models

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
