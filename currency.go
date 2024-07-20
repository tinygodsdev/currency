package currency

import "strings"

const (
	RUB = "RUB"
	EUR = "EUR"
	USD = "USD"
	GBP = "GBP"
	ITL = "ITL"
	JPY = "JPY"
	KRW = "KRW"
	INR = "INR"
	UAH = "UAH"
	ILS = "ILS"
	AUD = "AUD"
	CAD = "CAD"
	CHF = "CHF"
	CNY = "CNY"
	HKD = "HKD"
	NZD = "NZD"
	SEK = "SEK"
	SGD = "SGD"
	MXN = "MXN"
	NOK = "NOK"
	BRL = "BRL"
	ZAR = "ZAR"
)

func getCurrencyMapping() map[string]string {
	return map[string]string{
		"₽":    RUB,
		"руб.": RUB,
		"руб":  RUB,
		"rub":  RUB,
		"rub.": RUB,
		"р.":   RUB,
		"р":    RUB,

		"€":   EUR,
		"eur": EUR,

		"$":   USD,
		"usd": USD,

		"£":   GBP,
		"gbp": GBP,

		"₤":   ITL,
		"itl": ITL,

		"¥":   JPY,
		"jpy": JPY,

		"₩":   KRW,
		"krw": KRW,

		"₹":   INR,
		"inr": INR,

		"₴":   UAH,
		"uah": UAH,

		"₪":   ILS,
		"ils": ILS,

		"a$":  AUD,
		"aud": AUD,

		"c$":  CAD,
		"cad": CAD,

		"chf": CHF,

		"cn¥": CNY,
		"cny": CNY,

		"hk$": HKD,
		"hkd": HKD,

		"nz$": NZD,
		"nzd": NZD,

		"sek": SEK,

		"sgd": SGD,

		"mx$": MXN,
		"mxn": MXN,

		"nok": NOK,

		"r$":  BRL,
		"brl": BRL,

		"zar": ZAR,
	}
}

// NormalizeCurrency normalizes a currency string to a standard currency code
func NormalizeCurrency(currency string, customMapping map[string]string) (string, bool) {
	// if custom mapping is provided, use it
	if customMapping != nil {
		if v, ok := customMapping[currency]; ok {
			return v, true
		}
	}

	// otherwise use the default mapping
	mapping := getCurrencyMapping()
	if v, ok := mapping[strings.ToLower(currency)]; ok {
		return v, true
	}

	return currency, false
}

// SupportedCurrencies returns a list of supported currencies
func SupportedCurrencies() []string {
	mapping := getCurrencyMapping()
	currencies := make([]string, 0, len(mapping))
	seen := make(map[string]bool)

	for _, currency := range mapping {
		if !seen[currency] {
			seen[currency] = true
			currencies = append(currencies, currency)
		}
	}

	return currencies
}
