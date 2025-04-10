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
    9:optional string startAt ="" (api.raw = "startAt")
	10:optional string endAt ="" (api.raw = "endAt")
	11:optional double price =0(api.raw = "price")
	12:optional string name =""(api.raw = "name")
	13:optional string remark =""(api.raw = "remark")
//	14:optional i64 coachId =0(api.raw = "coachId")
//	15:optional i64 memberId =0(api.raw = "memberId")
//	16:optional i64 memberProductId =0(api.raw = "memberProductId")
//	17:optional i64 memberProductPropertyId =0(api.raw = "memberProductPropertyId")
	18:optional i64 status =0(api.raw = "status")
	19:optional string propertyName =""(api.raw = "propertyName")
	20:optional string venueName =""(api.raw = "venueName")
	21:optional string placeName =""(api.raw = "placeName")
//    22:optional string coachName =""(api.raw = "coachName")
//	23:optional string memberName =""(api.raw = "memberName")
//	24:optional string memberProductName =""(api.raw = "memberProductName")
//	25:optional string memberProductPropertyName =""(api.raw = "memberProductPropertyName")
	26:optional list<ScheduleMemberInfo> memberCourseRecord ={}(api.raw = "memberCourseRecord")
	27:optional list<ScheduleCoachInfo> coachCourseRecord ={}(api.raw = "coachCourseRecord")
    25:list<list<base.Seat>> seats={}  (api.raw = "seats")
	251:optional string createdAt =""(api.raw = "createdAt")
	252:optional string updatedAt =""(api.raw = "updatedAt")
	28: string date=""  (api.raw = "date")
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
	9:optional string startAt ="" (api.raw = "startAt")
	10:optional string endAt ="" (api.raw = "endAt")
	11:optional string signStartAt ="" (api.raw = "signStartAt")
	12:optional string signEndAt ="" (api.raw = "signEndAt")
	13:optional i64 status =0 (api.raw = "status")
	14:optional i64 memberProductId =0 (api.raw = "memberProductId")
	15:optional i64 memberProductPropertyId =0 (api.raw = "memberProductPropertyId")
    25: base.Seat seat={}  (api.raw = "seat")
	16:optional string venueName ="" (api.raw = "venueName")
	17:optional string memberName ="" (api.raw = "memberName")
	18:optional string memberProductName ="" (api.raw = "memberProductName")
	19:optional string memberProductPropertyName ="" (api.raw = "memberProductPropertyName")
	20:optional string gender ="" (api.raw = "gender")
	21:optional i64 birthday =0 (api.raw = "birthday")
	22:optional string mobile ="" (api.raw = "mobile")
	251:optional string createdAt =""(api.raw = "createdAt")
    252:optional string updatedAt =""(api.raw = "updatedAt")
    23: string date=""  (api.raw = "date")
}

struct CreateOrUpdateScheduleReq {
    1:optional i64 id =0 (api.raw = "id")
    2:optional string type ="" (api.raw = "type")
    3:optional i64 propertyId =0 (api.raw = "propertyId")
    4:optional i64 venueId =0 (api.raw = "venueId")
    5:optional i64 placeId  =0(api.raw = "placeId")
    6:optional i64 num =0 (api.raw = "num")
    7:optional string startAt=""  (api.raw = "startAt")
    8:optional double price =0 (api.raw = "price")
    9:optional string remark=""  (api.raw = "remark")
    10:optional i64 coachId =0 (api.raw = "coachId")
    11:optional i64 memberId =0 (api.raw = "memberId")
    12:optional i64 memberProductId =0 (api.raw = "memberProductId")
    13:optional i64 memberProductPropertyId =0 (api.raw = "memberProductPropertyId")
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

struct ScheduleListReq {
    3:  optional i64 memberId=0 (api.raw = "memberId")
    4:  optional list<i64> coachIds =""(api.raw = "coachIds")
    5:  optional list<i64> productIds =""(api.raw = "productIds")
    6:  optional i64 venueId=0 (api.raw = "venueId")
    7:  optional list<i64> propertyIds=0 (api.raw = "propertyIds")
    8:  optional string startAt ="" (api.raw = "startAt")
    9:  optional string type=""  (api.raw = "type")
	101:  optional i64 page = 0  (api.raw = "page")
    102:  optional i64 pageSize = 0  (api.raw = "pageSize")
}
struct ScheduleMemberListReq  {
	1:  optional string type=""      (api.raw = "type")
	2:  optional i64 memberId   =0  (api.raw = "memberId")
	3:  optional i64 scheduleId =0  (api.raw = "scheduleId")
	101:  optional i64 page = 0  (api.raw = "page")
    102:  optional i64 pageSize = 0  (api.raw = "pageSize")
}
struct ScheduleCoachListReq  {
    1:optional string type = "" (api.raw = "type")
    2:optional i64	coachId  = 0  (api.raw = "coachId")
    3:optional i64	scheduleId =0  (api.raw = "scheduleId")
	101:  optional i64 page = 0  (api.raw = "page")
    102:  optional i64 pageSize = 0  (api.raw = "pageSize")
}
struct ScheduleCoachInfo  {
	1:optional i64 id =0 (api.raw = "id")
	2:optional i64 coachId =0  (api.raw = "coachId")
	3:optional i64 venueId =0  (api.raw = "venueId")
	4:optional i64 placeId =0  (api.raw = "placeId")
	5:optional i64 propertyId =0  (api.raw = "propertyId")
	6:optional i64 scheduleId =0  (api.raw = "scheduleId")
	7:optional string type ="" (api.raw = "type")
	8:optional string date ="" (api.raw = "date")
	9:optional string startAt ="" (api.raw = "startAt")
	10:optional string endAt ="" (api.raw = "endAt")
	11:optional string signStartAt ="" (api.raw = "signStartAt")
	12:optional string signEndAt ="" (api.raw = "signEndAt")
	13:optional i64 status =0  (api.raw = "status")

