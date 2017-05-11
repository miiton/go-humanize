package humanize

import (
	"math"
	"testing"
)

func TestPostalCode(t *testing.T) {
	testList{
		{"0", PostalCode(0), "000-0000"},
		{"10", PostalCode(10), "000-0010"},
		{"100", PostalCode(100), "000-0100"},
		{"1000", PostalCode(1000), "000-1000"},
		{"10000", PostalCode(10000), "001-0000"},
		{"100000", PostalCode(100000), "010-0000"},
		{"1000000", PostalCode(1000000), "100-0000"},
		{"1234567", PostalCode(1234567), "123-4567"},
		{"10000000", PostalCode(10000000), "0"},
		{"10100000", PostalCode(10100000), "0"},
		{"10010000", PostalCode(10010000), "0"},
		{"10001000", PostalCode(10001000), "0"},
		{"123456789", PostalCode(123456789), "0"},
		{"maxint", PostalCode(9.223372e+18), "0"},
		{"math.maxint", PostalCode(math.MaxInt64), "0"},
		{"math.minint", PostalCode(math.MinInt64), "0"},
		{"minint", PostalCode(-9.223372e+18), "0"},
		{"-123,456,789", PostalCode(-123456789), "0"},
		{"-10,100,000", PostalCode(-10100000), "0"},
		{"-10,010,000", PostalCode(-10010000), "0"},
		{"-10,001,000", PostalCode(-10001000), "0"},
		{"-10,000,000", PostalCode(-10000000), "0"},
		{"-100,000", PostalCode(-100000), "0"},
		{"-10,000", PostalCode(-10000), "0"},
		{"-1,000", PostalCode(-1000), "0"},
		{"-100", PostalCode(-100), "0"},
		{"-10", PostalCode(-10), "0"},
	}.validate(t)
}
