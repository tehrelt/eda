package common

import "fmt"

type Currency int64

func (c Currency) Float64() float64 {
	return (float64(c) / 100)
}

func (c Currency) String() string {
	val := (float64(c) / 100)
	return fmt.Sprintf("%.2f", val)
}

func (c Currency) Multiply(f float64) Currency {
	val := (float64(c) * f) + 0.5
	return Currency(val)
}
