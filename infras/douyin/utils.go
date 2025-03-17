package douyin

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"net/url"
	"time"
)

// GetAccessToken 获取Token
func GetAccessToken() (string, error) {
	params := map[string]string{
		"client_key":    AppId,
		"client_secret": AppSecret,
		"grant_type":    "client_credential",
	}
	resp, err := GetUrl(consts.MethodPost, map[string]string{}, params, ClientToken)
	if err != nil {
		return "", err
	}
	if respMap, ok := resp.(map[string]interface{}); ok {
		if dataMap, ok := respMap["data"].(map[string]interface{}); ok {
			return dataMap["access_token"].(string), nil
		}
	}
	return "", fmt.Errorf("resp: %v", "resp is not map[string]interface{}")

}
func GetUrl(method string, headers map[string]string, params map[string]string, url string) (interface{}, error) {
	c, err := client.NewClient()
	if err != nil {
		return nil, err
	}
	req, res := &protocol.Request{}, &protocol.Response{}
	req.SetOptions(config.WithRequestTimeout(5 * time.Second))

	req.SetMethod(method)
	if method == consts.MethodPost {
		req.SetFormData(params)
	}

	if method == consts.MethodGet {
		url = URLEncode(params, url)
	}

	req.SetHeaders(headers)

	req.SetRequestURI(url)

	err = c.Do(context.Background(), req, res)

	var resp interface{}
	err = sonic.Unmarshal(res.Body(), &resp)
	if err != nil {
		hlog.Infof("resp: %v", err)
		return nil, err
	}

	return resp, err
}

func URLEncode(req map[string]string, URL string) string {
	encode := url.Values{}
	for k, v := range req {
		encode.Add(k, fmt.Sprintf("%v", v))
	}
	return fmt.Sprintf("%s?%s", URL, encode.Encode())
}

func GetHeadersToken() map[string]string {
	accessToken, _ := GetAccessToken()
	return map[string]string{
		"content-type": "application/json",
		"access-token": accessToken,
	}
}
