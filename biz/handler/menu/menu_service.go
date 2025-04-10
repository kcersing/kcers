// Code generated by hertz generator.

package menu

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "kcers/idl_gen/model/base"
	menu "kcers/idl_gen/model/menu"
)

// MenuAuth .
// @router /service/auth/menu/role [POST]
func MenuAuth(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MenuRole .
// @router /service/menu/role [POST]
func MenuRole(ctx context.Context, c *app.RequestContext) {
	var err error

	hlog.Info(err)
	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CreateApi .
// @router /service/api/create [POST]
func CreateApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.ApiInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateApi .
// @router /service/api/update [POST]
func UpdateApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.ApiInfo
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteApi .
// @router /service/api [POST]
func DeleteApi(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// ApiList .
// @router /service/api/list [POST]
func ApiList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.ApiPageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// ApiTree .
// @router /service/api/tree [POST]
func ApiTree(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.ApiPageReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// CreateMenu .
// @router /service/menu/create [POST]
func CreateMenu(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.CreateOrUpdateMenuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateMenu .
// @router /service/menu/update [POST]
func UpdateMenu(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.CreateOrUpdateMenuReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// DeleteMenu .
// @router /service/menu [POST]
func DeleteMenu(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MenuLists .
// @router /service/menu/list [POST]
func MenuLists(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.MenuListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MenuTree .
// @router /service/menu/tree [POST]
func MenuTree(ctx context.Context, c *app.RequestContext) {
	var err error
	var req menu.MenuListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MenuInfo .
// @router /service/menu/info [POST]
func MenuInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.IDReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}
