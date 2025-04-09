namespace go member

include "../base/base.thrift"

// Create or update user information request | 创建或更新用户信息
struct CreateOrUpdateMemberReq {
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional string name="" (api.raw = "name")
    3:  optional string username="" (api.raw = "username")
    4:  optional string mobile="" (api.raw = "mobile")
    5:  optional string gender="未知" (api.raw = "gender")
    6:  optional string birthday="" (api.raw = "birthday")
    7:  optional i64 source=0 (api.raw = "source")
    8:  optional i64 grade=0 (api.raw = "grade")
    9: optional i64  intention=0 (api.raw = "intention")
    10:  optional string email="" (api.raw = "email")
    11:  optional string wecom="" (api.raw = "wecom")

}

struct CreateOrUpdateMemberAvatarReq {
    1:  optional i64 memberId=0 (api.raw = "memberId")
    2:  optional string avatar="" (api.raw = "avatar")
}

struct MemberInfo {
     1:optional i64 id=0  (api.raw = "id")
     2:optional	string name ="" (api.raw = "name")

     4:optional i64 status=0 (api.raw = "status")
     5:optional	string username = "" (api.raw = "username")
     6:optional	string mobile ="" (api.raw = "mobile")
     7:optional	string avatar ="" (api.raw = "avatar")
     8:optional string createdAt ="" (api.raw = "createdAt")
     9:optional string updatedAt  =""(api.raw = "updatedAt")
     10:optional i64 condition=0 (api.raw = "condition")
     11:optional string conditionName  =""(api.raw = "conditionName")
     12:optional i64 createdId=0 (api.raw = "createdId")
     13:optional string createdName  =""(api.raw = "createdName")


     250:optional MemberProfile profile={} (api.raw = "profile")
     251:optional MemberPrivacy privacy={} (api.raw = "privacy")
     252:optional MemberDetail detail={} (api.raw = "detail")
}
struct MemberDetail{
	 /**消费总金额*/
     2:optional double moneySum=0 (api.raw = "moneySum")
     /**首次的产品*/
     3:optional i64 productId =0  (api.raw = "productId")
     4:optional string productName =""(api.raw = "productName")
     /**进馆总次数*/
     7:optional i64 entrySum=0 (api.raw = "entrySum")
     /**最后一次进馆时间*/
     8:optional string entryLastAt="" (api.raw = "entryLastAt")
     /**进馆最后期限时间*/
     9:optional string entryDeadlineAt =""(api.raw = "entryDeadlineAt")
     /**最后一次上课时间*/
     10:optional string classLastAt="" (api.raw = "classLastAt")

}
struct MemberProfile{
    1:  optional string name="" (api.raw = "name")
    2:  optional string fatherName="" (api.raw = "fatherName")
    3:  optional string motherName="" (api.raw = "motherName")
    4:  optional i64 grade=0 (api.raw = "grade")
    5:  optional i64 intention=0 (api.raw = "intention")
    6:  optional i64 source=0 (api.raw = "source")
    7:  optional string gradeName="" (api.raw = "gradeName")
    8:  optional string intentionName="" (api.raw = "intentionName")
    9:  optional string sourceName="" (api.raw = "sourceName")

    10:	optional string email="" (api.raw = "email")
    11:	optional string gender = "" (api.raw = "gender")
    12: optional i64 age  =0 (api.raw = "age")
    13: optional string wecom ="" (api.raw = "wecom")
    14: optional string birthday=""   (api.raw = "birthday")
    /**关联会员*/
    15:optional i64 relationMid=0   (api.raw = "relationMid")
    16:optional string relationMname="" (api.raw = "relationMname")
    /**关联员工*/
    17:optional i64 relationUid =0  (api.raw = "relationUid")
    18:optional string relationUname="" (api.raw = "relationUname")

}

