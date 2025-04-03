// Code generated by hertz generator.

package schedule

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	schedule "kcers/idl_gen/model/admin/schedule"
	base "kcers/idl_gen/model/base"
)

// CreateSchedule .
// @router /service/schedule/create [POST]
func CreateSchedule(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.CreateOrUpdateScheduleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateSchedule .
// @router /service/schedule/update [POST]
func UpdateSchedule(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.CreateOrUpdateScheduleReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateStatus .
// @router /service/schedule/status [POST]
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

// ListSchedule .
// @router /service/schedule/list [POST]
func ListSchedule(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.ScheduleListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// DateListSchedule .
// @router /service/schedule/date-list [POST]
func DateListSchedule(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.ScheduleListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetScheduleById .
// @router /service/schedule/info [POST]
func GetScheduleById(ctx context.Context, c *app.RequestContext) {
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

// GetScheduleMemberList .
// @router /service/schedule/schedule-member-list [POST]
func GetScheduleMemberList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.ScheduleMemberListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// SearchSubscribeByMember .
// @router /service/schedule/search-subscribe-by-member [POST]
func SearchSubscribeByMember(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.SearchSubscribeByMemberReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// MemberSubscribe .
// @router /service/schedule/member-subscribe [POST]
func MemberSubscribe(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.MemberSubscribeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateMemberStatus .
// @router /service/schedule/schedule-member-status [POST]
func UpdateMemberStatus(ctx context.Context, c *app.RequestContext) {
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

// GetScheduleCoachList .
// @router /service/schedule/schedule-coach-list [POST]
func GetScheduleCoachList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req schedule.ScheduleCoachListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}

// UpdateCoachStatus .
// @router /service/schedule/schedule-coach-status [POST]
func UpdateCoachStatus(ctx context.Context, c *app.RequestContext) {
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
