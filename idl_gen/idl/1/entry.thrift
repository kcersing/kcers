namespace go admin.entry

include "../../base/base.thrift"



struct EntryListReq{
    1:  optional i64 page (api.raw = "page")
    2:  optional i64 pageSize (api.raw = "pageSize")

    3:  optional i64 memberId (api.raw = "memberId")
    4:  optional i64 venueId (api.raw = "venueId")
    5:  optional i64 memberProductId (api.raw = "memberProductId")
    6:  optional i64 memberPropertyId (api.raw = "memberPropertyId")
    7:  optional i64 userId (api.raw = "userId")
    8:  optional string entryTime (api.raw = "entryTime")
    9:  optional string leavingTime (api.raw = "leavingTime")

}
service EntryService {
  base.NilResponse EntryList(1: EntryListReq req) (api.post = "/service/entry/list")
}