struct MemberPrivacy {
	//身份证号
     1: string identityCard (api.raw = "identityCard")
        //正面
     2: string faceIdentityCard (api.raw = "faceIdentityCard")
        //反面
     3: string backIdentityCard (api.raw = "backIdentityCard")
        //人脸照片
     4: string facePic (api.raw = "facePic")
        //特征值
     5: string faceEigenvalue (api.raw = "faceEigenvalue")
        //人脸更新时间
     6: string faceUpdateAt (api.raw = "faceUpdateAt")
     7: string facePicUpdatedAt (api.raw = "facePicUpdatedAt")
}

// Get user list request | 获取用户列表请求参数
struct MemberListReq {
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional string name = "" (api.raw = "name")
    4:  optional string mobile = "" (api.raw = "mobile")
    5:  optional i64 source = 0 (api.raw = "source")
    6:  optional i64 intention = 0 (api.raw = "intention")
    7:  optional i64 createdId = 0 (api.raw = "createdId")
    8:  optional string startCreatedAt = "" (api.raw = "startCreatedAt")
    9:  optional string endCreatedAt = "" (api.raw = "endCreatedAt")
}

struct MemberSearchReq {
    1:  string value="" (api.raw = "value")
    2:  string option="" (api.raw = "option")
}

struct MemberProductListReq {
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=10 (api.raw = "pageSize")
    3:  optional i64 memberId=0 (api.raw = "memberId")
    4:  optional string name="" (api.raw = "name")
    5:  optional i64 venueId=0 (api.raw = "venueId")
    6:  optional i64 status=0 (api.raw = "status")
   
}

struct MemberPropertyListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=10 (api.raw = "pageSize")
    3:  optional i64 memberId=0 (api.raw = "memberId")
    4:  optional string type="" (api.raw = "type")
    5:  optional string name="" (api.raw = "name")
    6:  optional i64 venueId=0 (api.raw = "venueId")
    7:  optional i64 status=0 (api.raw = "status")
    8:  optional i64 memberProductId=0 (api.raw = "memberProductId")
}

struct MemberContractListReq{
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional i64 memberId=0 (api.raw = "memberId")
    4:  optional i64 venueId =0 (api.raw = "venueId")
    5:  optional i64 contractId =0 (api.raw = "contractId")
}

struct MemberContractInfo{
    1:  optional string	name =""(api.raw = "name")
    2:  optional i64 memberId =0 (api.raw = "memberId")
    3:  optional string	memberName="" (api.raw = "memberName")
    4:  optional i64 venueId =0(api.raw = "venueId")
    5:  optional string	venueName="" (api.raw = "venueName")
    6:  optional i64 memberProductId =0(api.raw = "memberProductId")
    7:  optional string	memberProductName =""(api.raw = "memberProductName")
    8:  optional i64 contractId =0(api.raw = "contractId")
    10:  optional string sign =""(api.raw = "sign")
    11:  optional string signImg="" (api.raw = "signImg")
    12:  optional string content="" (api.raw = "content")
    13:optional string createdAt ="" (api.raw = "createdAt")
    14:optional string updatedAt  =""(api.raw = "updatedAt")
}


struct MemberProductSearchReq{
    1:  optional list<i64> members (api.raw = "members")
}

struct MemberPropertySearchReq{
    1:  optional list<i64> memberProductIds (api.raw = "memberProductIds")
}

struct UpdateMemberFollowReq{
    1:  optional list<i64> memberId=0 (api.raw = "memberId")
    2:  optional i64 followId=0 (api.raw = "followId")
}

struct MemberProductInfo  {
	1: optional i64 id=0 (api.raw ="id")
	2: optional string name="" (api.raw ="name")
	3: optional double price=0 (api.raw ="price")
	4: optional i64 status=0 (api.raw ="status")
	5: optional string statusName="" (api.raw ="statusName")
	6: optional string createdAt="" (api.raw ="createdAt")
	7: optional string sn="" (api.raw ="sn")
	8: optional i64 memberId=0 (api.raw ="memberId")
	9: optional i64 orderId=0 (api.raw ="orderId")
	10: optional i64 venueId=0 (api.raw ="venueId")
	11: optional string venueName="" (api.raw ="venueName")
	12: optional i64 productId=0 (api.raw ="productId")
	13: optional list<MemberPropertyInfo> propertys=0 (api.raw ="propertys")
}



