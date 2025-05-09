namespace go admin.role

include "../base/base.thrift"


//创建或更新角色信息参数
struct RoleInfo {
    1:  i64 id (api.raw = "id")
    2:  string name (api.raw = "name")
    3:  string value (api.raw = "value")
    4:  string defaultRouter (api.raw = "defaultRouter")
    5:  i64 status (api.raw = "status")
    6:  string remark  (api.raw = "remark")
    7:  i64 orderNo (api.raw = "orderNo")
    8:  string createdAt (api.raw = "createdAt")
    9:  string updatedAt (api.raw = "updatedAt")
}

// role service
service RoleService {

  // Create role information | 创建角色
  base.NilResponse CreateRole(1: RoleInfo req) (api.post = "/api/admin/role/create")

  // Update role information | 更新角色
  base.NilResponse UpdateRole(1: RoleInfo req) (api.post = "/api/admin/role/update")

  // Delete role information | 删除角色信息
  base.NilResponse DeleteRole(1: base.IdReq req) (api.post = "/api/admin/role/del")

  // Get role information | 获取角色信息
  base.NilResponse RoleById(1: base.IdReq req) (api.get = "/api/admin/role")

  // Get role list | 获取角色列表
  base.NilResponse RoleList(1: base.PageInfoReq req) (api.get = "/api/admin/role/list")

  // Set role status | 设置角色状态, 启用1/禁用0
  base.NilResponse UpdateRoleStatus(1: base.StatusCodeReq req) (api.post = "/api/admin/role/status")
}

// authorization message
// API授权数据
struct ApiAuthorityInfo {
    1:  string path (api.raw = "path")
    2:  string method (api.raw = "method")
}

// 创建或更新API授权信息
struct CreateOrUpdateApiAuthorityReq {
    1:  i64 roleId (api.raw = "roleId")
//    2:  ApiAuthorityInfo data (api.raw = "api_authority_info")
    2:  list<i64> apis (api.raw = "apis")
}

// 菜单授权请求数据
struct MenuAuthorityInfoReq {
    1: i64 roleId (api.raw = "roleId")
    2: list<i64> menu_ids (api.raw = "menu_ids")
}
type ApiAuthInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

// authorization service
service authorityService {

  // 创建API权限
  base.NilResponse CreateAuthority(1: CreateOrUpdateApiAuthorityReq req) (api.post = "/api/admin/authority/api/create")

  // 更新API权限
  base.NilResponse UpdateApiAuthority(1: CreateOrUpdateApiAuthorityReq req) (api.post = "/api/admin/authority/api/update")

  // 获取角色api权限列表
  base.NilResponse ApiAuthority(1: base.IdReq req) (api.post = "/api/admin/authority/api/role")

  // 创建菜单权限
  base.NilResponse CreateMenuAuthority(1: MenuAuthorityInfoReq req) (api.post = "/api/admin/authority/menu/create")

  // 更新菜单权限
  base.NilResponse UpdateMenuAuthority(1: MenuAuthorityInfoReq req) (api.post = "/api/admin/authority/menu/update")

  // 获取角色菜单权限列表
  base.NilResponse MenuAuthority(1: base.IdReq req) (api.post = "/api/admin/authority/menu/role")

}

// api message
// API信息
struct ApiInfo {
    1:  i64 id (api.raw = "id")
    2:  string createdAt (api.raw = "createdAt")
    3:  string updatedAt (api.raw = "updatedAt")
    4:  string path (api.raw = "path")
    5:  string description (api.raw = "description")
    6:  string group (api.raw = "group")
    7:  string method (api.raw = "method")
}

// API列表请求数据
struct ApiPageReq {
    1:  i64 page (api.raw = "page")
    2:  i64 pageSize (api.raw = "pageSize")
    3:  string path (api.raw = "path")
    4:  string description (api.raw = "description")
    5:  string method (api.raw = "method")
    6:  string group (api.raw = "group")
}


// api service
service ApisService {
  // 创建或API
  base.NilResponse CreateApi(1: ApiInfo req) (api.post = "/api/admin/api/create")

  // 更新API
  base.NilResponse UpdateApi(1: ApiInfo req) (api.post = "/api/admin/api/update")

  // 删除API信息
  base.NilResponse DeleteApi(1: base.IdReq req) (api.post = "/api/admin/api")

  // 获取API列表
  base.NilResponse ApiList(1: ApiPageReq req) (api.post = "/api/admin/api/list")

  base.NilResponse ApiTree(1: ApiPageReq req) (api.post = "/api/admin/api/tree")

}