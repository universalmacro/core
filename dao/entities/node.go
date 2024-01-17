package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/universalmacro/common/snowflake"
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	SecurityKey string `gorm:"type:varchar(64);uniqueIndex"`
	Name        string `gorm:"type:varchar(64)"`
	Description string `gorm:"type:varchar(200)"`
	Config      NodeConfig
}

var nodeIdGenerator = snowflake.NewIdGenertor(0)

func (a *Node) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = nodeIdGenerator.Uint()
	return err
}

type NodeConfig struct {
	gorm.Model
	NodeID    uint
	SecretKey string       `gorm:"type:varchar(64)"`
	Database  *DBConfig    `gorm:"type:json"`
	Redis     *RedisConfig `gorm:"type:json"`
}

type RedisConfig struct {
	Host     string `gorm:"type:varchar(64)"`
	Port     string `gorm:"type:varchar(64)"`
	Password string `gorm:"type:varchar(64)"`
}

func (j *RedisConfig) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := RedisConfig{}
	err := json.Unmarshal(bytes, &result)
	*j = RedisConfig(result)
	return err
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
}

func (j *DBConfig) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := DBConfig{}
	err := json.Unmarshal(bytes, &result)
	*j = DBConfig(result)
	return err
}

func (j DBConfig) Value() (driver.Value, error) {
	return json.Marshal(j)
}
