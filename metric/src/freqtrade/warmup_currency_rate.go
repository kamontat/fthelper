package freqtrade

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/loggers"
)

const (
	CACHE_CURRENCY_RATE = "currency-rate"
)

// update currency rate every 24 hours
func warmupCurrencyRate(conn *Connection, logger *loggers.Logger) error {
	return caches.Global.UpdateFn(CACHE_CURRENCY_RATE, func(o interface{}) (interface{}, error) {
		logger.Info("initial currency rate value")
		var balance = NewBalance(conn)
		if balance.FiatValue <= 0 || balance.CryptoValue <= 0 {
			return nil, fmt.Errorf("cannot get correct balance version, try again next time")
		}

		var rate = balance.FiatValue / balance.CryptoValue
		return rate, nil
	}, "24h")
}
