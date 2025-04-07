package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"kcers/biz/dal/cache"
	"kcers/biz/dal/config"
	db "kcers/biz/dal/db/mysql"
	"kcers/biz/dal/db/mysql/ent"
	banner2 "kcers/biz/dal/db/mysql/ent/banner"
	"kcers/biz/dal/db/mysql/ent/predicate"
	"kcers/biz/infras/do"
	"kcers/idl_gen/model/banner"
	"sync"
)

type Banner struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
	mu    sync.Mutex
}

var _ do.Banner = (*Banner)(nil)

func NewBanner(ctx context.Context, c *app.RequestContext) do.Banner {
	return &Banner{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}

func (b *Banner) Create(req *banner.BannerInfo) error {
	err := b.db.Banner.Create().
		SetName(req.Name).
		SetPic(req.Pic).
		SetLink(req.Link).
		SetIsShow(req.IsShow).
		Exec(b.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b *Banner) Update(req *banner.BannerInfo) error {
	err := b.db.Banner.Update().
		Where(banner2.IDEQ(req.ID)).
		SetName(req.Name).
		SetPic(req.Pic).
		SetLink(req.Link).
		SetIsShow(req.IsShow).
		Exec(b.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b *Banner) Delete(id int64) error {
	err := b.db.Banner.Update().
		Where(banner2.IDEQ(id)).
		SetDelete(1).
		Exec(b.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (b *Banner) List(req *banner.BannerListReq) (resp []*banner.BannerInfo, total int, err error) {
	var predicates []predicate.Banner

	predicates = append(predicates, banner2.Delete(0))

	lists, err := b.db.Banner.Query().Where(predicates...).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Order(ent.Desc(banner2.FieldID)).
		Limit(int(req.PageSize)).All(b.ctx)
	if err != nil {
		err = errors.Wrap(err, "get contest participant list failed")
		return resp, total, err
	}

	for _, v := range lists {
		resp = append(resp, entBannerInfo(v))
	}

	total, _ = b.db.Banner.Query().Where(predicates...).Count(b.ctx)
	return
}

func entBannerInfo(v *ent.Banner) *banner.BannerInfo {
	return &banner.BannerInfo{
		ID:     v.ID,
		Name:   v.Name,
		Pic:    v.Pic,
		Link:   v.Link,
		IsShow: v.IsShow,
	}
}

func (b *Banner) UpdateShow(ID int64, status int64) error {
	err := b.db.Banner.Update().
		Where(banner2.IDEQ(ID)).
		SetIsShow(status).
		Exec(b.ctx)
	if err != nil {
		return err
	}
	return nil
}
