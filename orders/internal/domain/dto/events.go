package dto

import "time"

type SaveOrderEvent struct {
	Order     Order
	Customer  Customer
	CreatedAt time.Time
}
