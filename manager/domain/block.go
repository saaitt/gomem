package domain

import "time"

type Block struct {
	Name      string
	Content   map[string]interface{}
	TTL       time.Duration
	CreatedAt time.Time
}
