package freqtrade

import (
	"errors"

	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/shared/datatype"
)

const BALANCE_CONST = "balance"

func NewBalance() *balance {
	return &balance{
		Currencies:   make([]*crypto, 0),
		CryptoValue:  0,
		CryptoSymbol: "UKN",
		FiatValue:    0,
		FiatSymbol:   "UKN",
	}
}

type crypto struct {
	Symbol          string  `json:"currency"`
	Balance         float64 `json:"balance"`
	Free            float64 `json:"free"`
	Used            float64 `json:"used"`
	EstStakeBalance float64 `json:"est_stake"`
	StakeSymbol     string  `json:"stake"`
}

type balance struct {
	Currencies   []*crypto
	CryptoValue  float64 `json:"total"`
	CryptoSymbol string  `json:"stake"`
	FiatValue    float64 `json:"value"`
	FiatSymbol   string  `json:"symbol"`
}

func (b *balance) Name() string {
	return BALANCE_CONST
}

func (b *balance) Build(connector connection.Connector, connection *connection.Connection, history *datatype.Queue) (interface{}, error) {
	err := connection.Http.GET(b.Name(), b)
	if err == nil && b.CryptoValue > 0 && b.FiatValue <= 0 {
		err = errors.New("invalid balance because fiat value is zero")
	}

	return b, err
}

func (b *balance) Empty() bool {
	return b.FiatValue == 0 && b.CryptoValue == 0
}

func ToBalance(connector connection.Connector) (*balance, error) {
	raw, err := connector.Connect(BALANCE_CONST)
	if err != nil {
		return nil, err
	}
	return raw.(*balance), nil
}
