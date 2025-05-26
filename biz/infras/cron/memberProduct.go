package cron

import (
	"context"
	"saas/biz/dal/db"
	memberproduct2 "saas/biz/dal/db/ent/memberproduct"
	"time"
)

// GT >
// GTE >=
// LT <
// LTE <=
func SetMemberProductStatus() {
	present := time.Now()

	db.DB.MemberProduct.Update().
		Where(
			memberproduct2.StatusEQ(1),
			memberproduct2.ValidityAtLTE(present),
		).
		SetStatus(2).
		Save(context.Background())

	db.DB.MemberProduct.Update().
		Where(
			memberproduct2.StatusEQ(2),
			memberproduct2.CancelAtLTE(present),
		).
		SetStatus(3).
		Save(context.Background())

}
