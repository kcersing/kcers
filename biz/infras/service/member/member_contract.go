package member

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent/membercontract"
	"kcers/biz/dal/db/mysql/ent/predicate"
	venueService "kcers/biz/infras/service/venue"
	"kcers/idl_gen/model/member"
)

func (m *Member) ContractList(req *member.MemberContractListReq) (resp []*member.MemberContractInfo, total int, err error) {
	var predicates []predicate.MemberContract
	if req.MemberId > 0 {
		predicates = append(predicates, membercontract.MemberID(req.MemberId))
	}
	if req.VenueId > 0 {
		predicates = append(predicates, membercontract.VenueID(req.VenueId))
	}
	if req.ContractId > 0 {
		predicates = append(predicates, membercontract.ContractID(req.ContractId))
	}

	lists, err := m.db.MemberContract.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(m.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Member Contract list failed")
		return resp, total, err
	}

	err = copier.Copy(&resp, &lists)
	if err != nil {
		err = errors.Wrap(err, "copy Member Contract info failed")
		return resp, 0, err
	}
	for i, v := range lists {
		vInfo, err := venueService.NewVenue(m.ctx, m.c).VenueInfo(v.VenueID)
		if err == nil {
			resp[i].VenueName = vInfo.Name
		}
		first, err := v.QueryContent().First(m.ctx)
		if err != nil {
			return nil, 0, err
		}
		resp[i].Content = first.Content
		resp[i].SignImg = first.SignImg
	}

	return
}
