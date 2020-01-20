package temperature

import "errors"

// Reporter - handles the designation of temperature conversion
// to the appropriate converter and displays the result
func Reporter(data int, from, to string) (int, error) {
	switch {
	case from == "c" && to == "f":
		return CelsiusToFahrenheit(data), nil
	}
	return 0, errors.New("No valid temperature unit provided, should be k, c, f")
}
