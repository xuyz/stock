package juhe

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xuyz/stock/sutils"
)

type stockSimple struct {
	Amount        int64  `json:"amount"`
	Buy           string `json:"buy"`
	Changepercent string `json:"changepercent"`
	Code          string `json:"code"`
	High          string `json:"high"`
	Low           string `json:"low"`
	Name          string `json:"name"`
	Open          string `json:"open"`
	Pricechange   string `json:"pricechange"`
	Sell          string `json:"sell"`
	Settlement    string `json:"settlement"`
	Symbol        string `json:"symbol"`
	Ticktime      string `json:"ticktime"`
	Trade         string `json:"trade"`
	Volume        int64  `json:"volume"`
}

type juheListResp struct {
	ErrorCode int    `json:"error_code"`
	Reason    string `json:"reason"`
	Result    struct {
		Data       []stockSimple `json:"data"`
		Num        string        `json:"num"`
		Page       string        `json:"page"`
		TotalCount string        `json:"totalCount"`
	} `json:"result"`
}

// 上海
func sh(key string) []stockSimple {
	return listAll("http://web.juhe.cn:8080/finance/stock/shall?key=" + key)
}

// 深圳
func sz(key string) []stockSimple {
	return listAll("http://web.juhe.cn:8080/finance/stock/szall?key=" + key)
}

// 美国
func usa(key string) []stockSimple {
	return listAll("http://web.juhe.cn:8080/finance/stock/usaall?key=" + key)
}

// 香港
func hk(key string) []stockSimple {
	return listAll("http://web.juhe.cn:8080/finance/stock/hkall?key=" + key)
}

func listAll(u string) []stockSimple {
	var res []stockSimple
	i := 1
	for {
		d, err := list(u + fmt.Sprintf("&page=%d", i))
		if err != nil {
			break
		}

		if len(d) == 0 {
			break
		}

		res = append(res, d...)
		i++
	}

	return res
}

func list(u string) ([]stockSimple, error) {
	resp, err := sutils.Get(u)
	if err != nil {
		return nil, err
	}

	res := juheListResp{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	if res.ErrorCode != 0 {
		return nil, errors.New(res.Reason)
	}

	return res.Result.Data, nil
}
