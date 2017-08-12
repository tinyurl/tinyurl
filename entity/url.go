package entity

import "time"

type URL struct {
	ID         int64 `gorm:"primary_key"`
	OriginUrl  string
	ShortPath  string
	CreateTime time.Time
}
