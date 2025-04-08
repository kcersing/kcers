package product

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"kcers/biz/dal/db/mysql/ent/productproperty"
	"kcers/biz/dal/db/mysql/ent/venue"
	"kcers/biz/dal/enums"
	"kcers/idl_gen/model/base"
	"kcers/idl_gen/model/product"
	"time"
)

func (p *Product) PropertyList(req *product.ListReq) (resp []*product.Property, total int, err error) {

	var predicates []predicate.ProductProperty
	if req.Name != "" {
		predicates = append(predicates, productproperty.NameEQ(req.Name))
	}
	if len(req.Status) > 0 {
		predicates = append(predicates, productproperty.StatusIn(req.Status...))
	}
	if len(req.CreatedAt) > 0 {
		s, _ := time.ParseInLocation(time.DateTime, req.CreatedAt[0], time.Local)
		d, _ := time.ParseInLocation(time.DateTime, req.CreatedAt[1], time.Local)
		predicates = append(predicates, productproperty.CreatedAtGTE(s))
		predicates = append(predicates, productproperty.CreatedAtLTE(d))
	}

	if len(req.VenueId) > 0 {
		predicates = append(predicates, productproperty.HasVenuesWith(venue.IDIn(req.VenueId...)))
	}

	if req.Type != "" {
		predicates = append(predicates, productproperty.TypeEQ(req.Type))
	}

	lists, err := p.db.ProductProperty.
		Query().
		Where(predicates...).
		Order(ent.Desc(productproperty.FieldID)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "get product list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Product info failed")
		return resp, 0, err
	}

	for i, v := range lists {

	}

	total, _ = p.db.ProductProperty.Query().Where(predicates...).Count(p.ctx)
	return

}

func (p *Product) entProperty(req *ent.ProductProperty) (info *product.Property) {

	info = &product.Property{
		ProductId:  req.ProductID,
		Name:       req.Name,
		Price:      req.Price,
		Duration:   req.Duration,
		Length:     req.Length,
		Count:      req.Count,
		Type:       req.Type,
		Data:       req.Data,
		ID:         req.ID,
		Status:     req.Status,
		StatusName: enums.ReturnPropertyValues(req.Status),
		Tags:       nil,
		Contracts:  nil,
	}

	venues, err := req.QueryVenues().Select(venue.FieldID, venue.FieldName).All(m.ctx)
	if err == nil {
		var ven []*product.PropertyVenue
		err = copier.Copy(&ven, &venues)
		info.Venue = ven
		for i2, v := range ven {
			if i2 == 0 {
				info.Venues = v.Name
			} else {
				info.Venues += ", " + v.Name
			}
			info.VenueId = append(info.VenueId, v.ID)
		}

	}

	venues, err := v.QueryVenues().Select(venue.FieldID, venue.FieldName).All(p.ctx)
	if err == nil {
		var ven []base.List
		err = copier.Copy(&ven, &venues)
		resp[i].Venue = ven

		for i2, v := range ven {
			if i2 == 0 {
				resp[i].Venues = v.Name
			} else {
				resp[i].Venues += ", " + v.Name
			}
			resp[i].VenueId = append(resp[i].VenueId, v.ID)
		}

	}

	return
}

func (p *Product) CreateProperty(req *product.CreateOrUpdatePropertyReq) error {

	_, err := p.db.ProductProperty.Create().
		SetName(req.Name).
		SetType(req.Type).
		SetLength(req.Length).
		SetStatus(req.Status).
		SetDuration(req.Duration).
		SetCount(req.Count).
		SetPrice(req.Price).
		SetData("").
		AddVenues(p.db.Venue.Query().Where(venue.IDIn(req.VenueId...)).AllX(p.ctx)...).
		Save(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "create user failed")
		return err
	}
	return nil

}

func (p *Product) UpdateProperty(req *product.CreateOrUpdatePropertyReq) error {

	_, err := p.db.ProductProperty.Update().
		Where(productproperty.IDEQ(req.ID)).
		SetName(req.Name).
		SetType(req.Type).
		SetLength(req.Length).
		SetStatus(req.Status).
		SetDuration(req.Duration).
		SetCount(req.Count).
		SetPrice(req.Price).
		SetData("").
		ClearVenues().
		AddVenues(p.db.Venue.Query().Where(venue.IDIn(req.VenueId...)).AllX(p.ctx)...).
		Save(p.ctx)

	if err != nil {
		err = errors.Wrap(err, "update Product Property failed")
		return err
	}

	return nil
}
func (p *Product) UpdatePropertyStatus(ID int64, status int64) error {
	_, err := p.db.ProductProperty.Update().Where(productproperty.IDEQ(ID)).SetStatus(int64(status)).Save(p.ctx)
	return err
}
func (p *Product) DeleteProperty(id int64) error {
	//TODO implement me
	panic("implement me")
}
