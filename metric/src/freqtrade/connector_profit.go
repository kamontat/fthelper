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
	if profit, err := FetchProfit(conn); err == nil {
		return profit
	}
	return EmptyProfit()
}

func FetchProfit(conn *Connection) (*Profit, error) {
	var name = API_PROFIT
	if data, err := conn.Cache(name, conn.ExpireAt(name), func() (interface{}, error) {
		return GetProfit(conn)
	}); err == nil {
		return data.(*Profit), nil
	} else {
		return nil, err
	}
}

func GetProfit(conn *Connection) (*Profit, error) {
	var target = new(Profit)
	var err = GetConnector(conn, API_PROFIT, &target)
	return target, err
}
