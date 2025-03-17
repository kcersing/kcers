package douyin

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func GetCrmQuery(params map[string]string) (interface{}, error) {
	headers := GetHeadersToken()
	resp, err := GetUrl(consts.MethodGet, headers, params, OrderQuery)
	if err != nil {
		return nil, err
	}
	hlog.Infof("resp: %v", resp)
	return nil, err
}

// GerOrder 获取订单信息
//
//	params map[string]string{
//		"account_id": "",
//		"order_id":   "",
//	}
func GerOrder(params map[string]string) (interface{}, error) {
	headers := GetHeadersToken()
	resp, err := GetUrl(consts.MethodGet, headers, params, OrderQuery)
	if err != nil {
		return nil, err
	}
	hlog.Infof("resp: %v", resp)
	return nil, err
}

// GetShopPoi 获取店铺信息
func GetShopPoi() error {
	headers := GetHeadersToken()
	params := map[string]string{
		"account_id": AccountId,
		"page":       "1",
		"size":       "100",
	}
	resp, err := GetUrl(consts.MethodGet, headers, params, ShopPoiQuery)
	if err != nil {
		return err
	}
	hlog.Infof("resp: %v", resp)
	return err
}

// CouponsVerifyPrepare 验券准备
//
//	params map[string]string{
//		"poi_id": "",
//		"code":   "",
//	}
func CouponsVerifyPrepare(params map[string]string) (interface{}, error) {
	headers := GetHeadersToken()
	resp, err := GetUrl(consts.MethodGet, headers, params, CouponsPrepare)
	if err != nil {
		return nil, err
	}
	hlog.Infof("resp: %v", resp)
	return nil, err
}

// CouponsVerifyOn 验券
//
//	params map[string]string{
//		"verify_token": "",
//		"poi_id":         "",
//	}
func CouponsVerifyOn(params map[string]string) (interface{}, error) {
	headers := GetHeadersToken()

	resp, err := GetUrl(consts.MethodGet, headers, params, CouponsVerify)
	if err != nil {
		return nil, err
	}
	hlog.Infof("resp: %v", resp)
	return nil, err
}

func GetGoods(params map[string]string) (interface{}, error) {
	headers := GetHeadersToken()
	resp, err := GetUrl(consts.MethodGet, headers, params, Goods)
	if err != nil {
		return nil, err
	}
	hlog.Infof("resp: %v", resp)
	return nil, err
}

// GetGoodsList 获取商品列表
// params map[string]string{
// "account_id" : “”,
// "count" : 50,
// "goods_creator_type" : 1,
// "cursor"=>0
// }
func GetGoodsList(params map[string]string) (interface{}, error) {
	headers := GetHeadersToken()
	resp, err := GetUrl(consts.MethodGet, headers, params, GoodsList)
	if err != nil {
		return nil, err
	}
	hlog.Infof("resp: %v", resp)
	return nil, err
}
