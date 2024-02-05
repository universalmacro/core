package entities

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/universalmacro/common/utils"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string
	Description string
}

type Permission struct {
	Domain      string
	Permissions []string
}

func (j *Permission) Scan(value any) error {
	return utils.ScanJson(value, j)
}

func (j Permission) Value() (driver.Value, error) {
	return json.Marshal(j)
}
