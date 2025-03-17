namespace go admin.schedule

include "../base/base.thrift"
struct ScheduleInfo  {
	1:optional i64 id =0 (api.raw = "id")
	2:optional string type ="" (api.raw = "type")
	3:optional i64 propertyId =0 (api.raw = "propertyId")
	4:optional i64 venueId =0 (api.raw = "venueId")
	5:optional i64 placeId =0 (api.raw = "placeId")
    6:optional i64 num =0  (api.raw = "num")
	7:optional i64 numSurplus =0  (api.raw = "numSurplus")
	8:optional string data ="" (api.raw = "data")
    9:optional string startTime ="" (api.raw = "startTime")
	10:optional string endTime ="" (api.raw = "endTime")
	11:optional double price =0(api.raw = "price")
	12:optional string name =""(api.raw = "name")
	13:optional string remark =""(api.raw = "remark")
	14:optional i64 coachId =0(api.raw = "coachId")
	15:optional i64 memberId =0(api.raw = "memberId")
	16:optional i64 memberProductId =0(api.raw = "memberProductId")
	17:optional i64 memberProductPropertyId =0(api.raw = "memberProductPropertyId")
	18:optional i64 status =0(api.raw = "status")
	19:optional string propertyName =""(api.raw = "propertyName")
	20:optional string venueName =""(api.raw = "venueName")
	21:optional string placeName =""(api.raw = "placeName")
    22:optional string coachName =""(api.raw = "coachName")
	23:optional string memberName =""(api.raw = "memberName")
	24:optional string memberProductName =""(api.raw = "memberProductName")
	25:optional string memberProductPropertyName =""(api.raw = "memberProductPropertyName")
	26:optional list<ScheduleMemberInfo> scheduleMemberList ={}(api.raw = "scheduleMemberList")
	27:optional list<ScheduleCoachInfo> scheduleCoachList ={}(api.raw = "scheduleCoachList")

	251:optional string createdAt =""(api.raw = "createdAt")
	252:optional string updatedAt =""(api.raw = "updatedAt")
}

struct ScheduleMemberInfo  {
	1:optional i64 id =0 (api.raw = "id")
	2:optional i64 memberId =0 (api.raw = "memberId")
	3:optional i64 venueId =0 (api.raw = "venueId")
	4:optional i64 placeId =0 (api.raw = "placeId")
	5:optional i64 propertyId =0 (api.raw = "propertyId")
	6:optional i64 scheduleId =0 (api.raw = "scheduleId")
	7:optional string scheduleName ="" (api.raw = "scheduleName")
	8:optional string type ="" (api.raw = "type")
	9:optional string startTime ="" (api.raw = "startTime")
	10:optional string endTime ="" (api.raw = "endTime")
	11:optional string signStartTime ="" (api.raw = "signStartTime")
	12:optional string signEndTime ="" (api.raw = "signEndTime")
	13:optional i64 status =0 (api.raw = "status")
	14:optional i64 memberProductId =0 (api.raw = "memberProductId")
	15:optional i64 memberProductItemId =0 (api.raw = "memberProductItemId")

	16:optional string venueName ="" (api.raw = "venueName")
	17:optional string memberName ="" (api.raw = "memberName")
	18:optional string memberProductName ="" (api.raw = "memberProductName")
	19:optional string memberProductPropertyName ="" (api.raw = "memberProductPropertyName")
	20:optional string gender ="" (api.raw = "gender")
	21:optional i64 birthday =0 (api.raw = "birthday")
	22:optional string mobile ="" (api.raw = "mobile")
	251:optional string createdAt =""(api.raw = "createdAt")
    252:optional string updatedAt =""(api.raw = "updatedAt")
}

struct CreateOrUpdateScheduleReq {
    1:optional i64 id =0 (api.raw = "id")
    2:optional string type ="" (api.raw = "type")
    3:optional i64 propertyId =0 (api.raw = "propertyId")
    4:optional i64 venueId =0 (api.raw = "venueId")
    5:optional i64 placeId  =0(api.raw = "placeId")
    6:optional i64 num =0 (api.raw = "num")
    7:optional string startTime=""  (api.raw = "startTime")
    8:optional double price =0 (api.raw = "price")
    9:optional string remark=""  (api.raw = "remark")
    10:optional i64 coachId =0 (api.raw = "coachId")
    11:optional i64 memberId =0 (api.raw = "memberId")
    12:optional i64 memberProductId =0 (api.raw = "memberProductId")
    13:optional i64 memberProductPropertyId =0 (api.raw = "memberProductPropertyId")
}

struct ScheduleListReq {
    1:  optional i64 page=0 (api.raw = "page")
    2:  optional i64 pageSize=0 (api.raw = "pageSize")
    3:  optional i64 memberId=0 (api.raw = "memberId")
    4:  optional list<i64> coachIds =""(api.raw = "coachIds")
    5:  optional list<i64> productIds =""(api.raw = "productIds")
    6:  optional i64 venueId=0 (api.raw = "venueId")
    7:  optional list<i64> propertyIds=0 (api.raw = "propertyIds")
    8:  optional string startTime ="" (api.raw = "startTime")
    9:  optional string type=""  (api.raw = "type")
}
struct ScheduleMemberReq {
    1:  optional i64 page =0(api.raw = "page")
    2:  optional i64 pageSize=0 (api.raw = "pageSize")
    3:  optional i64 memberId=0 (api.raw = "memberId")
    4:  optional i64 scheduleId=0 (api.raw = "scheduleId")
    5:  optional string type ="" (api.raw = "type")
}
struct ScheduleCoachReq{
    1:  optional i64 page=0 (api.raw = "page")
    2:  optional i64 pageSize =0(api.raw = "pageSize")
    3:  optional i64 memberId=0 (api.raw = "memberId")
    4:  optional i64 scheduleId=0 (api.raw = "scheduleId")
    5:  optional string type ="" (api.raw = "type")
}

