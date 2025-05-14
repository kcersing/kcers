package cron

import (
	"context"

	"time"
)

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
