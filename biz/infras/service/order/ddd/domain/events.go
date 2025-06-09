package domain

import "time"

type Event interface {
	GetEventId() string
	SetOrderSn() string
	GetCreatedAt() time.Time
	SetEventType() string
}
