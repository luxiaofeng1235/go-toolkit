package api

import (
	"go-novel/utils"

	"github.com/gin-gonic/gin"
)

type Pay struct{}

func (pay *Pay) AliPayNotify(c *gin.Context) {
	return
}

// 支付宝回调
func (pay *Pay) AliPayCallback(c *gin.Context) {
	utils.Success(c, "", "alipay callback")
	return
}

// 微信回调
func (pay *Pay) WxPayNotify(c *gin.Context) {
	utils.Success(c, "", "wxpay callback")
	return
}
