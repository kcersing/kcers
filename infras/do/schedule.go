package do

type Schedule interface {
	ScheduleCreate(req CreateOrUpdateScheduleReq) error
	ScheduleUpdate(req CreateOrUpdateScheduleReq) error
	ScheduleDelete(id int64) error
	ScheduleList(req ScheduleListReq) (resp []*ScheduleInfo, total int, err error)
	ScheduleDateList(req ScheduleListReq) (map[string][]*ScheduleInfo, int, error)
	ScheduleUpdateStatus(ID int64, status int64) error
	ScheduleInfo(ID int64) (roleInfo *ScheduleInfo, err error)

	MemberCreate(req ScheduleMemberCreate) error
	MemberUpdate(req ScheduleMemberInfo) error
	MemberDelete(id int64) error
	MemberList(req ScheduleMemberListReq) (resp []*ScheduleMemberInfo, total int, err error)
	UpdateMemberStatus(ID int64, status int64) error
	MemberInfo(ID int64) (roleInfo *ScheduleMemberInfo, err error)
	SearchSubscribeByMember(req SearchSubscribeByMemberReq) (list []SubscribeByMember, total int64, err error)

	CoachCreate(req ScheduleCoachInfo) error
	CoachUpdate(req ScheduleCoachInfo) error
	CoachDelete(id int64) error
	CoachList(req ScheduleCoachListReq) (resp []*ScheduleCoachInfo, total int, err error)
	CoachUpdateStatus(ID int64, status int64) error
	CoachInfo(ID int64) (roleInfo *ScheduleCoachInfo, err error)
}
