package freqtrade

type CryptoBalance struct {
	Symbol          string  `json:"currency"`
	Balance         float64 `json:"balance"`
	Free            float64 `json:"free"`
	Used            float64 `json:"used"`
	EstStakeBalance float64 `json:"est_stake"`
	StakeSymbol     string  `json:"stake"`
}

type Balance struct {
	Currencies   []*CryptoBalance
	CryptoValue  float64 `json:"total"`
	CryptoSymbol string  `json:"stake"`
	FiatValue    float64 `json:"value"`
	FiatSymbol   string  `json:"symbol"`
}

func EmptyBalance() *Balance {
	return &Balance{
		CryptoValue:  0,
		CryptoSymbol: "UKN",
		FiatValue:    0,
		FiatSymbol:   "UKN",
		Currencies:   make([]*CryptoBalance, 0),
	}
}

func NewBalance(conn *Connection) *Balance {
	var name, expireAt, query = Connector(conn, API_BALANCE)
	if data, err := conn.Cache(name, expireAt, func() (interface{}, error) {
		var target = new(Balance)
		err := conn.GET(name, query, &target)
		return target, err
	}); err == nil && data != nil {
		return data.(*Balance)
	}

	return EmptyBalance()
}
