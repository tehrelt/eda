package common_test

import (
	"testing"

	"github.com/tehrelt/eda/orders/internal/common"
)

func TestToCurrency(t *testing.T) {
	cases := []struct {
		name     string
		in       int64
		expected string
	}{
		{name: "positive", in: 10000, expected: "100.00"},
		{name: "negative", in: -10000, expected: "-100.00"},
		{name: "zero", in: 0, expected: "0.00"},
		{name: "cents", in: 99, expected: "0.99"},
		{name: "123.45", in: 12345, expected: "123.45"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := common.Currency(c.in)

			if got.String() != c.expected {
				t.Errorf("expected %s, got %s", c.expected, got.String())
			}
		})
	}
}
