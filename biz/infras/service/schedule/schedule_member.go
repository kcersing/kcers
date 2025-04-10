package schedule

import (
	"fmt"
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent/member"
	"kcers/biz/dal/db/mysql/ent/memberproduct"
	"kcers/biz/dal/db/mysql/ent/memberproductproperty"
	"kcers/biz/dal/db/mysql/ent/predicate"
	schedule2 "kcers/biz/dal/db/mysql/ent/schedule"
	"kcers/biz/dal/db/mysql/ent/schedulemember"
	"kcers/biz/dal/db/mysql/ent/venue"
	"kcers/biz/dal/minio"
	"kcers/idl_gen/model/admin/schedule"
)

func (s *Schedule) MemberCreate(req *schedule.MemberSubscribeReq) error {

	sc, err := s.db.Schedule.Query().
		Where(schedule2.ID(req.ScheduleId)).
		First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "未查询到该课程")
		return err
	}

	if (sc.NumSurplus - int64(len(req.MemberProductPropertyId))) < 0 {
		err = errors.Wrap(errors.New("超出约课人数"+fmt.Sprintf("%+v", sc.NumSurplus-int64(len(req.MemberProductPropertyId)))), "未查询到该课程")
		return err
	}
	for _, v := range req.MemberProductPropertyId {

		mpp, err := s.db.MemberProductProperty.Query().
			Where(memberproductproperty.ID(v)).
			First(s.ctx)

		if err != nil {
			err = errors.Wrap(err, "未查询到该会员属性")
			return err
		}
		m, err := s.db.Member.Query().
			Where(member.ID(mpp.MemberID)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员")
			return err
		}
		mp, err := s.db.MemberProduct.Query().
			Where(memberproduct.ID(mpp.MemberProductID)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员产品")
			return err
		}

		_, err = s.db.ScheduleMember.Create().
			SetSchedule(sc).
			SetScheduleName(sc.Name).
			SetMemberName(m.Name).
			SetStatus(1).
			SetStartAt(sc.StartAt).
			SetEndAt(sc.EndAt).
			SetMemberID(m.ID).
			SetType(sc.Type).
			SetVenueID(sc.VenueID).
			SetMemberProductID(mp.ID).
			SetMemberProductName(mp.Name).
			SetMemberProductPropertyID(mpp.ID).
			SetMemberProductPropertyName(mpp.Name).
			Save(s.ctx)
		if err != nil {
			return err
		}

	}

	return nil

}

func (s *Schedule) MemberUpdate(req *schedule.ScheduleMemberInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *Schedule) MemberDelete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Schedule) MemberList(req *schedule.ScheduleMemberListReq) (resp []*schedule.ScheduleMemberInfo, total int, err error) {

	var predicates []predicate.ScheduleMember
	if req.MemberId > 0 {
		predicates = append(predicates, schedulemember.MemberID(req.MemberId))
	}
	if req.ScheduleId > 0 {
		predicates = append(predicates, schedulemember.ScheduleID(req.ScheduleId))
	}

	lists, err := s.db.ScheduleMember.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Schedule Member list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, s.entScheduleMemberInfo(v, nil))

	}

	total, _ = s.db.ScheduleMember.Query().Where(predicates...).Count(s.ctx)
	return

}

func (s *Schedule) UpdateMemberStatus(id, status int64) error {
	_, err := s.db.ScheduleMember.Update().Where(schedulemember.ID(id)).SetStatus(status).Save(s.ctx)
	return err
}

func (s *Schedule) MemberInfo(id int64) (info *schedule.ScheduleMemberInfo, err error) {

	m, err := s.db.ScheduleMember.Query().Where(schedulemember.ID(id)).First(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get Schedule Member failed")
		return info, err
	}

	info = s.entScheduleMemberInfo(m, nil)

	return info, nil

}

func (s *Schedule) SearchSubscribeByMember(req *schedule.SearchSubscribeByMemberReq) (list []*schedule.SubscribeByMember, total int64, err error) {

	m, err := s.db.Member.Query().Where(member.Mobile(req.Mobile)).First(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return nil, 0, err
	}

	mpp, err := s.db.MemberProductProperty.
		Query().
		Where(
			memberproductproperty.PropertyID(req.PropertyId),
			memberproductproperty.HasVenuesWith(venue.ID(req.VenueId)),
		).
		All(s.ctx)

	if err != nil {
		return nil, 0, err
	}
	for _, v := range mpp {
		mp, _ := v.QueryOwner().First(s.ctx)
		list = append(list, &schedule.SubscribeByMember{
			Avatar:                    minio.URLconvert(s.ctx, s.c, m.Avatar),
			Mobile:                    m.Mobile,
			MemberId:                  m.ID,
			MemberName:                m.Name,
			MemberProductID:           v.MemberProductID,
			MemberProductPropertyId:   v.ID,
			MemberProductName:         mp.Name,
			MemberProductPropertyName: v.Name,
		})
	}

	total = int64(len(list))
	return
}
