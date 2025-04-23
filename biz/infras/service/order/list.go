package order

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent"
	order2 "kcers/biz/dal/db/mysql/ent/order"
	"kcers/biz/dal/db/mysql/ent/predicate"
	product2 "kcers/biz/dal/db/mysql/ent/product"
	"kcers/biz/dal/enums"
	"kcers/biz/infras/service/user"
	"kcers/idl_gen/model/order"
	"strconv"
	"time"
)

func (o *Order) Info(id int64) (info *order.OrderInfo, err error) {
	ret, err := o.db.Order.Query().
		Where(order2.IDEQ(id)).
		Limit(1).
		First(o.ctx)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	info = o.entOrderInfo(ret)
	return
}
func (o *Order) List(req *order.ListOrderReq) (resp []*order.OrderInfo, total int, err error) {
	var predicates []predicate.Order
	if req.OrderSn != "" {
		predicates = append(predicates, order2.OrderSnEQ(req.OrderSn))
	}
	lists, err := o.db.Order.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(o.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Order list failed")
		return resp, total, err
	}
	for _, v := range lists {
		resp = append(resp, o.entOrderInfo(v))
	}

	total, _ = o.db.Order.Query().Where(predicates...).Count(o.ctx)
	return
}
func (o *Order) entOrderAmount(v *ent.Order) (info *order.OrderAmount) {
	first, err := v.QueryAmount().First(o.ctx)
	if err != nil || first == nil {
		return nil
	}
	info = &order.OrderAmount{
		ID:        first.ID,
		Total:     first.Total,
		Actual:    first.Actual,
		Residue:   first.Residue,
		Remission: first.Remission,
		OrderId:   first.OrderID,
		CreatedAt: first.CreatedAt.Format(time.DateTime),
		UpdatedAt: first.UpdatedAt.Format(time.DateTime),
	}
	return

}
func (o *Order) entOrderItem(v *ent.Order) (info *order.OrderItem) {
	first, err := v.QueryItem().First(o.ctx)
	if err != nil || first == nil {
		return nil
	}
	product, err := o.db.Product.Query().Where(product2.IDEQ(first.ProductID)).First(o.ctx)
	if err != nil || product == nil {
		return nil
	}

	info = &order.OrderItem{
		ID:                   first.ID,
		ProductId:            first.ProductID,
		RelatedUserProductId: first.RelatedUserProductID,
		OrderId:              first.OrderID,
		Data:                 first.Data.String(),
		Name:                 first.Name,
		CreatedAt:            first.CreatedAt.Format(time.DateTime),
		UpdatedAt:            first.UpdatedAt.Format(time.DateTime),
		Number:               first.Number,
	}
	return

}
func (o *Order) entOrderPay(v *ent.Order) (info []*order.OrderPay) {
	all, err := v.QueryPay().All(o.ctx)
	if err != nil || all == nil {
		return nil
	}
	for _, first := range all {

		info = append(info, &order.OrderPay{
			ID:        first.ID,
			Pay:       first.Pay,
			Remission: first.Remission,
			OrderId:   first.OrderID,
			PayWay:    first.PayWay,
			PaySn:     first.PaySn,
			PrepayId:  first.PrepayID,
			PayExtra:  string(first.PayExtra),
			CreatedAt: first.CreatedAt.Format(time.DateTime),
			Note:      first.Note,
			UpdatedAt: first.UpdatedAt.Format(time.DateTime),
			PayAt:     first.PayAt.Format(time.DateTime),
		})
	}

	return

}
func (o *Order) entOrderSales(v *ent.Order) (info []*order.OrderSales) {
	all, _ := v.QuerySales().All(o.ctx)
	for _, first := range all {
		u, err := user.NewUser(o.ctx, o.c).Info(first.SalesID)
		if err != nil || u == nil {
			continue
		}

		info = append(info, &order.OrderSales{
			ID:        first.ID,
			SalesId:   first.SalesID,
			Ratio:     first.Ratio,
			OrderId:   first.OrderID,
			SalesName: u.Name,
		})

	}
	return
}
func (o *Order) entOrderInfo(v *ent.Order) (info *order.OrderInfo) {
	vf, _ := v.QueryOrderVenues().First(o.ctx)
	m, _ := v.QueryOrderMembers().First(o.ctx)
	StatusName := enums.ReturnOrderStatusValues(v.Status)

	info = &order.OrderInfo{
		ID:           v.ID,
		OrderSn:      v.OrderSn,
		Status:       v.Status,
		Source:       v.Source,
		Device:       v.Device,
		Nature:       strconv.FormatInt(v.Nature, 10),
		VenueId:      v.VenueID,
		MemberId:     v.MemberID,
		CreatedId:    v.CreatedID,
		CompletionAt: v.CompletionAt.Format(time.DateTime),
		CreatedAt:    v.CreatedAt.Format(time.DateTime),
		UpdatedAt:    v.UpdatedAt.Format(time.DateTime),
		MemberName:   m.Name,
		MemberMobile: m.Mobile,
		OrderAmount:  o.entOrderAmount(v),
		OrderItem:    o.entOrderItem(v),
		OrderPay:     o.entOrderPay(v),
		OrderSales:   o.entOrderSales(v),
		StatusName:   StatusName,
		SourceName:   v.Source,
		VenueName:    vf.Name,
	}

	u, err := user.NewUser(o.ctx, o.c).Info(v.CreatedID)
	if err != nil && u != nil {
		info.CreatedName = u.Name
	}

	return
}
