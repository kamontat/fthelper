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

// Create new balance
func NewBalance(conn *Connection) *Balance {
	if balance, err := FetchBalance(conn); err == nil {
		return balance
	}
	return EmptyBalance()
}

// Fetch balance to local cache
func FetchBalance(conn *Connection) (*Balance, error) {
	var name = API_BALANCE
	data, err := conn.Cache(name, conn.ExpireAt(name), func() (interface{}, error) {
		return GetBalance(conn)
	})
	return data.(*Balance), err
}

// Get balance without cache
func GetBalance(conn *Connection) (*Balance, error) {
	var name = API_BALANCE
	var target = new(Balance)
	var err = conn.GET(name, conn.QueryValues(name), &target)
	return target, err
}
