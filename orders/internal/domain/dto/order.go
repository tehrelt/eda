package dto

import "time"

type SaveOrder struct {
	Username string
	Total    string
}

type Order struct {
	Id    string
	Total float64
}

type Customer struct {
	Id       string
	Username string
}

type OrderSummary struct {
	Username      string
	Total         string
	LastOrderedAt time.Time
}
