package douyin

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	urll "net/url"
	"strconv"
	"strings"
	"time"
)

func AuthWithBind() {

	fast, err := GenAuthWithBindValidUrlFast(
		"4", AppId, "", []string{"1", "2", "5", "6", "9", "10", "12", "15", "16"})
	if err != nil {
		return
	}
	hlog.Info(fast)
}

func GenAuthWithBindValidUrlFast(solutionKey string, outShopId string, extra string, permissionKeys []string) (result string, err error) {
	dto := &GenAuthWithBindValidUrlDto{
		Timestamp:      strconv.FormatInt(time.Now().Unix(), 10),
		Extra:          extra,
		SolutionKey:    solutionKey,
		OutShopId:      outShopId,
		PermissionKeys: permissionKeys,
	}
	return GenAuthValidUrl(dto)
}

// GenAuthValidUrl 生成一个合法的签名url
func GenAuthValidUrl(dto *GenAuthWithBindValidUrlDto) (result string, err error) {
	parsedUrl, err := urll.Parse(AuthUrl)
	if err != nil {
		return "", err
	}

	query := map[string]string{}
	query["client_key"] = AppId
	query["timestamp"] = dto.Timestamp
	query["charset"] = "UTF-8"
	query["solution_key"] = dto.SolutionKey
	query["permission_keys"] = strings.Join(dto.PermissionKeys, ",")
	query["out_shop_id"] = dto.OutShopId
	if dto.Extra != "" {
		query["extra"] = dto.Extra
	}

	signResult := SignV2(AppSecret, "", query)

	// set final url params
	parsedQuery := parsedUrl.Query()
	parsedQuery.Add("sign", signResult)
	for k, v := range query {
		parsedQuery.Add(k, v)
	}

	parsedUrl.RawQuery = parsedQuery.Encode()
	return parsedUrl.String(), nil
}

type GenAuthWithBindValidUrlDto struct {
	// url开始生效时间
	Timestamp string
	// 服务商自定义字符串，长度不可超过1000字节
	Extra string
	// 解决方案 具体值参照
	SolutionKey string
	// 授权能力列表
	PermissionKeys []string
	// 外部门店ID
	OutShopId string
}
