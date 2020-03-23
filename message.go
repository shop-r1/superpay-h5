package superpay_h5

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

type CreateOrderReq struct {
	MerchantId         string  `url:"merchant_id"`
	AuthenticationCode string  `url:"-"`
	ProductTitle       string  `url:"product_title"`
	MerchantTradeNo    string  `url:"merchant_trade_no"`
	Currency           string  `url:"currency"`
	TotalAmount        float64 `url:"total_amount"`
	CreateTime         string  `url:"create_time"`
	NotificationUrl    string  `url:"notification_url"`
	Token              string  `url:"token"`
	ReturnUrl          string  `url:"return_url"`
	RmbAmount          float64 `url:"rmb_amount"`
}

type CreateOrderReqWx struct {
	CreateOrderReq `url:"struct"`
	ReturnTarget   string `url:"return_target"`
	RedirectUrl    string `url:"redirect_url"`
}

type CreateOrderRspWx struct {
	Result          string `json:"result"`
	Reason          string `json:"reason"`
	QRCodeImgPath   string `json:"QRCodeImgPath"`
	QRCodeURL       string `json:"QRCodeURL"`
	SupayCashierURL string `json:"supayCashierURL"`
}

type CreateOrderReqAlipay struct {
	CreateOrderReq `url:"struct"`
	MobileFlag     string `url:"mobile_flag"`
}

func getEncodeQuery(e interface{}, skipEmpty bool, u url.Values) url.Values {
	obj1 := reflect.TypeOf(e)
	obj2 := reflect.ValueOf(e)
	for i := 0; i < obj1.NumField(); i++ {
		switch obj1.Field(i).Tag.Get("url") {
		case "total_amount", "rmb_amount":
			if obj1.Field(i).Tag.Get("url") == "rmb_amount" && obj2.Field(i).Float() == 0 {
				continue
			}
			u.Set(obj1.Field(i).Tag.Get("url"), strconv.FormatFloat(obj2.Field(i).Float(), 'f', 2, 64))
			continue
		}
		if (skipEmpty && obj2.Field(i).IsZero()) || obj1.Field(i).Tag.Get("url") == "-" {
			continue
		}
		if obj1.Field(i).Tag.Get("url") == "struct" {
			u = getEncodeQuery(obj2.Field(i).Interface(), skipEmpty, u)
			continue
		}
		u.Set(obj1.Field(i).Tag.Get("url"), obj2.Field(i).String())
	}
	return u
}

func (e *CreateOrderReq) setToken() {
	s := fmt.Sprintf("merchant_id=%s&authentication_code=%s&merchant_trade_no=%s&total_amount=%s",
		e.MerchantId, e.AuthenticationCode, e.MerchantTradeNo,
		strconv.FormatFloat(e.TotalAmount, 'f', 2, 64))
	fmt.Println(s)
	e.Token = md5V(s)
}
