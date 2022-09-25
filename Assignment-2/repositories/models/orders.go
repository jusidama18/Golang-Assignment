package models

import "time"

type Orders struct {
	OrderID      int
	CustomerName string
	OrderedAt    time.Time
	Items        []Items
}
