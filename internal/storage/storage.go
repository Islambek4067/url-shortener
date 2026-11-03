package storage

import "time"

type Link struct {
	Original  string
	Clicks    int
	CreatedAt time.Time
}

type Storage interface {
	Save(code, url string) error
	Get(code string) (Link, bool, error)
	AddClick(code string) error
}
