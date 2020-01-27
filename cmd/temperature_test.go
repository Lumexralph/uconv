package cmd

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	cases := []struct {
		name        string
		input, want int64
	}{
		{"positive data provided", 60, 140},
		{"negative data provided", -60, -76},
	}

	for _, tc := range cases {
		got := celsiusToFahrenheit(tc.input)

		t.Run(tc.name, func(t *testing.T) {
			if got != tc.want {
				t.Errorf("celsiusToFahrenheit(%d) got=%d; want=%d", tc.input, got, tc.want)
			}
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	cases := []struct {
		name        string
		input, want int64
	}{
		{"positive data provided", 60, 333},
		{"negative data provided", -60, 213},
	}

	for _, tc := range cases {
		got := celsiusToKelvin(tc.input)

		t.Run(tc.name, func(t *testing.T) {
			if got != tc.want {
				t.Errorf("celsiusToKelvin(%d) got=%d; want=%d", tc.input, got, tc.want)
			}
		})
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	cases := []struct {
		name        string
		input, want int64
	}{
		{"positive data provided", 60, 15},
		{"negative data provided", -60, -51},
	}

	for _, tc := range cases {
		got := fahrenheitToCelsius(tc.input)

		t.Run(tc.name, func(t *testing.T) {
			if got != tc.want {
				t.Errorf("fahrenheitToCelsius(%d) got=%d; want=%d", tc.input, got, tc.want)
			}
		})
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	cases := []struct {
		name        string
		input, want int64
	}{
		{"positive data provided", 60, 288},
		{"negative data provided", -60, 222},
	}

	for _, tc := range cases {
		got := fahrenheitToKelvin(tc.input)

		t.Run(tc.name, func(t *testing.T) {
			if got != tc.want {
				t.Errorf("fahrenheitToKelvin(%d) got=%d; want=%d", tc.input, got, tc.want)
			}
		})
	}
}

func TestKelvinToCelsius(t *testing.T) {
	cases := []struct {
		name        string
		input, want int64
	}{
		{"positive data provided", 60, -213},
		{"negative data provided", -60, -333},
	}

	for _, tc := range cases {
		got := kelvinToCelsius(tc.input)

		t.Run(tc.name, func(t *testing.T) {
			if got != tc.want {
				t.Errorf("kelvinToCelsius(%d) got=%d; want=%d", tc.input, got, tc.want)
			}
		})
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	cases := []struct {
		name        string
		input, want int64
	}{
		{"positive data provided", 60, -351},
		{"negative data provided", -60, -567},
	}

	for _, tc := range cases {
		got := kelvinToFahrenheit(tc.input)

		t.Run(tc.name, func(t *testing.T) {
			if got != tc.want {
				t.Errorf("kelvinToFahrenheit(%d) got=%d; want=%d", tc.input, got, tc.want)
			}
		})
	}
}
