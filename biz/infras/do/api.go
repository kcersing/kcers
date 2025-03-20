package do

type Api interface {
	Create(req ApiInfo) error
	Update(req ApiInfo) error
	Delete(id int64) error
	List(req ListApiReq) (resp []*ApiInfo, total int, err error)
	ApiTree(req ListApiReq) (resp []*Tree, total int, err error)
}