struct MemberPropertyInfo  {
	1: optional i64 id=0 (api.raw ="id")
	2: optional string sn="" (api.raw ="sn")
	3: optional i64 memberId=0 (api.raw ="memberId")
	4: optional i64 memberProductId=0 (api.raw ="memberProductId")
	5: optional i64 propertyId=0 (api.raw ="propertyId")
	6: optional string name="" (api.raw ="name")
	7: optional double price=0 (api.raw ="price")
	8: optional i64 duration=0 (api.raw ="duration")
	9: optional i64 length=0 (api.raw ="length")
	10: optional i64 count=0 (api.raw ="count")
	11: optional i64 countSurplus=0 (api.raw ="countSurplus")
	12: optional string type="" (api.raw ="type")
	13: optional i64 status=0 (api.raw ="status")
	14: optional string createdAt="" (api.raw ="createdAt")
	15: optional string validityAt="" (api.raw ="validityAt")
	16: optional string cancelAt="" (api.raw ="cancelAt")
	17: optional list<base.List> venue={} (api.raw ="venue")
	18: optional list<i64> venueId=0 (api.raw ="venueId")
	19: optional string venues="" (api.raw ="venues")
}

struct MemberEntryInfo  {
	1: optional i64 id=0 (api.raw ="id")
	2: optional string createdAt="" (api.raw ="createdAt")
	3: optional i64 memberPropertyId=0 (api.raw ="memberPropertyId")
	4: optional string entryAt="" (api.raw ="entryAt")
	5: optional string leavingAt="" (api.raw ="leavingAt")
	6: optional i64 memberId=0 (api.raw ="memberId")
	7: optional i64 memberProductId=0 (api.raw ="memberProductId")
	8: optional i64 userId=0 (api.raw ="userId")
	9: optional i64 venueId=0 (api.raw ="venueId")
}
struct LoginReq{
     1:optional	string mobile ="" (api.raw = "mobile")
     2:optional	string smsCaptcha ="" (api.raw = "smsCaptcha")
}


service MemberService {

  // 新增用户
  base.NilResponse CreateMember(1: CreateOrUpdateMemberReq req) (api.post = "/service/member/create")

  // 更新用户
  base.NilResponse UpdateMember(1: CreateOrUpdateMemberReq req) (api.post = "/service/member/update")

  // 获取用户基本信息
  base.NilResponse MemberInfo(1: base.IDReq req) (api.post = "/service/member/info")

  // 获取用户列表
  base.NilResponse MemberList(1: MemberListReq req) (api.post = "/service/member/list")
  base.NilResponse MemberListExport(1: MemberListReq req) (api.post = "/service/member/list/export")

  // 更新用户状态
  base.NilResponse UpdateMemberStatus(1: base.StatusCodeReq req) (api.post = "/service/member/status")
  base.NilResponse DelMember(1: base.Ids req) (api.post = "/service/member/del")
   base.NilResponse UpdateMemberFollow(1: UpdateMemberFollowReq req) (api.post = "/service/member/update-follow")

  // 更新用户状态
  base.NilResponse MemberSearch(1: MemberSearchReq req) (api.post = "/service/member/search")

  base.NilResponse MemberProductList(1: MemberProductListReq req) (api.post = "/service/member/product-list")

  base.NilResponse MemberPropertyList(1: MemberPropertyListReq req) (api.post = "/service/member/property-list")

  base.NilResponse MemberProductDetail(1: base.IDReq req) (api.post = "/service/member/product-detail")

  base.NilResponse MemberPropertyDetail(1: base.IDReq req) (api.post = "/service/member/property-detail")

  base.NilResponse MemberPropertyUpdate(1: MemberPropertyListReq req) (api.post = "/service/member/property-update")

  base.NilResponse MemberContractList(1: MemberPropertyListReq req) (api.post = "/service/member/contract-list")


  base.NilResponse MemberProductSearch(1: MemberProductSearchReq req) (api.post = "/service/member/search-product")

  base.NilResponse MemberPropertySearch(1: MemberPropertySearchReq req) (api.post = "/service/member/search-property")

}
