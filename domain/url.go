package domain

import "time"

type URL struct {
	ID         int64 `gorm:"primary_key"`
	OriginURL  string
	ShortPath  string
	CreateTime time.Time
}

type PostURL struct {
	OriginURL string `json:"oritin_url"`
}
