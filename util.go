package superpay_h5

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
)

func newWxReqUrl() *url.URL {
	u := &url.URL{}
	u.Host = "api.superpayglobal.com"
	u.Scheme = "https"
	u.Path = "/payment/wxpayproxy/merchant_request"
	return u
}

func newAlipay() *url.URL {
	u := &url.URL{}
	u.Host = "api.superpayglobal.com"
	u.Scheme = "https"
	u.Path = "/payment/bridge/merchant_request"
	return u
}

func mapToUrl(params map[string]string, key string) string {
	u := url.Values{}
	for k, v := range params {
		if v == "" {
			continue
		}
		u.Set(k, v)
	}
	s := u.Encode()
	if key != "" {
		s += "&key=" + key
	}
	return s
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
