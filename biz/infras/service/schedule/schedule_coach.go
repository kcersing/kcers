package schedule

import (
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent/predicate"
	schedule2 "kcers/biz/dal/db/mysql/ent/schedule"
	"kcers/biz/dal/db/mysql/ent/schedulecoach"
	"kcers/idl_gen/model/admin/schedule"
)

func (s *Schedule) CoachCreate(req *schedule.ScheduleCoachInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *Schedule) CoachUpdate(req *schedule.ScheduleCoachInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *Schedule) CoachDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Schedule) CoachList(req *schedule.ScheduleCoachListReq) (resp []*schedule.ScheduleCoachInfo, total int, err error) {

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

	for _, v := range lists {
		resp = append(resp, s.entScheduleCoachInfo(v, nil))
	}

	total, _ = s.db.ScheduleCoach.Query().Where(predicates...).Count(s.ctx)
	return

}

func (s *Schedule) CoachUpdateStatus(id int64, status int64) error {
	sc, err := s.db.ScheduleCoach.Query().Where(schedulecoach.IDEQ(id)).First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule Coach info failed")
		return err
	}
	_, err = s.db.ScheduleCoach.Update().Where(schedulecoach.ID(id)).SetStatus(status).Save(s.ctx)
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

func (s *Schedule) CoachInfo(id int64) (info *schedule.ScheduleCoachInfo, err error) {
	m, err := s.db.ScheduleCoach.Query().Where(schedulecoach.ID(id)).First(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get Schedule Member failed")
		return info, err
	}

	info = s.entScheduleCoachInfo(m, nil)

	return info, nil
}
