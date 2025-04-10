package schedule

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"kcers/biz/dal/cache"
	casbin2 "kcers/biz/dal/casbin"
	"kcers/biz/dal/config"
	db "kcers/biz/dal/db/mysql"
	"kcers/biz/dal/db/mysql/ent"
	"kcers/biz/dal/db/mysql/ent/predicate"
	schedule2 "kcers/biz/dal/db/mysql/ent/schedule"
	"kcers/biz/infras/do"
	"kcers/idl_gen/model/admin/schedule"
	"time"
)

type Schedule struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	Cbs   *casbin.Enforcer
	db    *ent.Client
	cache *ristretto.Cache
}

func NewSchedule(ctx context.Context, c *app.RequestContext) do.Schedule {
	return &Schedule{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
		Cbs:   casbin2.CasbinEnforcer(),
	}
}

func (s *Schedule) ScheduleUpdate(req *schedule.CreateOrUpdateScheduleReq) error {

	return nil
}

func (s *Schedule) ScheduleDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Schedule) ScheduleList(req *schedule.ScheduleListReq) (resp []*schedule.ScheduleInfo, total int, err error) {
	var predicates []predicate.Schedule

	if req.StartAt != "" {
		startTime, _ := time.Parse(time.DateOnly, req.StartAt)
		//大于
		predicates = append(predicates, schedule2.StartAtGTE(startTime))
		//小于
		predicates = append(predicates, schedule2.EndAtLTE(startTime.Add(7*24*time.Hour)))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, schedule2.VenueID(req.VenueId))
	}

	lists, err := s.db.Schedule.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, s.entScheduleInfo(v))
	}

	total, _ = s.db.Schedule.Query().Where(predicates...).Count(s.ctx)
	return
}

func (s *Schedule) ScheduleDateList(req *schedule.ScheduleListReq) (map[string][]*schedule.ScheduleInfo, int, error) {
	req.Page = 1
	req.PageSize = 1000
	lists, total, err := s.ScheduleList(req)
	m := make(map[string][]*schedule.ScheduleInfo)
	for _, v := range lists {
		m[v.Date] = append(m[v.Date], v)
	}

	return m, total, err
}

func (s *Schedule) ScheduleUpdateStatus(ID int64, status int64) error {
	_, err := s.db.Schedule.Update().Where(schedule2.IDEQ(ID)).SetStatus(status).Save(s.ctx)
	return err
}