struct SearchSubscribeByMemberReq{
    1:  optional i64 propertyId=0 (api.raw = "propertyId")
    2:  optional string	mobile ="" (api.raw = "mobile")
    3:  optional i64 venueId=0 (api.raw = "venueId")
    4:  optional string	type ="" (api.raw = "type")
}

struct MemberSubscribeReq{
    1:  optional list<i64> memberProductPropertyId =0 (api.raw = "memberProductPropertyId")
    2:  optional i64 scheduleId =0(api.raw = "scheduleId")
    3:  optional string remark =""(api.raw = "remark")
}


struct ScheduleMemberCreate  {
	MemberProductPropertyID []int64 (api.raw = "memberProductPropertyId")
	Schedule                int64   (api.raw = "schedule")
	Remark                  string  (api.raw = "remark")
}
struct ScheduleMemberListReq  {
	Type     string (api.raw = "type")
	Page     int64  (api.raw = "page")
	PageSize int64  (api.raw = "pageSize")
	Member   int64  (api.raw = "member")
	Schedule int64  (api.raw = "schedule")
}

struct ScheduleCoachInfo  {
	1:optional i64 id =0 (api.raw = "id")
	CoachId       int64  (api.raw = "CoachId")
	VenueId       int64  (api.raw = "VenueId")
	PlaceID       int64  (api.raw = "PlaceId")
	PropertyId    int64  (api.raw = "PropertyId")
	ScheduleId    int64  (api.raw = "ScheduleId")
	Type          string (api.raw = "type")
	Date          string (api.raw = "date")
	StartTime     string (api.raw = "startTime")
	EndTime       string (api.raw = "endTime")
	SignStartTime string (api.raw = "signStartTime")
	SignEndTime   string (api.raw = "signEndTime")
	Status        int64  (api.raw = "status")

	ScheduleName string (api.raw = "scheduleName")
	PropertyName string (api.raw = "propertyName")
	VenueName    string (api.raw = "venueName")
	PlaceName    string (api.raw = "placeName")
	CoachName    string (api.raw = "coachName")
	CoachAvatar  string (api.raw = "coachAvatar")

	Mobile                    string (api.raw = "mobile")
	MemberName                string (api.raw = "memberName")
	MemberAvatar              string (api.raw = "memberAvatar")
	MemberProductName         string (api.raw = "memberProductName")
	MemberProductPropertyName string (api.raw = "memberProductPropertyName")
	Remark                    string (api.raw = "remark")
	MRemark                   string (api.raw = "mRemark")
		251:optional string createdAt =""(api.raw = "createdAt")
    	252:optional string updatedAt =""(api.raw = "updatedAt")
}

struct ScheduleCoachListReq  {
	Type     string (api.raw = "type")
	Page     int64  (api.raw = "page")
	PageSize int64  (api.raw = "pageSize")
	Coach    int64  (api.raw = "coach")
	Schedule int64  (api.raw = "schedule")
}

struct SubscribeByMember  {
	Avatar                  string (api.raw = "avatar")
	Mobile                  string (api.raw = "mobile")
	MemberId                int64  (api.raw = "memberId")
	MemberProductID         int64  (api.raw = "memberProductID")
	MemberProductPropertyId int64  (api.raw = "memberProductPropertyId")

	MemberName                string (api.raw = "memberName")
	MemberProductName         string (api.raw = "memberProductName")
	MemberProductPropertyName string (api.raw = "memberProductPropertyName")
}



service ScheduleService {
    base.NilResponse CreateSchedule(1: CreateOrUpdateScheduleReq req)  (api.post = "/api/admin/schedule/create")

    base.NilResponse UpdateSchedule(1: CreateOrUpdateScheduleReq req) (api.post = "/api/admin/schedule/update")

    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/schedule/status")

    base.NilResponse ListSchedule(1: ScheduleListReq req )(api.post = "/api/admin/schedule/list")

    base.NilResponse DateListSchedule(1: ScheduleListReq req )(api.post = "/api/admin/schedule/date-list")

    base.NilResponse GetScheduleById(1: base.IdReq req) (api.post = "/api/admin/schedule/info")

    base.NilResponse GetScheduleMemberList(1: ScheduleMemberReq req) (api.post = "/api/admin/schedule/schedule-member-list")

    base.NilResponse SearchSubscribeByMember(1: SearchSubscribeByMemberReq req) (api.post = "/api/admin/schedule/search-subscribe-by-member")

    base.NilResponse MemberSubscribe(1: MemberSubscribeReq req) (api.post = "/api/admin/schedule/member-subscribe")

    base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/schedule/schedule-member-status")


    base.NilResponse GetScheduleCoachList(1: ScheduleMemberReq req) (api.post = "/api/admin/schedule/schedule-coach-list")

    base.NilResponse UpdateCoachStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/schedule/schedule-coach-status")

}