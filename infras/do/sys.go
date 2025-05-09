package do

import "kcers/idl_gen/model/admin/sys"

type Sys interface {
	ProductList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error)
	PropertyList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error)
	PropertyType(req *sys.SysListReq) (list []*sys.SysList, total int64, err error)
	VenueList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error)
	MemberList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error)
	ContractList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error)
	StaffList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error)
	PlaceList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error)
	RoleList(req *sys.SysListReq) (list []*sys.SysList, total int64, err error)
}
