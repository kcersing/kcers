// Code generated by hertz generator.

package venue

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "kcers/idl_gen/model/base"
	venue "kcers/idl_gen/model/venue"
)

// CreatePlace .
// @router /api/admin/place/create [POST]
func CreatePlace(ctx context.Context, c *app.RequestContext) {
	var err error
	var req venue.CreateOrUpdateVenuePlaceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdatePlace .
// @router /api/admin/place/update [POST]
func UpdatePlace(ctx context.Context, c *app.RequestContext) {
	var err error
	var req venue.CreateOrUpdateVenuePlaceReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// PlaceUpdateStatus .
// @router /api/admin/place/status [POST]
func PlaceUpdateStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// PlaceList .
// @router /api/admin/place/list [POST]
func PlaceList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req venue.VenuePlaceListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// Create .
// @router /api/admin/venue/create [POST]
func Create(ctx context.Context, c *app.RequestContext) {
	var err error
	var req venue.CreateOrUpdateVenueReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// Update .
// @router /api/admin/venue/update [POST]
func Update(ctx context.Context, c *app.RequestContext) {
	var err error
	var req venue.CreateOrUpdateVenueReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateStatus .
// @router /api/admin/venue/status [POST]
func UpdateStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.StatusCodeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// List .
// @router /api/admin/venue/list [POST]
func List(ctx context.Context, c *app.RequestContext) {
	var err error
	var req venue.VenueListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}
