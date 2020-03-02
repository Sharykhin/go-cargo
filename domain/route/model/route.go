package model

import "time"

type (
	OrderID string
	Route struct {
		OrderID OrderID
		DateArrival time.Time
	}
)