package schedule

import (
	"fmt"
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent/member"
	"kcers/biz/dal/db/mysql/ent/memberproduct"
	"kcers/biz/dal/db/mysql/ent/memberproductproperty"
	"kcers/biz/dal/db/mysql/ent/productproperty"
	"kcers/biz/dal/db/mysql/ent/user"
	"kcers/biz/dal/db/mysql/ent/venue"
	"kcers/biz/dal/db/mysql/ent/venueplace"
	"kcers/biz/infras/service"
	"kcers/idl_gen/model/admin/schedule"
	"time"
)

func (s *Schedule) ScheduleCreate(req schedule.CreateOrUpdateScheduleReq) error {

	// 解析字符串到 time.Time 类型
	startTime, _ := time.Parse(time.DateTime, req.StartAt)
	tx, err := s.db.Tx(s.ctx)
	if err != nil {
		return fmt.Errorf("starting a transaction: %w", err)
	}

	if req.Type == "course" {
		mpp, err := tx.MemberProductProperty.Query().
			Where(memberproductproperty.ID(req.MemberProductPropertyId)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员产品属性")
			return service.Rollback(tx, err)
		}
		req.PropertyId = mpp.PropertyID
	}

	property, err := s.db.ProductProperty.Query().
		Where(productproperty.ID(req.PropertyId)).
		First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "未查询到改课程属性")
		return err
	}

	venueName := ""
	if req.VenueId != 0 {
		ven, err := s.db.Venue.Query().
			Where(venue.ID(req.PropertyId)).
			First(s.ctx)
		if err == nil {
			venueName = ven.Name
		}
	}
	placeName := ""
	if req.PlaceId != 0 {
		place, err := s.db.VenuePlace.Query().
			Where(venueplace.ID(req.PropertyId)).
			First(s.ctx)
		if err == nil {
			placeName = place.Name
		}
	}

	one, err := tx.Schedule.Create().
		SetType(req.Type).
		SetStatus(1).
		SetVenueID(req.VenueId).
		SetPlaceID(req.PlaceId).
		SetPropertyID(req.PropertyId).
		SetNum(req.Num).
		SetNumSurplus(req.Num).
		SetDate(startTime.Format(time.DateOnly)).
		SetStartAt(startTime).
		SetEndAt(startTime.Add(time.Duration(property.Length) * time.Minute)).
		SetPrice(req.Price).
		SetRemark(req.Remark).
		SetLength(property.Length).
		SetName(property.Name).
		SetVenueName(venueName).
		SetPlaceName(placeName).
		Save(s.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Course Record Schedule failed")
		return service.Rollback(tx, err)
	}

	if req.CoachId != 0 {

		u, err := tx.User.Query().
			Where(user.ID(req.CoachId)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该员工")
			return service.Rollback(tx, err)
		}

		_, err = tx.ScheduleCoach.Create().
			SetSchedule(one).
			SetScheduleName(one.Name).
			SetType(req.Type).
			SetVenueID(req.VenueId).
			SetCoachID(req.CoachId).
			SetStartAt(startTime).
			SetEndAt(startTime.Add(time.Duration(property.Length) * time.Minute)).
			SetStatus(1).
			SetCoachName(u.Name).
			Save(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return service.Rollback(tx, err)
		}

	}

	if req.MemberId != 0 {
		m, err := tx.Member.Query().
			Where(member.ID(req.MemberId)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员")
			return service.Rollback(tx, err)
		}
		mp, err := tx.MemberProduct.Query().
			Where(memberproduct.ID(req.MemberProductId)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员产品")
			return service.Rollback(tx, err)
		}
		mpp, err := tx.MemberProductProperty.Query().
			Where(memberproductproperty.ID(req.MemberProductPropertyId)).
			First(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "未查询到该会员产品属性")
			return service.Rollback(tx, err)
		}
		if (mpp.CountSurplus - 1) < 0 {
			err = errors.Wrap(err, "该会员课程不足")
			return service.Rollback(tx, err)
		}

		_, err = tx.ScheduleMember.Create().
			SetSchedule(one).
			SetType(req.Type).
			SetMemberID(req.MemberId).
			SetVenueID(req.VenueId).
			SetStartAt(startTime).
			SetEndAt(startTime.Add(time.Duration(property.Length) * time.Minute)).
			SetMemberProductID(req.MemberProductId).
			SetMemberProductPropertyID(req.MemberProductPropertyId).
			SetStatus(0).
			SetMemberName(m.Name).
			SetMemberProductName(mp.Name).
			SetMemberProductPropertyName(mpp.Name).
			Save(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "create Course Record Coach failed")
			return service.Rollback(tx, err)
		}

		_, err = tx.MemberProductProperty.
			Update().
			Where(memberproductproperty.IDEQ(req.MemberProductPropertyId)).
			AddCountSurplus(-1).
			Save(s.ctx)
		if err != nil {
			err = errors.Wrap(err, "Member Product Item failed")
			return service.Rollback(tx, err)
		}

	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
