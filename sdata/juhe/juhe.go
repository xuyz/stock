package juhe

import (
	"strconv"

	"github.com/xuyz/stock/sdata"
)

type Juhe struct {
	key string
}

func NewJuhe(key string) *Juhe {
	return &Juhe{key: key}
}

func StockSimpleConvert(simple stockSimple) *sdata.StockSimple {
	open, err := strconv.ParseFloat(simple.Open, 64)
	if err != nil {
		return nil
	}

	high, err := strconv.ParseFloat(simple.High, 64)
	if err != nil {
		return nil
	}
	low, err := strconv.ParseFloat(simple.Low, 64)
	if err != nil {
		return nil
	}
	trade, err := strconv.ParseFloat(simple.Trade, 64)
	if err != nil {
		return nil
	}

	settlement, err := strconv.ParseFloat(simple.Settlement, 64)
	if err != nil {
		return nil
	}
	return &sdata.StockSimple{
		Code:          simple.Code,
		Name:          simple.Name,
		Open:          int64(open * 1000),
		High:          int64(high * 1000),       // 最高 单位厘
		Low:           int64(low * 1000),        // 最低 单位厘
		Trade:         int64(trade * 1000),      // 最新价 单位厘
		PriceChange:   simple.Pricechange,       // 涨跌额
		ChangePercent: simple.Changepercent,     // 涨跌幅
		Settlement:    int64(settlement * 1000), // 昨收
		Volume:        simple.Volume,            // 成交量
		Amount:        simple.Amount,            // 成效额
	}
}

func (j *Juhe) Stocks() []*sdata.StockSimple {
	var res []*sdata.StockSimple

	dsh := sh(j.key)
	for _, v := range dsh {
		vv := StockSimpleConvert(v)
		if vv != nil {
			res = append(res, vv)
		}
	}

	dsz := sz(j.key)
	for _, v := range dsz {
		vv := StockSimpleConvert(v)
		if vv != nil {
			res = append(res, vv)
		}
	}

	return res
}

func (j *Juhe) StockInfo(code string) map[string]string {
	return nil
}
