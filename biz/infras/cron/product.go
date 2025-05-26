package cron

import (
	"context"
	"saas/biz/dal/db"
	product2 "saas/biz/dal/db/ent/product"
	"time"
)

func SetProductStatus() {
	present := time.Now()

	db.DB.Product.Update().
		Where(
			product2.IsSalesIn([]int64{1}...),
			product2.SignSalesAtLTE(present),
			product2.EndSalesAtGTE(present),
		).
		SetIsSales(2).
		Save(context.Background())

	db.DB.Product.Update().
		Where(
			product2.Or(
				product2.IsSalesIn([]int64{2}...),
				product2.SignSalesAtGTE(present),
				product2.EndSalesAtLTE(present),
			),
		).
		SetIsSales(1).
		Save(context.Background())
}
