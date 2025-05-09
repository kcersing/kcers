namespace go admin.logs

include "../base/base.thrift"


//日志列表请求数据
struct LogsListReq {
    1:  i64 page (api.raw = "page")
    2:  i64 pageSize (api.raw = "pageSize")
    3:  string type (api.raw = "type")
    4:  string method (api.raw = "method")
    5:  string api (api.raw = "api")
    6:  string success (api.raw = "success")
    7:  string operators (api.raw = "operators")
}
//日志信息
struct LogsInfo {
    1:  string type (api.raw = "type")
    2:  string method (api.raw = "method")
    3:  string api (api.raw = "api")
    4:  string success (api.raw = "success")
    5:  string reqContent (api.raw = "reqContent")
    6:  string respContent (api.raw = "respContent")
    7:  string ip (api.raw = "ip")
    8:  string userAgent (api.raw = "userAgent")
    9:  string operator (api.raw = "operator")
    10:  string time (api.raw = "time")
    11:  string createdAt (api.raw = "createdAt")
    12:  string updatedAt (api.raw = "updatedAt")
}

service LogsService{
  // Get logs list | 获取日志列表
  base.NilResponse GetLogsList(1: LogsListReq req) (api.post = "/api/admin/logs/list")

  // Delete logs | 删除日志信息
  base.NilResponse DeleteLogs(1: base.Empty req) (api.post = "/api/admin/logs/deleteAll")

}