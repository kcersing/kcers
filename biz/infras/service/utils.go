package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
	"kcers/biz/dal/db/mysql/ent"
	"kcers/biz/dal/db/mysql/ent/member"
	"kcers/biz/dal/db/mysql/ent/memberproduct"
	"kcers/biz/dal/db/mysql/ent/user"
	"kcers/biz/dal/db/mysql/ent/venue"

	product2 "kcers/biz/dal/db/mysql/ent/product"

	"strconv"
)

func GetTokenUserName(ctx context.Context, c *app.RequestContext, db *ent.Client) string {
	userEnt, _ := GetTokenUser(ctx, c, db)
	if userEnt != nil {
		return userEnt.Name
	}
	return ""
}

func GetTokenUserID(c *app.RequestContext) int64 {
	id, exist := c.Get("userId")
	if exist || id != nil {
		uId, ok := id.(string)
		if ok {
			uid, err := strconv.ParseInt(uId, 10, 64)
			if err != nil {
				return 0
			}
			return uid
		}
		return 0
	}
	return 0
}

func GetTokenUser(ctx context.Context, c *app.RequestContext, db *ent.Client) (one *ent.User, err error) {
	id := GetTokenUserID(c)
	if id >= 0 {
		one, err = GetUser(db, id)
		if err != nil {
			return nil, err
		}
		return one, nil
	}
	return nil, err
}

func GetTokenMemberID(c *app.RequestContext) int64 {
	id, exist := c.Get("memberId")
	if exist || id != nil {
		uId, ok := id.(string)
		if ok {
			uid, err := strconv.ParseInt(uId, 10, 64)
			if err != nil {
				return 0
			}
			return uid
		}
		return 0
	}
	return 0
}

func GetTokenMember(ctx context.Context, c *app.RequestContext, db *ent.Client) (one *ent.Member, err error) {

	id := GetTokenMemberID(c)
	if id >= 0 {
		one, err = GetMember(db, id)
		if err != nil {
			return nil, err
		}
		return one, nil
	}
	return nil, err
}

func GetMember(db *ent.Client, id int64) (one *ent.Member, err error) {
	one, err = db.Member.Query().Where(member.IDEQ(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该会员")
	}
	return one, nil

}
func GetMemberProduct(db *ent.Client, id int64) (one *ent.MemberProduct, err error) {
	one, err = db.MemberProduct.Query().Where(memberproduct.ID(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该会员产品")
	}
	return one, err
}

func GetUser(db *ent.Client, id int64) (one *ent.User, err error) {
	one, err = db.User.Query().Where(user.ID(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该员工")
	}
	return one, err
}

func GetVenue(db *ent.Client, id int64) (one *ent.Venue, err error) {
	one, err = db.Venue.Query().Where(venue.ID(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该场馆")
	}
	return one, err
}

func GetProduct(db *ent.Client, id int64) (one *ent.Product, err error) {
	one, err = db.Product.Query().Where(product2.ID(id)).First(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "未查询到该产品")
	}
	return one, err
}
func DeductNumberSurplus(db *ent.Client, id int64) (err error) {
	first, err := GetMemberProduct(db, id)
	if err != nil {
		return err
	}
	if (first.NumberSurplus - 1) < 0 {
		hlog.Error("剩余课耗不足")
		return errors.New("剩余课耗不足")
	}
	_, err = db.MemberProduct.UpdateOneID(id).
		AddNumberSurplus(first.NumberSurplus - 1).
		Save(context.Background())
	if err != nil {
		hlog.Errorf("扣除会员产品剩余数量失败ID：%v", id)
		return err
	}

	return nil

}
