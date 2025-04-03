// Code generated by hertz generator.

package sms

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	base "kcers/idl_gen/model/base"
	sms "kcers/idl_gen/model/sms"
)

// SmsInfo .
// @router /service/sms/info [POST]
func SmsInfo(ctx context.Context, c *app.RequestContext) {
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

// SmsSendList .
// @router /service/sms/send-list [POST]
func SmsSendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sms.SmsSendListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(base.NilResponse)

	c.JSON(consts.StatusOK, resp)
}
