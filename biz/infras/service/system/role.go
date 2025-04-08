package system

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
	"kcers/biz/dal/cache"
	"kcers/biz/dal/config"
	db "kcers/biz/dal/db/mysql"
	"kcers/biz/dal/db/mysql/ent"
	"kcers/biz/dal/db/mysql/ent/menu"
	"kcers/biz/dal/db/mysql/ent/role"
	"kcers/biz/dal/db/mysql/ent/user"
	"kcers/biz/infras/do"
	"kcers/biz/pkg/utils"
	"kcers/idl_gen/model/auth"
	"strconv"
	"time"
)

type Role struct {
	ctx   context.Context
	c     *app.RequestContext
	salt  string
	db    *ent.Client
	cache *ristretto.Cache
}

func (r *Role) Create(req *auth.RoleInfo) error {
	roleEnt, err := r.db.Role.Create().
		SetName(req.Name).
		SetValue(req.Value).
		SetDefaultRouter(req.DefaultRouter).
		SetStatus(req.Status).
		SetRemark(req.Remark).
		SetOrderNo(req.OrderNo).
		Save(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "create Role failed")
		return err
	}

	// set roleEnt to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(roleEnt.ID)), roleEnt, 1, 1*time.Hour)
	return nil
}

func (r *Role) Update(req *auth.RoleInfo) error {
	roleEnt, err := r.db.Role.UpdateOneID(req.ID).
		SetName(req.Name).
		SetValue(req.Value).
		SetDefaultRouter(req.DefaultRouter).
		SetStatus(req.Status).
		SetRemark(req.Remark).
		SetOrderNo(req.OrderNo).
		SetUpdatedAt(time.Now()).
		Save(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Role failed")
		return err
	}

	// set roleEnt to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(roleEnt.ID)), roleEnt, 1, 1*time.Hour)
	return nil
}

func (r *Role) Delete(id int64) error {
	// whether role is used by user
	exist, err := r.db.User.Query().Where(user.RoleIDEQ(id)).Exist(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "query user - role failed")
		return err
	}
	if exist {
		return errors.New("role is used by user")
	}
	// delete role from db
	err = r.db.Role.DeleteOneID(id).Exec(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "delete Role failed")
		return err
	}
	// delete role from cache
	r.cache.Del("roleData" + strconv.Itoa(int(id)))
	return nil
}

func (r *Role) RoleInfoById(id int64) (roleInfo *auth.RoleInfo, err error) {
	roleInterface, ok := r.cache.Get("roleData" + strconv.Itoa(int(id)))
	if ok {
		if l, ok := roleInterface.(*ent.Role); ok {
			return &auth.RoleInfo{
				ID:            l.ID,
				Name:          l.Name,
				Value:         l.Value,
				DefaultRouter: l.DefaultRouter,
				Status:        l.Status,
				Remark:        l.Remark,
				OrderNo:       l.OrderNo,
				CreatedAt:     l.CreatedAt.Format(time.DateTime),
				UpdatedAt:     l.UpdatedAt.Format(time.DateTime),
			}, nil
		}
	}
	// get role from db
	roleEnt, err := r.db.Role.Query().Where(role.IDEQ(id)).Only(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "get Role failed")
		return nil, err
	}
	// set role to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(id)), roleEnt, 1, 1*time.Hour)
	// convert to RoleInfo
	roleInfo = &auth.RoleInfo{
		ID:            roleEnt.ID,
		Name:          roleEnt.Name,
		Value:         roleEnt.Value,
		DefaultRouter: roleEnt.DefaultRouter,
		Status:        int64(roleEnt.Status),
		Remark:        roleEnt.Remark,
		OrderNo:       roleEnt.OrderNo,
		CreatedAt:     roleEnt.CreatedAt.Format(time.DateTime),
		UpdatedAt:     roleEnt.UpdatedAt.Format(time.DateTime),
	}
	return
}

func (r *Role) List(req *auth.RoleListReq) (roleInfoList []*auth.RoleInfo, total int, err error) {

	roleEntList, err := r.db.Role.Query().Order(ent.Asc(role.FieldOrderNo)).
		Offset(int(req.Page-1) * int(req.PageSize)).
		Limit(int(req.PageSize)).All(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "get RoleList failed")
		return nil, 0, err
	}
	// convert to List
	for _, roleEnt := range roleEntList {
		menuArr, _ := roleEnt.QueryMenus().GroupBy(menu.FieldID).Ints(r.ctx)

		roleInfoList = append(roleInfoList, &auth.RoleInfo{
			ID:            roleEnt.ID,
			Name:          roleEnt.Name,
			Value:         roleEnt.Value,
			DefaultRouter: roleEnt.DefaultRouter,
			Status:        roleEnt.Status,
			Remark:        roleEnt.Remark,
			OrderNo:       roleEnt.OrderNo,
			CreatedAt:     roleEnt.CreatedAt.Format(time.DateTime),
			UpdatedAt:     roleEnt.UpdatedAt.Format(time.DateTime),
			Menus:         utils.ConvertIntSliceToInt64Slice(menuArr),
			Apis:          utils.ConvertIntSliceToInt64Slice(roleEnt.Apis),
		})
	}
	total, _ = r.db.Role.Query().Count(r.ctx)
	return
}

func (r *Role) UpdateStatus(id, status int64) error {

	roleEnt, err := r.db.Role.UpdateOneID(id).SetStatus(status).Save(r.ctx)
	if err != nil {
		err = errors.Wrap(err, "update Role status failed")
		return err
	}
	// set role to cache
	r.cache.SetWithTTL("roleData"+strconv.Itoa(int(id)), roleEnt, 1, 1*time.Hour)

	return nil
}

func NewRole(ctx context.Context, c *app.RequestContext) do.Role {
	return &Role{
		ctx:   ctx,
		c:     c,
		salt:  config.GlobalServerConfig.MySQLInfo.Salt,
		db:    db.DB,
		cache: cache.Cache,
	}
}
