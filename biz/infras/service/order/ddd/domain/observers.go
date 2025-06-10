package domain

import (
	"errors"
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
	Reserve(sn OrderSn, items []OrderItem) error //预留库存
	Release(sn OrderSn) error                    //扣除库存
	RestoreForRefund(sn OrderSn) error           //恢复库存
}

// InventoryHandler 库存处理器
type InventoryHandler struct {
	InventoryService InventoryService
}

func (h *InventoryHandler) Handle(event Event) error {
	if h.InventoryService == nil {
		return errors.New("库存服务未初始化")
	}
	switch e := event.(type) {
	case *OrderCreatedEvent:
		// 处理订单创建事件
		return h.InventoryService.Reserve(e.OrderSn, e.Items)
	case *OrderCancelledEvent:
		// 处理订单支付事件
		return h.InventoryService.Release(e.OrderSn)
	case *OrderRefundedEvent:
		// 处理订单退款事件
		return h.InventoryService.RestoreForRefund(e.OrderSn)
	default:
		return nil
	}
}

// PayService 支付服务接口
type PayService interface {
	Refund(sn OrderSn, amount float64, reason string, createdId int64) error
}
type PayHandler struct {
	PayService PayService
}

func (h *PayHandler) Handle(event Event) error {
	if h.PayService == nil {
		return errors.New("支付服务未初始化")
	}
	switch e := event.(type) {
	case *OrderRefundedEvent:
		return h.PayService.Refund(e.OrderSn, e.RefundedAmount, e.Reason, e.CreatedId)
	default:
		return nil
	}

}

// NotificationServer 通知服务接口
type NotificationServer interface {
	SendOrderCreatedNotification(sn OrderSn) error
	SendOrderPaidNotification(sn OrderSn) error
	SendOrderShippedNotification(sn OrderSn) error
	SendOrderCompletedNotification(sn OrderSn) error
	SendOrderCancelledNotification(sn OrderSn) error
	SendOrderRefundNotification(sn OrderSn) error
}
type NotificationHandler struct {
	NotificationServer NotificationServer
}

func (h *NotificationHandler) Handler(event Event) error {
	if h.NotificationServer == nil {
		return errors.New("通知服务未初始化")
	}
	switch e := event.(type) {
	case *OrderCreatedEvent:
		// 处理订单创建事件
		return h.NotificationServer.SendOrderCreatedNotification(e.OrderSn)
	case *OrderPaidEvent:
		// 处理订单支付事件
		return h.NotificationServer.SendOrderPaidNotification(e.OrderSn)
	case *OrderShippedEvent:
		// 处理订单发货事件
		return h.NotificationServer.SendOrderShippedNotification(e.OrderSn)
	case *OrderCompletedEvent:
		// 处理订单完成事件
		return h.NotificationServer.SendOrderCompletedNotification(e.OrderSn)
	case *OrderCancelledEvent:
		// 处理订单取消事件
		return h.NotificationServer.SendOrderCancelledNotification(e.OrderSn)
	case *OrderRefundedEvent:
		// 处理订单退款事件
		return h.NotificationServer.SendOrderRefundNotification(e.OrderSn)
	default:
		return nil
	}
}
