package wechat

import (
	"github.com/ArtisanCloud/PowerLibs/v3/logger/drivers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"kcers/biz/dal/config"
	"kcers/biz/pkg/consts"
	"sync"
)

var MiniProgramApp *miniProgram.MiniProgram

var onceMiniProgramApp sync.Once

func InitMiniProgramApp() {
	onceMiniProgramApp.Do(func() {
		MiniProgramApp = NewMiniMiniProgramService()
	})
}

const TIMEZONE = "asia/shanghai"
const DATETIME_FORMAT = "20060102"

func NewMiniMiniProgramService() *miniProgram.MiniProgram {
	conf := config.GlobalServerConfig.Wechat
	var cache kernel.CacheInterface
	if config.GlobalServerConfig.Redis.Host != "" {
		cache = kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs: []string{config.GlobalServerConfig.Redis.Host},
		})
	}
	wechatFilePath := consts.WechatFilePath

	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:        conf.Appid, // 小程序、公众号或者企业微信的appid
		Secret:       conf.MchId, // 商户号 appID
		ResponseType: response.TYPE_MAP,
		Http:         miniProgram.Http{},
		Log: miniProgram.Log{
			Driver: &drivers.SimpleLogger{},
			Level:  "debug",
			File:   wechatFilePath + "/mini_log.log",
		},
		Cache:     cache,
		HttpDebug: true,
		Debug:     false,
	})
	// 2. 调用小程序的授权登陆接口
	//code := "CODE" // 前端小程序登录时，从微信获取的code
	//rs, err := app.Auth.Session(context.Background(), code)
	//hlog.Info(rs.OpenID)
	if err != nil {
		hlog.Error(err)
	}

	return app
}
