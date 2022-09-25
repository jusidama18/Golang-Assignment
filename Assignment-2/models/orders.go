package models

import "time"

type Orders struct {
	OrderID      int
	CustomerName string
	OrderAt      time.Time
	Items        []Items
}
