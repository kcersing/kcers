package do

import "kcers/idl_gen/model/auth"

type Role interface {
	Create(req *auth.RoleInfo) error
	Update(req *auth.RoleInfo) error
	Delete(id *int64) error
	RoleInfoByID(ID *int64) (roleInfo *auth.RoleInfo, err error)
	List(req *auth.RoleListReq) (roleInfoList []*auth.RoleInfo, total int, err error)
	UpdateStatus(ID *int64, status *int64) error
}
