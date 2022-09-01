package db

import "time"

type Base struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
