namespace go venue

include "../base/base.thrift"


struct VenueInfo {
    1:  optional i64 id=0 (api.raw = "id")
    3:  optional string name = "" (api.raw = "name")
    4:  optional string address = ""  (api.raw = "address")
    5:  optional string addressDetail  = "" (api.raw = "addressDetail")
    6:  optional string latitude  = "" (api.raw = "latitude")
    7:  optional string longitude  = "" (api.raw = "longitude")
    8:  optional string mobile = ""  (api.raw = "mobile")
    9:  optional string pic  = "" (api.raw = "pic")
    10:  optional string information  = "" (api.raw = "information")
    11:  optional i64 status =0 (api.raw = "status")
    12:  optional string CreatedAt  = "" (api.raw = "CreatedAt")
    13:  optional string UpdatedAt = ""  (api.raw = "UpdatedAt")
}

struct VenueListReq {
    1:  optional i64 page =0  (api.raw = "page")
    2:  optional i64 pageSize =0  (api.raw = "pageSize")
    3:  optional string name  = "" (api.raw = "name")
    11:  optional i64 status =0  (api.raw = "status")
}
struct VenuePlaceListReq {
    1:  optional i64 page =0  (api.raw = "page")
    2:  optional i64 pageSize =0  (api.raw = "pageSize")
    3:  optional string name = ""  (api.raw = "name")
    11:  optional i64 status =0  (api.raw = "status")
}

struct VenuePlaceInfo{
	1:  optional i64 id =0  (api.raw = "id")
	2:  optional string name = ""  (api.raw = "name")
	3:  optional string pic = ""  (api.raw = "pic")
	4:  optional i64 venueId =0  (api.raw = "venueId")
	5:  optional i64 status =0  (api.raw = "status")
	6:  optional string createdAt = ""  (api.raw = "createdAt")
	7:  optional string updatedAt = ""  (api.raw = "updatedAt")
}

struct CreateOrUpdateVenueReq {
    1: optional i64 id = 0 (api.raw = "id")
    2: optional string name = ""  (api.raw = "name")
    3: optional string address = ""  (api.raw = "address")
    4: optional string addressDetail = ""  (api.raw = "addressDetail")
    5: optional string latitude = ""  (api.raw = "latitude")
    6: optional string longitude  = "" (api.raw = "longitude")
    7: optional string mobile = ""  (api.raw = "mobile")
    8: optional string pic = ""  (api.raw = "pic")
    9: optional string information = ""  (api.raw = "information")
    10: optional i64 status =0  (api.raw = "status")
}

struct CreateOrUpdateVenuePlaceReq {
    1: optional i64 id  =0 (api.raw = "id")
    2: optional string name  = "" (api.raw = "name")
    3: optional i64 venueId =0  (api.raw = "venueId")
    4: optional string pic = ""  (api.raw = "pic")
    5: optional i64 status =0  (api.raw = "status")
}

service VenueService {

    // 添加
    base.NilResponse CreatePlace(1: CreateOrUpdateVenuePlaceReq req) (api.post = "/api/admin/place/create")
    // 编辑
    base.NilResponse UpdatePlace(1: CreateOrUpdateVenuePlaceReq req) (api.post = "/api/admin/place/update")
    // 删除
    base.NilResponse PlaceUpdateStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/place/status")
    // 列表
    base.NilResponse PlaceList(1: VenuePlaceListReq req) (api.post = "/api/admin/place/list")

    // 添加
    base.NilResponse Create(1: CreateOrUpdateVenueReq req) (api.post = "/api/admin/venue/create")
    // 编辑
    base.NilResponse Update(1: CreateOrUpdateVenueReq req) (api.post = "/api/admin/venue/update")
    base.NilResponse UpdateStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/venue/status")
    // 列表
    base.NilResponse List(1: VenueListReq req) (api.post = "/api/admin/venue/list")
}