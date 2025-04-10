package do

import "kcers/idl_gen/model/admin/schedule"

type Schedule interface {
	ScheduleCreate(req *schedule.CreateOrUpdateScheduleReq) error
	ScheduleUpdate(req *schedule.CreateOrUpdateScheduleReq) error
	ScheduleDelete(id int64) error
	ScheduleList(req *schedule.ScheduleListReq) (resp []*schedule.ScheduleInfo, total int, err error)
	ScheduleDateList(req *schedule.ScheduleListReq) (map[string][]*schedule.ScheduleInfo, int, error)
	ScheduleUpdateStatus(id, status int64) error
	ScheduleInfo(id int64) (roleInfo *schedule.ScheduleInfo, err error)

	MemberCreate(req *schedule.MemberSubscribeReq) error
	MemberUpdate(req *schedule.ScheduleMemberInfo) error
	MemberDelete(id int64) error
	MemberList(req *schedule.ScheduleMemberListReq) (resp []*schedule.ScheduleMemberInfo, total int, err error)
	UpdateMemberStatus(id, status int64) error
	MemberInfo(id int64) (roleInfo *schedule.ScheduleMemberInfo, err error)
	SearchSubscribeByMember(req *schedule.SearchSubscribeByMemberReq) (list []*schedule.SubscribeByMember, total int64, err error)

	CoachCreate(req *schedule.ScheduleCoachInfo) error
	CoachUpdate(req *schedule.ScheduleCoachInfo) error
	CoachDelete(id int64) error
	CoachList(req *schedule.ScheduleCoachListReq) (resp []*schedule.ScheduleCoachInfo, total int, err error)
	CoachUpdateStatus(ID int64, status int64) error
	CoachInfo(id int64) (info *schedule.ScheduleCoachInfo, err error)
}
