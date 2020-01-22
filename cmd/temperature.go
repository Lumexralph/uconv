package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tempCmd)
}

var tempCmd = &cobra.Command{
	Use:   "temperature",
	Short: "Convert temperature for different units",
	Long: `temp is a sub-command for uconv (unit converter),
			it helps to convert temperature from one unit to another, Kelvin,
			Fahrenheit and Celsius`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the temperature converter for the unit converter")
	},
}

// CelsiusToFahrenheit - converts Celsius to Fahrenheit
func CelsiusToFahrenheit(data int) int {
	return (data * 9 / 5) + 32
}

// CelsiusToKelvin - converts Celsius to Kelvin
// func CelsiusToKelvin(data int) int {

// }

// // FahrenheitToCelsius - converts Fahrenheit to Celsius
// func FahrenheitToCelsius(data int) int {

// }

// // FahrenheitToKelvin - converts Fahrenheit to Kelvin
// func FahrenheitToKelvin(data int) int {

// }

// // KelvinToCelsius - converts Kelvin to Celsius
// func KelvinToCelsius(data int) int {

// }

// // KelvinToFahrenheit - converts Kelvin to Fahrenheit
// func KelvinToFahrenheit(data int) int {

// }

// Reporter - handles the designation of temperature conversion
// to the appropriate converter and displays the result
func Reporter(data int, from, to string) (int, error) {
	switch {
	case from == "c" && to == "f":
		return CelsiusToFahrenheit(data), nil
	}
	return 0, errors.New("No valid temperature unit provided, should be k, c, f")
}
