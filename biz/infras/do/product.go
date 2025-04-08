package do

import "kcers/idl_gen/model/product"

type Product interface {
	CreateProperty(req *product.CreateOrUpdatePropertyReq) error
	UpdateProperty(req *product.CreateOrUpdatePropertyReq) error
	DeleteProperty(id int64) error
	UpdatePropertyStatus(id, status int64) error
	PropertyList(req *product.ListReq) (resp []*product.Property, total int, err error)

	Create(req *product.CreateOrUpdateReq) error
	Update(req *product.CreateOrUpdateReq) error
	Delete(id int64) error
	List(req *product.ListReq) (resp []*product.Product, total int, err error)
	UpdateStatus(id, status int64) error
	Info(id int64) (info *product.Product, err error)
}
