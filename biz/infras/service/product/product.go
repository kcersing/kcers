package product

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"kcers/biz/dal/cache"
	"kcers/biz/dal/config"
	db "kcers/biz/dal/db/mysql"
	"kcers/biz/dal/db/mysql/ent"
	"kcers/biz/dal/db/mysql/ent/predicate"
	product2 "kcers/biz/dal/db/mysql/ent/product"
	"kcers/biz/dal/db/mysql/ent/productproperty"
	"kcers/biz/dal/db/mysql/ent/user"
	"kcers/biz/dal/enums"
	"kcers/biz/infras/do"
	"kcers/idl_gen/model/product"
	"strconv"
	"time"
)

type Product struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (p *Product) Info(id int64) (info *product.Product, err error) {
	inter, exist := p.cache.Get("productInfo" + strconv.Itoa(int(id)))
	if exist {
		if v, ok := inter.(*product.Product); ok {
			return v, nil
		}
	}
	one, err := p.db.Product.Query().Where(product2.IDEQ(id)).First(p.ctx)
	if err != nil {
		err = errors.Wrap(err, "get product failed")
		return info, err
	}

	info = p.entProduct(one)

	p.cache.SetWithTTL("productInfo"+strconv.Itoa(int(info.ID)), &info, 1, 1*time.Hour)
	return
}

func (p *Product) Create(req *product.CreateOrUpdateReq) error {

	properties, err := p.db.ProductProperty.Query().Where(
		productproperty.IDIn(req.Propertys...),
	).All(p.ctx)
	if err != nil {
		return errors.Wrap(err, "未找到属性")
	}
	_, err = p.db.Product.Create().
		SetName(req.Name).
		SetPic(req.Pic).
		SetDescription(req.Description).
		SetPrice(req.Price).
		SetStock(req.Stock).
		AddPropertys(properties...).
		SetCreateID(req.CreateId).
		Save(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "create Product failed")
		return err
	}

	return nil
}

func (p *Product) Update(req *product.CreateOrUpdateReq) error {
	_, err := p.db.Product.Update().
		Where(product2.IDEQ(req.ID)).
		SetName(req.Name).
		SetPic(req.Pic).
		SetDescription(req.Description).
		SetStatus(req.Status).
		SetPrice(req.Price).
		SetStock(req.Stock).
		SetCreateID(req.CreateId).
		Save(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "update Product failed")
		return err
	}

	return nil
}

func (p *Product) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (p *Product) List(req *product.ListReq) (resp []*product.Product, total int, err error) {
	var predicates []predicate.Product
	if req.Name != "" {
		predicates = append(predicates, product2.NameEQ(req.Name))
	}
	lists, err := p.db.Product.Query().Where(predicates...).
		Order(ent.Desc(product2.FieldCreatedAt)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(p.ctx)
	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy product info failed")
		return resp, 0, err
	}

	for i, v := range lists {
		//propertys

		ds, _ := v.QueryPropertys().IDs(p.ctx)

		all, err := p.db.ProductProperty.Query().Where(productproperty.IDIn(ds...)).All(p.ctx)
		if err != nil {
			return nil, 0, err
		}

		for _, p := range all {
			pi := product.Property{
				ID:       p.ID,
				Name:     p.Name,
				Price:    p.Price,
				Duration: p.Duration,
				Length:   p.Length,
				Count:    p.Count,
				Type:     p.Type,
				Data:     p.Data,
				Status:   p.Status,
			}
			resp[i].Propertys = append(resp[i].Propertys, &pi)
		}
	}

	total, _ = p.db.Product.Query().Where(predicates...).Count(p.ctx)
	return

}

func (p *Product) UpdateStatus(ID int64, status int64) error {
	_, err := p.db.Product.Update().Where(product2.IDEQ(ID)).SetStatus(int64(status)).Save(p.ctx)
	return err
}

func (p *Product) entProduct(req *ent.Product) (info *product.Product) {
	var statusName = enums.ReturnProductValues(req.Status)

	info = &product.Product{
		ID:          req.ID,
		Name:        req.Name,
		Pic:         req.Pic,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Status:      req.Status,
		CreatedId:   req.CreateID,

		Propertys:  nil,
		StatusName: statusName,

		IsSales:     req.IsSales,
		SignSalesAt: req.SignSalesAt.Format(time.DateTime),
		EndSalesAt:  req.EndSalesAt.Format(time.DateTime),
		CreatedAt:   req.CreatedAt.Format(time.DateTime),
		UpdatedAt:   req.UpdatedAt.Format(time.DateTime),
	}

	if req.CreateID != 0 {
		create, err := p.db.User.Query().Where(user.IDEQ(req.CreateID)).First(p.ctx)
		if err != nil {
			err = errors.Wrap(err, "get user failed")
			return info
		}
		info.CreatedName = create.Name
	}
	return
}

func NewProduct(ctx context.Context, c *app.RequestContext) do.Product {
	return &Product{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
