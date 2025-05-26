package cron

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"saas/biz/dal/db"
	schedule2 "saas/biz/dal/db/ent/schedule"
	schedulecoach2 "saas/biz/dal/db/ent/schedulecoach"
	schedulemember2 "saas/biz/dal/db/ent/schedulemember"
	"saas/biz/infras/service"
	"time"
)

// SetScheduleStatus 设置课程状态
func SetScheduleStatus() {
	present := time.Now()
	db.DB.Schedule.Update().
		Where(
			schedule2.StatusIn([]int64{2}...),
			//小于等于
			schedule2.StartTimeLTE(present),
			//大于等于
			//schedule2.EndTimeGTE(present),
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

	//db.DB.Schedule.Update().
	//	Where(
	//		schedule2.HasCoachsWith(schedulecoach2.StatusEQ(4)),
	//		schedule2.StatusNEQ(6),
	//	).
	//	SetStatus(6).
	//	Save(context.Background())

}

// SetScheduleMemberStatus 设置课程成员状态
func SetScheduleMemberStatus() {
	present := time.Now()

	db.DB.ScheduleMember.Update().
		Where(
			schedulemember2.StatusIn([]int64{1}...),
			schedulemember2.EndTimeLTE(present),
		).
		SetStatus(4).
		Save(context.Background())

	db.DB.Debug().ScheduleMember.Update().
		Where(
			schedulemember2.StatusIn([]int64{2}...),
			schedulemember2.EndTimeLTE(present),
		).
		SetStatus(3).
		Save(context.Background())
}

// SetDeductMemberNumberSurplus 扣除会员产品剩余数量
func SetDeductMemberNumberSurplus() {
	all, _ := db.DB.ScheduleMember.Query().
		Where(
			schedulemember2.StatusNotIn([]int64{1, 2}...),
			schedulemember2.IsDeduct(0),
		).
		All(context.Background())
	for _, v := range all {
		_, err := db.DB.ScheduleMember.UpdateOne(v).SetIsDeduct(1).Save(context.Background())
		if err != nil {
			hlog.Errorf("扣除会员产品剩余数量失败：%v", err)
		}

		err = service.DeductNumberSurplus(db.DB, v.MemberProductID)
		if err != nil {
			hlog.Errorf("扣除会员产品剩余数量失败ID：%v", v.MemberProductID)
			hlog.Errorf("扣除会员产品剩余数量失败原因：%v", err)
		}
	}

}

// SetScheduleCoachStatus 设置课程教练状态
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
