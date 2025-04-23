package order

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"kcers/biz/dal/cache"
	"kcers/biz/dal/config"
	db "kcers/biz/dal/db/mysql"
	"kcers/biz/dal/db/mysql/ent"
	order2 "kcers/biz/dal/db/mysql/ent/order"
	"kcers/biz/infras/do"
	"kcers/idl_gen/model/order"
)

type Order struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (o *Order) Update(req *order.OrderInfo) error {
	_, err := o.db.Order.Update().
		Where(order2.IDEQ(req.ID)).
		SetVenueID(req.VenueId).
		//SetSource(req.Source).
		//SetDevice(req.Device).
		Save(o.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Order failed")
		return err
	}

	return nil
}

func (o *Order) UpdateStatus(id int64, status int64) error {
	_, err := o.db.Order.Update().Where(order2.IDEQ(id)).SetStatus(int64(status)).Save(o.ctx)
	return err
}

func (o *Order) GetBySnOrder(sn string) (info *order.OrderInfo, err error) {
	ret, err := o.db.Order.Query().
		Where(order2.OrderSnEQ(sn)).
		Limit(1).
		All(o.ctx)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	if len(ret) == 0 {
		return nil, errors.New("OrdersNoFound")
	}
	row := ret[0]
	info = &order.OrderInfo{
		ID: row.ID,
	}
	return
}

func NewOrder(ctx context.Context, c *app.RequestContext) do.Order {
	return &Order{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
