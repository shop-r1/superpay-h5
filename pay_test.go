package superpay_h5

import (
	"testing"
	"time"

	"github.com/sanity-io/litter"
)

var (
	MerchantId         = "71874"
	AuthenticationCode = "cPZhP0rPHDrQxYgdepSgPbjUqjBeN5AG"
)

func TestPay_CreateOrderAlipay(t *testing.T) {
	type fields struct {
		MerchantId         string
		AuthenticationCode string
	}
	type args struct {
		req *CreateOrderReqAlipay
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test01",
			fields{
				MerchantId:         MerchantId,
				AuthenticationCode: AuthenticationCode,
			},
			args{
				&CreateOrderReqAlipay{
					CreateOrderReq: CreateOrderReq{
						ProductTitle:    "abca",
						MerchantTradeNo: "abcb",
						Currency:        "AUD",
						TotalAmount:     0.01,
						CreateTime:      time.Now().Add(-8 * time.Hour).Format("2006-01-02 15:04:05"),
						NotificationUrl: "https://www.baidu.com",
						ReturnUrl:       "https://jiajia.r1shop.net/m",
						//RmbAmount:          0.01,
					},
					MobileFlag: "T",
				},
			},
			false,
		},
		{
			"test01",
			fields{
				MerchantId:         MerchantId,
				AuthenticationCode: AuthenticationCode,
			},
			args{
				&CreateOrderReqAlipay{
					CreateOrderReq: CreateOrderReq{
						ProductTitle:    "abca",
						MerchantTradeNo: "abcb",
						Currency:        "AUD",
						CreateTime:      time.Now().Add(-8 * time.Hour).Format("2006-01-02 15:04:05"),
						NotificationUrl: "https://www.baidu.com",
						ReturnUrl:       "https://jiajia.r1shop.net/m",
						RmbAmount:       0.05,
					},
					MobileFlag: "T",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Pay{
				MerchantId:         tt.fields.MerchantId,
				AuthenticationCode: tt.fields.AuthenticationCode,
			}
			got, err := e.CreateOrderAlipay(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrderAlipay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			litter.Dump(got)
		})
	}
}

func TestPay_CreateOrderWx(t *testing.T) {
	type fields struct {
		MerchantId         string
		AuthenticationCode string
	}
	type args struct {
		req *CreateOrderReqWx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test01",
			fields{
				MerchantId:         MerchantId,
				AuthenticationCode: AuthenticationCode,
			},
			args{req: &CreateOrderReqWx{
				CreateOrderReq: CreateOrderReq{
					ProductTitle:    "abca",
					MerchantTradeNo: "540598752433143809",
					Currency:        "AUD",
					TotalAmount:     0.01,
					CreateTime:      time.Now().Add(-8 * time.Hour).Format("2006-01-02 15:04:05"),
					NotificationUrl: "https://www.baidu.com",
					ReturnUrl:       "https://jiajia.r1shop.net/m",
					//RmbAmount:          0.01,
				},
				RedirectUrl: "https://jiajia.r1shop.net/m",
			}},
			false,
		},
		{
			"test02",
			fields{
				MerchantId:         MerchantId,
				AuthenticationCode: AuthenticationCode,
			},
			args{req: &CreateOrderReqWx{
				CreateOrderReq: CreateOrderReq{
					ProductTitle:    "abca",
					MerchantTradeNo: "540598752433143809",
					Currency:        "AUD",
					//TotalAmount:        0.05,
					CreateTime:      time.Now().Add(-8 * time.Hour).Format("2006-01-02 15:04:05"),
					NotificationUrl: "https://www.baidu.com",
					ReturnUrl:       "https://jiajia.r1shop.net/m",
					RmbAmount:       0.05,
				},
				RedirectUrl: "https://jiajia.r1shop.net/m",
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Pay{
				MerchantId:         tt.fields.MerchantId,
				AuthenticationCode: tt.fields.AuthenticationCode,
			}
			got, err := e.CreateOrderWx(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateOrderWx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			litter.Dump(got)
		})
	}
}
