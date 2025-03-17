package douyin

var AppId = ""
var AppSecret = ""
var AccountId = ""

const (
	// AuthUrl 赋权
	AuthUrl string = "https://auth.dylk.com/auth-isv/"
	// ClientToken 获取token
	ClientToken = "https://open.douyin.com/oauth/client_token/"
	// ShopPoiQuery 查询店铺信息
	ShopPoiQuery = "https://open.douyin.com/goodlife/v1/shop/poi/query/"
	// OrderQuery 查询订单信息
	OrderQuery = "https://open.douyin.com/goodlife/v1/trade/order/query/"
	// CouponsPrepare 查询订单信息
	CouponsPrepare = "https://open.douyin.com/goodlife/v1/fulfilment/certificate/prepare/"
	// CouponsVerify 验券
	CouponsVerify = "https://open.douyin.com/goodlife/v1/fulfilment/certificate/verify/"

	CrmQuery = "https://open.douyin.com/goodlife/v1/open_api/crm/clue/query/"

	Goods = "https://open.douyin.com/goodlife/v1/goods/product/online/get/"

	GoodsList = "https://open.douyin.com/goodlife/v1/goods/product/online/get/"
)
