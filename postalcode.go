package humanize

import "fmt"

// PostalCode produces a string form of the given number in
// Japanese postal code format
//
// e.g. PostalCode(1234567) -> 123-4567
func PostalCode(v int) string {
	if v > 9999999 {
		return "0"
	} else if v < 0 {
		return "0"
	}
	pre := v / 10000 % 1000
	suf := v % 10000
	return fmt.Sprintf("%03d-%04d", pre, suf)
}
