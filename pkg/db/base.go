package db

import "time"

// Base base
type Base struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement;comment:主键"`
	CreatedAt time.Time `gorm:"column:created_at;comment:创建时间"`
}
