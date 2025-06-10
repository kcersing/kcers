package domain

import "time"

type Event interface {
	GetEventId() string
	SetOrderSn() OrderSn
	GetCreatedAt() time.Time
	SetEventType() string
}
