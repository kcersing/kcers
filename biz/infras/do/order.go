package do

import (
	"kcers/biz/dal/db/mysql/ent"
	"kcers/idl_gen/model/order"
)

type Order interface {
	Buy(req *order.BuyReq) (string, error)
	Update(req *order.UpdateOrderReq) error
	List(req *order.ListOrderReq) (resp []**order.OrderItem, total int, err error)
	OrderListExport(req *order.ListOrderReq) (string, error)
	OrderAllCount(req *order.OrderAllCountReq) (resp []*order.OrderCountInfo, actual float64, total int, err error)

	UpdateStatus(id int64, status int64) error
	Info(id int64) (info *order.OrderInfo, err error)
	GetBySnOrder(sn string) (info *order.OrderInfo, err error)

	CompletedOrder(orderEnt *ent.Order) (err error)
	RefundOrder(id int64) (err error)
	Pay(req PayOrder) (err error)
}
type PayOrder struct {
	Sn            string
	Fee           int64
	TransactionId string
	Transaction   []byte
	OrderEnt      *ent.Order
}

// OrderStatus 订单状态
type OrderStatus int

const (
	OrderStatusUnpaid = iota
	OrderStatusPartial
	OrderStatusCompleted
	OrderStatusPending
	OrderStatusRefunded
	OrderStatusReversed
)

var OrderStatusNames = map[OrderStatus]string{
	OrderStatusUnpaid:    "未付款",
	OrderStatusPartial:   "部分付款",
	OrderStatusCompleted: "已完成",
	OrderStatusPending:   "退款中",
	OrderStatusRefunded:  "已退款",
	OrderStatusReversed:  "已取消",
}

func (s OrderStatus) String() string {
	return OrderStatusNames[s]
}
