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

type Admin struct {
	ID          string       `json:"id"`
	Account     string       `json:"account"`
	PhoneNumber *PhoneNumber `json:"phoneNumber"`
	Role        string       `json:"role"`
}

type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}

type CreateNodeRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
