package entity

import "time"

type Order struct {
	Id       string
	Customer struct {
		Id       string
		Username string
	}
	Total     float64
	CreatedAt time.Time
	UpdatedAt *time.Time
}
