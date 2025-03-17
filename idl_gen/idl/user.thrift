namespace go user

include "../base/base.thrift"

// 用户基本信息
struct UserInfo  {
  1:  optional i64 id =0 (api.raw = "id")
  2:  optional i64 status=0 (api.raw = "status")
  3:  optional string username =""(api.raw = "username")
  4:  optional string password="" (api.raw = "password")
  5:  optional string name="" (api.raw = "name")
  6:  optional string sideMode="" (api.raw = "sideMode")
  7:  optional string baseColor="" (api.raw = "baseColor")
  8:  optional string activeColor="" (api.raw = "activeColor")
  9:  optional i64 roleId=0 (api.raw = "roleId")
  10: optional string mobile =""(api.raw = "mobile")
  11: optional string email="" (api.raw = "email")
  12: optional string wecom =""(api.raw = "wecom")
  13: optional string avatar="" (api.raw = "avatar")
  14: optional string createdAt="" (api.raw = "createdAt")
  15: optional string updatedAt="" (api.raw = "updatedAt")
  16: optional string roleName =""(api.raw = "roleName")
  17: optional string roleValue =""(api.raw = "roleValue")
  18: optional string gender="" (api.raw = "gender")
  19: optional i64 age =0(api.raw = "age")
  20: optional string birthday="" (api.raw = "birthday")
  21: optional string job =""(api.raw = "job")
  22: optional string jobName =""(api.raw = "jobName")
  23: optional string organization =""(api.raw = "organization")
  24: optional string organizationName="" (api.raw = "organizationName")
  25: optional i64 defaultVenueId =0(api.raw = "defaultVenueId")
}

// 登录请求
struct UserLoginReq {
    1:  string username =""(api.raw = "username")
    2:  string password =""(api.raw = "password")
    3:  i64 captchaId=0 (api.raw = "captchaId")
    4:  string captcha =""(api.raw = "captcha")
}

// 注册请求
struct RegisterReq {
    1:  optional string username =""(api.raw = "username")
    2:  optional string password =""(api.raw = "password")
    3:  optional i64 captchaId =0(api.raw = "captchaId")
    4:  optional string captcha="" (api.raw = "captcha")
    5:  optional string email =""(api.raw = "email")
}

// 修改密码请求
struct ChangePasswordReq {
    1:  i64 userId=0 (api.raw = "userId")
    2:  string newPassword="" (api.raw = "newPassword")
}

// 创建或更新用户信息请求
struct CreateOrUpdateUserReq {
    1:  optional i64 id =0(api.raw = "id")
    2:  optional string avatar =""(api.raw = "avatar")
    4:  optional string mobile =""(api.raw = "mobile")
    5:  optional string email =""(api.raw = "email")
    6:  optional i64 status =0(api.raw = "status")
    7:  optional string name =""(api.raw = "name")
    8:  optional string gender="" (api.raw = "gender")
    9:  optional string wecom="" (api.raw = "wecom")
    10: optional i64 createId =0(api.raw = "createId")
    11: optional string birthday="" (api.raw = "birthday")
    251:  optional string createdAt = ""  (api.raw = "createdAt")
    252:  optional string updatedAt = ""  (api.raw = "updatedAt")
}

// 获取用户列表请求
struct UserListReq {
    1:  optional i64 page =1(api.raw = "page")
    2:  optional i64 pageSize=100 (api.raw = "pageSize")
    3:  optional string username="" (api.raw = "username")
    4:  optional string name =""(api.raw = "name")
    5:  optional string email="" (api.raw = "email")
    6:  optional string mobile="" (api.raw = "mobile")
    7:  optional i64 roleId =0(api.raw = "roleId")
}

// 设置用户角色
struct SetUserRole{
    1:  optional i64 id=0 (api.raw = "id")
    2:  optional i64 roleId=0 (api.raw = "roleId")
}

// 设置默认场馆
struct SetDefaultVenueReq{
  1:  optional i64 venueId =0(api.raw = "venueId")
}

service UserService {
  // 注册
  base.NilResponse Register(1: RegisterReq req) (api.post = "/api/register")

  // 获取用户权限码
  base.NilResponse UserPermCode(1: base.Empty req) (api.post = "/api/admin/user/perm")

  // 修改密码
  base.NilResponse ChangePassword(1: ChangePasswordReq req) (api.post = "/api/admin/user/change-password")

  // 新增用户
  base.NilResponse CreateUser(1: CreateOrUpdateUserReq req) (api.post = "/api/admin/user/create")

  // 更新用户
  base.NilResponse UpdateUser(1: CreateOrUpdateUserReq req) (api.post = "/api/admin/user/update")

  // 获取用户基本信息
  base.NilResponse UserInfo(1: base.Empty req) (api.get = "/api/admin/user/info")

  // 获取用户列表
  base.NilResponse UserList(1: UserListReq req) (api.post = "/api/admin/user/list")

  // 删除用户信息
  base.NilResponse DeleteUser(1: base.IdReq req) (api.post = "/api/admin/user")

  // 获取用户个人信息
  base.NilResponse UserProfile(1: base.Empty req) (api.get = "/api/admin/user/profile")

  // 更新用户状态
  base.NilResponse UpdateUserStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/user/status")

  // 设置用户角色
  base.NilResponse SetUserRole(1: SetUserRole req) (api.post = "/api/admin/user/set-role")

  // 设置默认场馆
  base.NilResponse SetDefaultVenue(1: SetDefaultVenueReq req) (api.post = "/api/admin/user/set-default-venue")
}