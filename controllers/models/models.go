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

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required"`
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
	// Config      NodeConfig `json:"config"`
}

type NodeConfig struct {
	DBConfig    *DBConfig    `json:"dbConfig"`
	RedisConfig *RedisConfig `json:"redisConfig"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}
