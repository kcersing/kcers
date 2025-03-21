namespace go sms

include "../base/base.thrift"


struct SmsSendListReq {
    1:  optional i64 page=1 (api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional string mobile="" (api.raw = "mobile")
}

struct SmsSend{
    1: optional string createdAt = ""  (api.raw = "createdAt")
    2: optional i64 status=0  (api.raw = "status")
    3:  optional string mobile="" (api.raw = "mobile")
    5:  optional string code="" (api.raw = "code")
    7: optional string bizId = ""  (api.raw = "bizId")
    /**通知类型[1会员;2员工]*/
    8: optional i64 notifyType = 1  (api.raw = "notifyType")
    9: optional string content = ""  (api.raw = "content")
    10: optional string templates = ""  (api.raw = "templates")
}

struct SmsInfo{
    /**通知短信数量*/
    1: optional i64 noticeCount=0  (api.raw = "noticeCount")
    /**已用通知*/
    2: optional i64 usedNotice=0  (api.raw = "usedNotice")

    16: optional string createdAt = ""  (api.raw = "createdAt")
    17: optional string updatedAt = "" (api.raw = "updatedAt")
}

service SmsService {

    /**短信信息*/
    base.NilResponse SmsInfo(1: base.IDReq req) (api.post = "/service/sms/info")
    /**发送记录*/
    base.NilResponse SmsSendList(1: SmsSendListReq req) (api.post = "/service/sms/send-list")

}