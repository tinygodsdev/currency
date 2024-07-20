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
	}
}

func NormalizeCurrency(currency string, customMapping map[string]string) (string, bool) {
	// if custom mapping is provided, use it
	if customMapping != nil {
		if v, ok := customMapping[currency]; ok {
			return v, true
		}
	}

	// if no custom mapping is provided, use the default mapping
	mapping := getCurrencyMapping()
	if v, ok := mapping[strings.ToLower(currency)]; ok {
		return v, true
	}

	return currency, false
}
