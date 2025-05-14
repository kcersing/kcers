package cron

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"time"
)

func SetScheduleStatus() {
	present := time.Now()
	db.DB.Schedule.Update().
		Where(
			schedule2.StatusIn([]int64{2}...),
			schedule2.StartTimeLTE(present),
			schedule2.EndTimeGTE(present),
		).
		SetStatus(3).
		Save(context.Background())
	db.DB.Schedule.Update().
		Where(
			schedule2.StatusIn([]int64{3}...),
			schedule2.EndTimeLTE(present),
		).
		SetStatus(4).
		Save(context.Background())
}

func SetScheduleMemberStatus() {
	present := time.Now()

	db.DB.ScheduleMember.Update().
		Where(
			schedulemember2.StatusIn([]int64{1}...),
			schedulemember2.EndTimeLTE(present),
		).
		SetStatus(4).
		Save(context.Background())

	db.DB.ScheduleMember.Update().
		Where(
			schedulemember2.StatusIn([]int64{2}...),
			schedulemember2.EndTimeLTE(present),
		).
		SetStatus(3).
		Save(context.Background())
}

func SetDeductMemberNumberSurplus() {
	all, _ := db.DB.ScheduleMember.Query().
		Where(
			schedulemember2.StatusIn([]int64{3, 4}...),
			schedulemember2.IsDeduct(0),
		).
		All(context.Background())
	for _, v := range all {
		err := service.DeductNumberSurplus(db.DB, v.MemberProductID)
		if err != nil {
			hlog.Errorf("扣除会员产品剩余数量失败ID：%v", v.MemberProductID)
			hlog.Errorf("扣除会员产品剩余数量失败原因：%v", err)
		}
	}

}

func SetScheduleCoachStatus() {
	present := time.Now()
	db.DB.ScheduleCoach.Update().
		Where(
			schedulecoach2.StatusIn([]int64{1}...),
			schedulecoach2.EndTimeLTE(present),
		).
		SetStatus(4).
		Save(context.Background())
	db.DB.ScheduleCoach.Update().
		Where(
			schedulecoach2.StatusIn([]int64{2}...),
			schedulecoach2.EndTimeLTE(present),
		).
		SetStatus(3).
		Save(context.Background())
}
