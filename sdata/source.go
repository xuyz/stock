package sdata

// StockSimple 股票基本信息
type StockSimple struct {
	Code          string // 股票编码
	Name          string // 股票名
	Open          int64  // 开盘价 单位厘
	High          int64  // 最高 单位厘
	Low           int64  // 最低 单位厘
	Trade         int64  // 最新价 单位厘
	PriceChange   string // 涨跌额
	ChangePercent string // 涨跌幅
	Settlement    int64  // 昨收
	Volume        int64  // 成交量
	Amount        int64  // 成效额
}

// Source 数据源接口
type Source interface {
	// Stocks 股票列表
	Stocks() []*StockSimple

	// StockInfo 单个股票信息
	StockInfo(code string) map[string]string
}
