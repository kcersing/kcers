namespace go entry

include "../base/base.thrift"

service EntryService {
  base.NilResponse CreateEntry(1: CreateEntry req) (api.post = "/service/entry/create")
  base.NilResponse EntryList(1: EntryListReq req) (api.post = "/service/entry/list")
}
struct CreateEntry  {
    1:i64 memberProductId (api.raw = "memberProductId")
	3:string entryAt (api.raw = "entryAt")
	4:string leavingAt (api.raw = "leavingAt")
	5:i64 memberId (api.raw = "memberId")
	6:i64 userId (api.raw = "userId")
	7:i64 venueId (api.raw = "venueId")
}

struct EntryInfo  {
	1:i64 id
	2:i64 memberProductId
	4:string entryAt
	5:string leavingAt
	6:i64 memberId
	7:i64 userId
	8:i64 venueId
	9:string createdAt
	10:string updatedAt
	11:string venueName
	12:string userName
	13:string memberName
	14:string memberProductName
	15:string type
}


struct EntryListReq{
    1: optional i64 page=1 (api.raw = "page")
    2: optional i64 pageSize=100 (api.raw = "pageSize")
    3: optional i64 memberId = 0 (api.raw = "memberId")
    4: optional i64 venueId= 0 (api.raw = "venueId")
    5: optional i64 memberProductId= 0 (api.raw = "memberProductId")
    10: optional i64 memberPropertyId= 0 (api.raw = "memberPropertyId")
    7: optional i64 userId= 0 (api.raw = "userId")
    8: optional string entryAt="" (api.raw ="entryAt")
    9: optional string leavingAt="" (api.raw ="leavingAt")
}


