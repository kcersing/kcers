package dal

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	aliyun_sms "kcers/biz/dal/aliyun-sms"
	"kcers/biz/dal/cache"
	"kcers/biz/dal/casbin"
	"kcers/biz/dal/config"
	db "kcers/biz/dal/db/mysql"
	"kcers/biz/dal/logger"
	"kcers/biz/dal/minio"
)

func Init() {
	config.InitConfig()
	hlog.Info("Init config ok!")
	logger.InitLogger()
	hlog.Info("Init logger ok!")
	cache.InitCache()
	hlog.Info("Init cache ok!")
	db.InitDB()
	hlog.Info("Init db ok!")
	casbin.InitCasbin()
	hlog.Info("Init casbin ok!")
	minio.InitMinio()
	hlog.Info("Init minio ok!")
	aliyun_sms.InitAliyunSms()
	hlog.Info("Init aliyun sms ok!")
	//go func() {
	//	wechat.InitWXPaymentApp()
	//	wechat.InitMiniProgramApp()
	//}()
	hlog.Info("Init ok!")

}
