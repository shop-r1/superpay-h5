# superpay-h5

merchant_id	String(10)
必填，10088，商户 ID，由 Supay Tech 提供

authentication_code可选	String(32)
c3c5134dea12900d29c5deb2d5775162，商户验证码，由 Supay Tech 提供

product_title	String(255)
必填， Billy Lynn's Long Halftime Walk，商户商品名

merchant_trade_no	String(64)
必填，88888888，商户自身系统交易单号

currency	String(8)
必填，AUD，币种（现阶段仅支持 AUD）

total_amount	Number
必填，Number(8,2)，15，交易金额

create_time	DateTime
必填，2012-11-11 18:08:50，交易创建时间（GMT 时区）

notification_url	String(255)
必填， http%3A%2F%2Fwww.samplemerchant.com%2Fnotification，支付结果异步回调通知链接

token	String(32)
必填，基于参数字符串的 MD5 加密令牌

product_desc可选	String(255)
商户商品明细描述

return_url可选	String(255)
当客户支付成功后同步返回页面

merchant_validation_可选	String(128)
code] 商户自定义验证码。供同步返回时使用。

rmb_amount可选	Number
Number(8,2)，如果商户需要交易仅显示为人民币金额时， 请在此字段填入人民币金额

return_target可选	String(2)
如果商户的应用是在微信内部，则请指定此 参数的值为 WX, 否则忽略此参数

redirect_url可选	String(255)
用户使用手机在微信商城支付取消的时候 手机微信内部跳转回这个 url