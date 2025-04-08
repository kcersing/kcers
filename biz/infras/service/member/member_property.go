package member

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent"
	"kcers/biz/dal/db/mysql/ent/memberproductproperty"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"kcers/biz/dal/db/mysql/ent/venue"
	"kcers/idl_gen/model/member"
	"time"
)

func (m *Member) PropertyUpdate(req *member.MemberPropertyInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m *Member) PropertySearch(memberProducts []int64) (info *member.MemberPropertyInfo, err error) {
	m.db.MemberProductProperty.Query().Where(memberproductproperty.MemberProductIDIn(memberProducts...)).All(m.ctx)

	return
}
func (m *Member) PropertyList(req *member.MemberPropertyListReq) (resp []*member.MemberPropertyInfo, total int, err error) {
	var predicates []predicate.MemberProductProperty
	if req.MemberId > 0 {
		predicates = append(predicates, memberproductproperty.MemberID(req.MemberId))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, memberproductproperty.HasVenuesWith(venue.ID(req.VenueId)))
	}
	if req.MemberProductId > 0 {
		predicates = append(predicates, memberproductproperty.MemberProductID(req.MemberProductId))
	}
	if req.Type != "" {
		predicates = append(predicates, memberproductproperty.Type(req.Type))
	}
	if req.Name != "" {
		predicates = append(predicates, memberproductproperty.Name(req.Name))
	}
	if req.Status > 0 {
		predicates = append(predicates, memberproductproperty.Status(req.Status))
	}

	lists, err := m.db.MemberProductProperty.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member Product list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, m.entMemberPropertyInfo(v))
	}

	total, _ = m.db.MemberProductProperty.Query().Where(predicates...).Count(m.ctx)
	return
}

func (m *Member) PropertyDetail(id int64) (info *member.MemberPropertyInfo, err error) {
	first, err := m.db.MemberProductProperty.Query().Where(memberproductproperty.IDEQ(id)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member Product failed")
		return info, err
	}
	info = m.entMemberPropertyInfo(first)
	return info, err
}

func (m *Member) PropertyUpdateStatus(id, status int64) error {
	_, err := m.db.MemberProductProperty.Update().Where(memberproductproperty.IDEQ(id)).SetStatus(status).Save(m.ctx)
	return err
}

func (m *Member) entMemberPropertyInfo(req *ent.MemberProductProperty) (info *member.MemberPropertyInfo) {

	info = &member.MemberPropertyInfo{
		ID:              req.ID,
		Sn:              req.Sn,
		MemberId:        req.MemberID,
		MemberProductId: req.MemberProductID,
		PropertyId:      req.PropertyID,
		Name:            req.Name,
		Price:           req.Price,
		Duration:        req.Duration,
		Length:          req.Length,
		Count:           req.Count,
		CountSurplus:    req.CountSurplus,
		Type:            req.Type,
		Status:          req.Status,
		CreatedAt:       req.CreatedAt.Format(time.DateTime),
		ValidityAt:      req.ValidityAt.Format(time.DateTime),
		CancelAt:        req.CancelAt.Format(time.DateTime),
	}

	venues, err := req.QueryVenues().Select(venue.FieldID, venue.FieldName).All(m.ctx)
	if err == nil {
		var ven []*member.PropertyVenue
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
	return
}
