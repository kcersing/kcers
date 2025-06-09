package domain

import "sync"

// EventHandler 事件处理器接口
type EventHandler interface {
	Handle(event interface{}) error
}

// EventDispatcher 事件
type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}
