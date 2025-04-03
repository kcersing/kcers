package do

import "kcers/idl_gen/model/banner"

type Banner interface {
	Create(req *banner.BannerInfo) error
	Update(req *banner.BannerInfo) error
	Delete(id int64) error
	List(req *banner.BannerListReq) (resp []*banner.BannerInfo, total int, err error)
	UpdateShow(ID, status int64) error
}
