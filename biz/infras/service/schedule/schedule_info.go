package schedule

import (
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent"
	member2 "kcers/biz/dal/db/mysql/ent/member"

	"kcers/biz/dal/db/mysql/ent/memberprofile"
	schedule2 "kcers/biz/dal/db/mysql/ent/schedule"
	"kcers/biz/dal/db/mysql/ent/venueplace"
	"kcers/biz/dal/enums"
	"kcers/idl_gen/model/admin/schedule"
	"time"
)

func (s *Schedule) ScheduleInfo(id int64) (resp *schedule.ScheduleInfo, err error) {
	one, err := s.db.Schedule.Query().Where(schedule2.IDEQ(id)).First(s.ctx)
	if err != nil {
		err = errors.Wrap(err, "get venue failed")
		return resp, err
	}
	resp = s.entScheduleInfo(one)

	coachAll, _ := one.QueryCoachs().All(s.ctx)
	if coachAll != nil {
		for _, v := range coachAll {
			resp.CoachCourseRecord = append(resp.CoachCourseRecord, s.entScheduleCoachInfo(v, one))
		}
	}

	memberAll, _ := one.QueryMembers().All(s.ctx)
	if memberAll != nil {
		for _, v := range memberAll {
			resp.MemberCourseRecord = append(resp.MemberCourseRecord, s.entScheduleMemberInfo(v, one))
		}
	}

	return
}

func (s *Schedule) entScheduleMemberInfo(req *ent.ScheduleMember, sch *ent.Schedule) (info *schedule.ScheduleMemberInfo) {

	info = &schedule.ScheduleMemberInfo{

		ID:                        req.ID,
		MemberId:                  req.MemberID,
		VenueId:                   req.VenueID,
		ScheduleId:                req.ScheduleID,
		Type:                      req.Type,
		CreatedAt:                 req.CreatedAt.Format(time.DateTime),
		UpdatedAt:                 req.UpdatedAt.Format(time.DateTime),
		Status:                    req.Status,
		MemberProductId:           req.MemberProductID,
		MemberName:                req.MemberName,
		MemberProductName:         req.MemberProductName,
		Date:                      req.Date.Format(time.DateOnly),
		Seat:                      &req.Seat,
		StartAt:                   req.StartAt.Format("15:04"),
		EndAt:                     req.EndAt.Format("15:04"),
		SignStartAt:               req.SignStartAt.Format("15:04"),
		SignEndAt:                 req.SignEndAt.Format("15:04"),
		MemberProductPropertyId:   req.MemberProductPropertyID,
		MemberProductPropertyName: req.MemberProductPropertyName,
	}
	m, _ := s.db.Member.Query().Where(member2.ID(req.MemberID)).First(s.ctx)

	if m != nil {
		info.Mobile = m.Mobile
		mProfile, _ := s.db.MemberProfile.Query().Where(memberprofile.ID(req.MemberID)).First(s.ctx)
		info.Gender = enums.ReturnMemberGenderValues(mProfile.Gender)

		if !mProfile.Birthday.IsZero() {
			info.Birthday = int64(time.Now().Sub(mProfile.Birthday).Hours() / 24 / 365)
		}
	}

	if sch != nil {
		sch, _ = req.QuerySchedule().First(s.ctx)
	}
	if sch != nil {
		info.PlaceId = sch.PlaceID
		info.PropertyId = sch.PropertyID
		info.ScheduleName = sch.Name
		info.VenueName = sch.VenueName
	}

	return
}
func (s *Schedule) entScheduleCoachInfo(req *ent.ScheduleCoach, sch *ent.Schedule) (info *schedule.ScheduleCoachInfo) {

	info = &schedule.ScheduleCoachInfo{
		ID:          req.ID,
		CoachId:     req.CoachID,
		VenueId:     req.VenueID,
		ScheduleId:  req.ScheduleID,
		Type:        req.Type,
		CreatedAt:   req.CreatedAt.Format(time.DateTime),
		UpdatedAt:   req.UpdatedAt.Format(time.DateTime),
		Status:      req.Status,
		Date:        req.Date.Format(time.DateOnly),
		CoachName:   req.CoachName,
		StartAt:     req.StartAt.Format("15:04"),
		EndAt:       req.EndAt.Format("15:04"),
		SignStartAt: req.SignStartAt.Format("15:04"),
		SignEndAt:   req.SignEndAt.Format("15:04"),
		CoachAvatar: "",
	}
	if sch == nil {
		sch, _ = req.QuerySchedule().First(s.ctx)
	}
	if sch != nil {

		info.PlaceId = sch.PlaceID
		info.PropertyId = sch.PropertyID
		info.ScheduleName = sch.Name
		info.PropertyName = sch.PlaceName
		info.VenueName = sch.VenueName
		info.PlaceName = sch.PlaceName
		info.Remark = sch.Remark
	}

	return
}

func (s *Schedule) entScheduleInfo(req *ent.Schedule) (info *schedule.ScheduleInfo) {

	info = &schedule.ScheduleInfo{
		ID:         req.ID,
		Type:       req.Type,
		VenueId:    req.VenueID,
		PlaceId:    req.PlaceID,
		Num:        req.Num,
		NumSurplus: req.NumSurplus,
		Name:       req.Name,
		Status:     req.Status,
		VenueName:  req.VenueName,
		PlaceName:  req.PlaceName,
		CreatedAt:  req.CreatedAt.Format(time.DateTime),
		UpdatedAt:  req.UpdatedAt.Format(time.DateTime),

		PropertyId:   req.PropertyID,
		Date:         req.Date.Format(time.DateOnly),
		StartAt:      req.StartAt.Format("15:04"),
		EndAt:        req.EndAt.Format("15:04"),
		Price:        req.Price,
		Remark:       req.Remark,
		PropertyName: req.Name,

		MemberCourseRecord: nil,
		CoachCourseRecord:  nil,
	}

	seats, _ := s.db.VenuePlace.Query().Where(venueplace.IDEQ(req.PlaceID)).First(s.ctx)
	if seats != nil {
		info.Seats = seats.Seat
	}

	return
}
