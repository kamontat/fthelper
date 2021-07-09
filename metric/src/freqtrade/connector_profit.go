package freqtrade

type Profit struct {
	UnrealizedCryptoProfit float64 `json:"profit_closed_coin"`
	UnrealizedFiatProfit   float64 `json:"profit_closed_fiat"`
	RealizedCryptoProfit   float64 `json:"profit_all_coin"`
	RealizedFiatProfit     float64 `json:"profit_all_fiat"`

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

func EmptyProfit() *Profit {
	return &Profit{
		UnrealizedCryptoProfit: 0,
		UnrealizedFiatProfit:   0,
		RealizedCryptoProfit:   0,
		RealizedFiatProfit:     0,

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

func NewProfit(conn *Connection) *Profit {
	var name, expireAt, query = Connector(conn, API_PROFIT)
	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = new(Profit)
		err := conn.GET(name, query, &target)
		return target, err
	}); err == nil && data != nil {
		return data.(*Profit)
	}

	return EmptyProfit()
}
