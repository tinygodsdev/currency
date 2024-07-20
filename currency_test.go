package currency

import (
	"testing"
)

func TestNormalizeCurrency(t *testing.T) {
	tests := []struct {
		input         string
		customMapping map[string]string
		expected      string
		expectedOK    bool
	}{
		{"₽", nil, "RUB", true},
		{"руб.", nil, "RUB", true},
		{"$", nil, "USD", true},
		{"usd", nil, "USD", true},
		{"£", nil, "GBP", true},
		{"A$", nil, "AUD", true},
		{"HK$", nil, "HKD", true},
		{"ZAR", nil, "ZAR", true},
		{"unknown", nil, "unknown", false},
		{"custom", map[string]string{"custom": "CSTM"}, "CSTM", true},
		{"usd", map[string]string{"usd": "USDT"}, "USDT", true},
		{"$", map[string]string{"$": "DOLLAR"}, "DOLLAR", true},
	}

	for _, test := range tests {
		result, ok := NormalizeCurrency(test.input, test.customMapping)
		if result != test.expected || ok != test.expectedOK {
			t.Errorf("NormalizeCurrency(%q, %v) = %q, %v; want %q, %v",
				test.input, test.customMapping, result, ok, test.expected, test.expectedOK)
		}
	}
}

func TestSupportedCurrencies(t *testing.T) {
	expected := []string{
		"RUB", "EUR", "USD", "GBP", "ITL", "JPY", "KRW", "INR", "UAH", "ILS",
		"AUD", "CAD", "CHF", "CNY", "HKD", "NZD", "SEK", "SGD", "MXN", "NOK", "BRL", "ZAR",
	}

	result := SupportedCurrencies()

	if len(result) != len(expected) {
		t.Errorf("SupportedCurrencies() = %d currencies; want %d", len(result), len(expected))
	}

	seen := make(map[string]bool)
	for _, currency := range result {
		seen[currency] = true
	}

	for _, currency := range expected {
		if !seen[currency] {
			t.Errorf("SupportedCurrencies() missing %q", currency)
		}
	}
}
