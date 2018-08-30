package srule

import "github.com/xuyz/stock/sdata"

// Hot 比较热的股票
// 成交量大于100万 或 成交额大于1亿
func Hot(ss *sdata.StockSimple) bool {
	if ss.Volume > 1000000 {
		return true
	}

	if ss.Amount > 100000000 {
		return true
	}

	return false
}
