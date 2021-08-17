package freqtrade

import (
	"fmt"
	"time"

	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
)

const PROFIT_CONST = "profit"

func NewProfit() *profit {
	return &profit{
		UnrealizedCryptoProfit:  0,
		UnrealizedFiatProfit:    0,
		UnrealizedPercentProfit: 0,
		RealizedCryptoProfit:    0,
		RealizedFiatProfit:      0,
		RealizedPercentProfit:   0,

		TotalTrades:  0,
		ClosedTrades: 0,
		WinTrades:    0,
		LossTrades:   0,

		FirstTradeTimestamp: 0,
		LastTradeTimestamp:  0,
		AverageDuration:     "00:00:00",

		BestPair: "",
		BestRate: 0,
	}
}

type profit struct {
	RealizedCryptoProfit float64 `json:"profit_closed_coin"`
	RealizedFiatProfit   float64 `json:"profit_closed_fiat"`
	// Percent is number from 0 to 1 represent percentage of profit from start balance
	RealizedPercentProfit float64 `json:"profit_closed_ratio"`

	UnrealizedCryptoProfit float64 `json:"profit_all_coin"`
	UnrealizedFiatProfit   float64 `json:"profit_all_fiat"`
	// Percent is number from 0 to 1 represent percentage of profit from start balance
	UnrealizedPercentProfit float64 `json:"profit_all_ratio"`

	TotalTrades  int `json:"trade_count"`
	ClosedTrades int `json:"closed_trade_count"`
	WinTrades    int `json:"winning_trades"`
	LossTrades   int `json:"losing_trades"`

	// first trade timestamp (millisecond)
	FirstTradeTimestamp int64 `json:"first_trade_timestamp"`

	// latest trade timestamp (millisecond)
	LastTradeTimestamp int64 `json:"latest_trade_timestamp"`

	// format 00:00:00
	AverageDuration string `json:"avg_duration"`

	BestPair string  `json:"best_pair"`
	BestRate float64 `json:"best_rate"`
}

func (p *profit) GetAverageDuration() time.Duration {
	var h, m, s int
	n, err := fmt.Sscanf(p.AverageDuration, "%d:%d:%d", &h, &m, &s)
	if err != nil || n != 3 {
		return -1
	}

	var second = (h * 3600) + (m * 60) + s
	return time.Duration(second) * time.Second
}

func (p *profit) Name() string {
	return PROFIT_CONST
}

func (p *profit) Build(connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	err := connection.Http.GET(p.Name(), p)
	return p, err
}

func ToProfit(connector connection.Connector) (*profit, error) {
	raw, err := connector.Connect(PROFIT_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*profit), nil
}