	14:optional string scheduleName  ="" (api.raw = "scheduleName")
	15:optional string propertyName ="" (api.raw = "propertyName")
	16:optional string venueName ="" (api.raw = "venueName")
	17:optional string placeName ="" (api.raw = "placeName")
	18:optional string coachName ="" (api.raw = "coachName")
	19:optional string coachAvatar ="" (api.raw = "coachAvatar")

//	20:optional string mobile ="" (api.raw = "mobile")
//	21:optional string memberName ="" (api.raw = "memberName")
//	22:optional string memberAvatar ="" (api.raw = "memberAvatar")
//	23:optional string memberProductName ="" (api.raw = "memberProductName")
//	24:optional string memberProductPropertyName  ="" (api.raw = "memberProductPropertyName")
	25:optional string remark ="" (api.raw = "remark")
//	26:optional string mRemark ="" (api.raw = "mRemark")
	251:optional string createdAt =""(api.raw = "createdAt")
   	252:optional string updatedAt =""(api.raw = "updatedAt")
}



struct SubscribeByMember  {
	1:optional string  avatar =""  (api.raw = "avatar")
	2:optional string  mobile =""  (api.raw = "mobile")
	3:optional i64  memberId =0   (api.raw = "memberId")
	4:optional i64  memberProductID =0   (api.raw = "memberProductID")
	5:optional i64  memberProductPropertyId =0  (api.raw = "memberProductPropertyId")

	6:optional string  memberName =""  (api.raw = "memberName")
	7:optional string  memberProductName = ""  (api.raw = "memberProductName")
	8:optional string  memberProductPropertyName = ""  (api.raw = "memberProductPropertyName")
}



service ScheduleService {
    base.NilResponse CreateSchedule(1: CreateOrUpdateScheduleReq req)  (api.post = "/service/schedule/create")

    base.NilResponse UpdateSchedule(1: CreateOrUpdateScheduleReq req) (api.post = "/service/schedule/update")

    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/status")

    base.NilResponse ListSchedule(1: ScheduleListReq req )(api.post = "/service/schedule/list")

    base.NilResponse DateListSchedule(1: ScheduleListReq req )(api.post = "/service/schedule/date-list")

    base.NilResponse GetScheduleById(1: base.IDReq req) (api.post = "/service/schedule/info")

    base.NilResponse GetScheduleMemberList(1: ScheduleMemberListReq req) (api.post = "/service/schedule/schedule-member-list")

    base.NilResponse SearchSubscribeByMember(1: SearchSubscribeByMemberReq req) (api.post = "/service/schedule/search-subscribe-by-member")

    base.NilResponse MemberSubscribe(1: MemberSubscribeReq req) (api.post = "/service/schedule/member-subscribe")

    base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/schedule-member-status")

    base.NilResponse GetScheduleCoachList(1: ScheduleCoachListReq req) (api.post = "/service/schedule/schedule-coach-list")

    base.NilResponse UpdateCoachStatus(1: base.StatusCodeReq req) (api.post = "/service/schedule/schedule-coach-status")

}