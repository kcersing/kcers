package member

import (
	"entgo.io/ent/dialect/sql"
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent"
	"kcers/biz/dal/db/mysql/ent/memberproduct"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"kcers/biz/dal/enums"
	"kcers/biz/infras/service/venue"
	"kcers/idl_gen/model/member"
	"time"
)

func (m *Member) ProductSearch(members []int64) (info *member.MemberProductInfo, err error) {
	//m.db.MemberProduct.Query().Where(memberproduct.MemberIDIn(members...)).All(m.ctx)
	m.db.Debug().MemberProduct.Query().
		Modify(func(s *sql.Selector) {
			s.GroupBy(memberproduct.FieldProductID)
			s.Having(
				sql.In(
					memberproduct.FieldMemberID,
					members,
				),
			)
		}).
		All(m.ctx)
	return
}

func (m *Member) ProductList(req *member.MemberProductListReq) (resp []*member.MemberProductInfo, total int, err error) {
	var predicates []predicate.MemberProduct
	if req.MemberId > 0 {
		predicates = append(predicates, memberproduct.MemberID(req.MemberId))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, memberproduct.VenueID(req.VenueId))
	}
	if req.Name != "" {
		predicates = append(predicates, memberproduct.Name(req.Name))
	}
	if req.Status > 0 {
		predicates = append(predicates, memberproduct.Status(req.Status))
	}

	lists, err := m.db.MemberProduct.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member Product list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, m.entMemberProductInfo(v))
	}

	total, _ = m.db.MemberProduct.Query().Where(predicates...).Count(m.ctx)
	return
}

func (m *Member) ProductDetail(id int64) (info *member.MemberProductInfo, err error) {
	first, err := m.db.MemberProduct.Query().Where(memberproduct.IDEQ(id)).First(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member Product failed")
		return info, err
	}
	info = m.entMemberProductInfo(first)
	return info, err
}

func (m *Member) ProductUpdate(req *member.MemberProductInfo) error {
	//TODO implement me
	panic("implement me")
}

func (m *Member) ProductUpdateStatus(id, status int64) error {
	_, err := m.db.MemberProduct.Update().Where(memberproduct.IDEQ(id)).SetStatus(int64(status)).Save(m.ctx)
	return err
}

func (m *Member) entMemberProductInfo(req *ent.MemberProduct) (info *member.MemberProductInfo) {
	var statusName = enums.MPStatusNames[enums.MPStatus(req.Status)]

	info = &member.MemberProductInfo{
		ID:         req.ID,
		Name:       req.Name,
		Price:      req.Price,
		Status:     req.Status,
		StatusName: statusName,
		CreatedAt:  req.CreatedAt.Format(time.DateTime),
		Sn:         req.Sn,
		MemberId:   req.MemberID,
		OrderId:    req.OrderID,
		VenueId:    req.VenueID,
		ProductId:  req.ProductID,
	}
	vInfo, err := venue.NewVenue(m.ctx, m.c).VenueInfo(req.VenueID)
	if err == nil {
		info.VenueName = vInfo.Name
	}
	return
}
