package report

import "github.com/xuyz/stock/sconfig"

func StockWechat(content string) error {
	return WechatReport(sconfig.WechatAgentID(), sconfig.WechatCorID(), sconfig.WechatSecret(), "", content)
}
