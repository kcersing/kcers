package mw

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"kcers/idl_gen/model/logs"
	"strconv"
	"time"
)

func LogMw() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()
		c.Next(ctx)
		var log logs.LogsInfo
		log.Type = "Interface"
		log.Method = string(c.Request.Method())
		log.API = string(c.Request.Path())
		log.UserAgent = string(c.Request.Header.UserAgent())
		log.IP = c.ClientIP()

		reqBodyStr := string(c.Request.Body())
		if len(reqBodyStr) > 200 {
			reqBodyStr = reqBodyStr[:200]
		}
		log.ReqContent = reqBodyStr

		respBodyStr := string(c.Request.Body())
		if len(respBodyStr) > 200 {
			respBodyStr = respBodyStr[:200]
		}

		if c.Response.Header.StatusCode() == 200 {
			log.Success = true
		}

		costTime := time.Since(start).Milliseconds()
		log.Time = costTime

		v, exist := c.Get("user_id")
		if exist || v == nil {
			v = "0"
		}
		var userIDStr string
		var username = "Anonymous"
		var ok bool

		userIDStr, ok = v.(string)

		if !ok {
			userIDStr = "0"
		}

		userID, _ := strconv.Atoi(userIDStr)

		userInfo, _ := admin.NewUser(ctx, c).Info(int64(userID))

		if userInfo != nil {
			username = userInfo.Nickname
		}

		log.Operators = username

		err := admin.NewLogs(ctx, c).Create(&logs)

		if err != nil {
			hlog.Error(err)
		}

	}
}
