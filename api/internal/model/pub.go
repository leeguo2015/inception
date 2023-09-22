package model

import (
	"database/sql"

	"time"
)

type Base struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT"`
}

type BaseTime struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	DeletedAt sql.NullTime `json:"-"`
}
