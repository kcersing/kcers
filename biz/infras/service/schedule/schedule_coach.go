package schedule

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent/predicate"
	schedule2 "kcers/biz/dal/db/mysql/ent/schedule"
	"kcers/biz/dal/db/mysql/ent/schedulecoach"
	"kcers/biz/infras/service/member"
	"kcers/biz/infras/service/user"
	"kcers/biz/infras/service/venue"
	"kcers/idl_gen/model/admin/schedule"

	"time"
)

func (s *Schedule) CoachCreate(req schedule.ScheduleCoachInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *Schedule) CoachUpdate(req schedule.ScheduleCoachInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *Schedule) CoachDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Schedule) CoachList(req schedule.ScheduleCoachListReq) (resp []*schedule.ScheduleCoachInfo, total int, err error) {

	var predicates []predicate.ScheduleCoach
	if req.CoachId > 0 {
		predicates = append(predicates, schedulecoach.CoachIDEQ(req.CoachId))
	}
	if req.ScheduleId > 0 {
		predicates = append(predicates, schedulecoach.ScheduleID(req.ScheduleId))
	}
	if req.Type != "" {
		predicates = append(predicates, schedulecoach.Type(req.Type))
	}
	lists, err := s.db.ScheduleCoach.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule Coach list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Schedule Coach info failed")
		return resp, 0, err
	}
	for i, v := range lists {
		vInfo, err := venue.NewVenue(s.ctx, s.c).VenueInfo(v.VenueID)
		if err == nil {
			resp[i].VenueName = vInfo.Name
		}
		coach, err := user.NewUser(s.ctx, s.c).Info(v.CoachID)
		if err == nil {
			resp[i].CoachAvatar = coach.Avatar
		}

		sl, err := v.QuerySchedule().First(s.ctx)
		if err == nil {
			resp[i].ScheduleName = sl.Name
			resp[i].VenueName = sl.VenueName
			resp[i].PlaceName = sl.PlaceName
			resp[i].Remark = sl.Remark
		}
		sm, err := sl.QueryMembers().First(s.ctx)
		if err == nil {
			resp[i].MRemark = sm.Remark
			resp[i].MemberName = sm.MemberName
			resp[i].MemberProductName = sm.MemberProductName
			resp[i].MemberProductPropertyName = sm.MemberProductPropertyName

			m, err := member.NewMember(s.ctx, s.c).Info(sm.MemberID)
			if err == nil {
				resp[i].Mobile = m.Mobile
				resp[i].MemberAvatar = m.Avatar
			}
		}

		resp[i].CreatedAt = v.CreatedAt.Format(time.DateTime)
		resp[i].StartAt = v.StartAt.Format(time.TimeOnly)
		resp[i].EndAt = v.EndAt.Format(time.TimeOnly)
		resp[i].SignStartAt = v.SignStartAt.Format(time.TimeOnly)
		resp[i].SignEndAt = v.SignEndAt.Format(time.TimeOnly)
		resp[i].Date = v.StartAt.Format(time.DateOnly)

	}

	total, _ = s.db.ScheduleCoach.Query().Where(predicates...).Count(s.ctx)
	return

}

func (s *Schedule) CoachUpdateStatus(ID int64, status int64) error {
	sc, err := s.db.ScheduleCoach.Query().Where().First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule Coach info failed")
		return err
	}
	_, err = s.db.ScheduleCoach.Update().Where(schedulecoach.ID(ID)).SetStatus(status).Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Schedule Coach status failed")
		return err
	}
	_, err = s.db.Schedule.Update().Where(schedule2.ID(sc.ScheduleID)).SetStatus(status).Save(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Schedule status failed")
		return err
	}

	return nil
}

func (s *Schedule) CoachInfo(ID int64) (roleInfo *schedule.ScheduleCoachInfo, err error) {
	//TODO implement me
	panic("implement me")
}
