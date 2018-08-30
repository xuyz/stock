package scron

import (
	"fmt"
	"strings"
	"time"

	"github.com/xuyz/stock/scache"
	"github.com/xuyz/stock/sconfig"
	"github.com/xuyz/stock/sdata"
	"github.com/xuyz/stock/sdata/juhe"
	"github.com/xuyz/stock/srule"

	"github.com/xuyz/stock/report"

	"github.com/robfig/cron"
)

func Start() {
	c := cron.New()

	//c.AddFunc("@every 10m", hot) // 每10分钟
	//c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") }) // 没半小时
	c.AddFunc("1 1 10-15 * * *", hot) // 每天(10-15):01:01

	c.Start()
}

func hot() {
	var stocks []*sdata.StockSimple
	cacheKey := "stocks"
	cached, ok := scache.Get(cacheKey)
	if !ok || cached == nil {
		stocks = juhe.NewJuhe(sconfig.JuheKey()).Stocks()
		if len(stocks) != 0 {
			scache.Set(cacheKey, stocks, time.Hour*12)
		}
	} else {
		stocks = cached.([]*sdata.StockSimple)
	}

	var res []string
	res = append(res, "名称  编号  成交量  成交金额")
	for _, v := range stocks {
		if srule.Hot(v) {
			res = append(res, fmt.Sprintf("%s %s %d %d", v.Name, v.Code, v.Volume, v.Amount))
		}
	}
	report.StockWechat(strings.Join(res, "\n"))
}
