package superpay_h5

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Pay struct {
	MerchantId         string
	AuthenticationCode string
}

//创建订单
func NewPay(merchantId, authenticationCode string) *Pay {
	return &Pay{
		MerchantId:         merchantId,
		AuthenticationCode: authenticationCode,
	}
}

func (e Pay) CreateOrderWx(req *CreateOrderReqWx) (*CreateOrderRspWx, error) {
	u := newWxReqUrl()
	req.MerchantId = e.MerchantId
	req.AuthenticationCode = e.AuthenticationCode
	req.setToken()
	req.ReturnTarget = "WX"
	u.RawQuery = getEncodeQuery(*req, true, u.Query()).Encode()
	fmt.Println(u.RawQuery)
	fmt.Println(u.String())
	rsp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	var rb []byte
	rb, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	response := &CreateOrderRspWx{}
	err = json.Unmarshal(rb, response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (e Pay) CreateOrderAlipay(req *CreateOrderReqAlipay) (string, error) {
	u := newAlipay()
	req.MerchantId = e.MerchantId
	req.AuthenticationCode = e.AuthenticationCode
	req.setToken()
	u.RawQuery = getEncodeQuery(*req, true, u.Query()).Encode()
	return u.String(), nil
}
