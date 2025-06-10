package domain

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"sync"
)

// EventHandler 事件处理器接口
type EventHandler interface {
	Handle(event Event) error
}

// EventDispatcher
type EventDispatcher struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

// NewEventDispatcher
func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

// RegisterHandler 注册事件处理器
func (d *EventDispatcher) RegisterHandler(eventType string, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

// Dispatch 分发事件
func (d *EventDispatcher) Dispatch(event Event) {
	d.mu.RLock() // 读取锁
	defer d.mu.RUnlock()

	eventType := event.SetEventType()
	if handlers, ok := d.handlers[eventType]; ok {
		var wg sync.WaitGroup
		for _, handler := range handlers {
			wg.Add(1)
			go func(h EventHandler) {
				defer wg.Done()

				if err := h.Handle(event); err != nil {
					// 处理错误
					hlog.Errorf("事件处理失败：%v", err)
				}

			}(handler)
		}

	}

}

// InventoryService 库存服务接口
type InventoryService interface {
	Reserve() error          //预留库存
	Release() error          //扣除库存
	RestoreForRefund() error //恢复库存
}
type PayService interface {
	Refund(sn OrderSn, amount float64, reason string, createdId int64) error
}

type NotificationServer interface {
	SendCreatedNotification(sn OrderSn) error
	SendPaidNotification(sn OrderSn) error
}
