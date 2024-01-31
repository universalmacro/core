package entities

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/universalmacro/common/snowflake"
	"github.com/universalmacro/common/utils"
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	SecurityKey string `gorm:"type:varchar(64);uniqueIndex"`
	Name        string `gorm:"type:varchar(64)"`
	Description string `gorm:"type:varchar(200)"`
}

var nodeIdGenerator = snowflake.NewIdGenertor(0)

func (a *Node) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = nodeIdGenerator.Uint()
	return err
}

type NodeConfig struct {
	gorm.Model
	NodeID             uint
	SecretKey          string              `gorm:"type:varchar(64)"`
	FrontendDomain     string              `json:"frontendDomain" gorm:"type:varchar(64);uniqueIndex"`
	Api                *ApiConfig          `gorm:"type:json"`
	Server             *ServerConfig       `gorm:"type:json"`
	Database           *DBConfig           `gorm:"type:json"`
	Redis              *RedisConfig        `gorm:"type:json"`
	TencentCloudConfig *TencentCloudConfig `json:"type:json"`
}

type ApiConfig struct {
	MerchantUrl string `json:"merchantUrl" gorm:"type:varchar(256)"`
}

func (j *ApiConfig) Scan(value any) error {
	return utils.ScanJson(value, j)
}

func (j ApiConfig) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type ServerConfig struct {
	Port      string `json:"port" gorm:"type:varchar(64)"`
	JwtSecret string `json:"jwtSecret" gorm:"type:varchar(64)"`
}

func (j *ServerConfig) Scan(value any) error {
	return utils.ScanJson(value, j)
}

func (j ServerConfig) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type RedisConfig struct {
	Host     string `json:"host" gorm:"type:varchar(64)"`
	Port     string `json:"port" gorm:"type:varchar(64)"`
	Password string `json:"password" gorm:"type:varchar(64)"`
}

func (j *RedisConfig) Scan(value any) error {
	return utils.ScanJson(value, j)
}

func (j RedisConfig) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type DBConfig struct {
	Host     string `json:"host" gorm:"type:varchar(64)"`
	Port     string `json:"port" gorm:"type:varchar(64)"`
	Username string `json:"username" gorm:"type:varchar(64)"`
	Password string `json:"password" gorm:"type:varchar(64)"`
	Database string `json:"database" gorm:"type:varchar(64)"`
	Type     string `json:"type" gorm:"type:varchar(64)"`
}

func (j *DBConfig) Scan(value any) error {
	return utils.ScanJson(value, j)
}

func (j DBConfig) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type TencentCloudConfig struct {
	SecretId  string `json:"secretId" gorm:"type:varchar(128)"`
	SecretKey string `json:"secretKey" gorm:"type:varchar(128)"`
}

func (j *TencentCloudConfig) Scan(value any) error {
	return utils.ScanJson(value, j)
}

func (j TencentCloudConfig) Value() (driver.Value, error) {
	return json.Marshal(j)
}
