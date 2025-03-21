package do

import (
	"kcers/idl_gen/model/base"
	"kcers/idl_gen/model/menu"
)

type Api interface {
	Create(req *menu.ApiInfo) error
	Update(req *menu.ApiInfo) error
	Delete(id *int64) error
	List(req *menu.ListApiReq) (resp []*menu.ApiInfo, total int, err error)
	ApiTree(req *menu.ListApiReq) (resp []*base.Tree, total int, err error)
}
