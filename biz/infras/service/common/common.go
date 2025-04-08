package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"kcers/biz/dal/cache"
	db "kcers/biz/dal/db/mysql"
	"kcers/biz/dal/db/mysql/ent"
	"kcers/biz/infras/do"
)

type Common struct {
	ctx   context.Context
	c     *app.RequestContext
	db    *ent.Client
	cache *ristretto.Cache
}

func NewCommon(ctx context.Context, c *app.RequestContext) do.Common {
	return &Common{
		ctx:   ctx,
		c:     c,
		db:    db.DB,
		cache: cache.Cache,
	}
}
