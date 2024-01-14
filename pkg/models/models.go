package models

import (
	"errors"
	"time"
)

var ErrorMessage = errors.New("models: no matching record found")

type News struct {
	ID       int
	Title    string
	Content  string
	Date     time.Time
	Category string
}
