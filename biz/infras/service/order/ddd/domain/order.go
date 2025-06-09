package domain

import (
	"fmt"
	"sync"
)

type OrderSn string

type Order struct {
	OrderSn     OrderSn
	Items       []OrderItem
	TotalAmount float64
	Status      OrderStatus
	Version     int          // 乐观锁版本号
	Events      []Event      // 事件列表
	mu          sync.RWMutex // 读写锁
}

// OrderItem 订单商品项
type OrderItem struct {
	ProductId int64
	Quantity  int
	Price     float64
}

func CreateOrder(orderSn OrderSn, items []OrderItem) (order *Order, err error) {
	return
}
func (o *Order) Pay(payMethod string) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.Status != OrderCreated {
		return fmt.Errorf("订单状态错误，当前状态：%s ", o.OrderSn)
	}

	if !o.Status.CanTransitionTo(OrderCancelled) {
		return fmt.Errorf("订单状态不允许从 %s 转换到 %s ", o.Status, OrderCancelled)
	}

	return
}
