package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var to, from string

func init() {
	tempCmd.Flags().StringVarP(&from, "from", "f", "c", "the unit to convert from")
	tempCmd.Flags().StringVarP(&to, "to", "t", "f", "the unit to convert to")

	rootCmd.AddCommand(tempCmd)
}

var tempCmd = &cobra.Command{
	Use:   "temperature",
	Short: "Convert temperature for different units",
	Long: `temp is a sub-command for uconv (unit converter),
	it helps to convert temperature from one unit to another, Kelvin - k,
	Fahrenheit - f and Celsius - c`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// convert the string to int
		arg, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}

		result := celsiusToFahrenheit(arg)
		// °F
		fmt.Println("I need to convert this unit: ", result, "°F")
	},
}

// CelsiusToFahrenheit - converts Celsius to Fahrenheit
func celsiusToFahrenheit(data int64) int64 {
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
// func Reporter(data int, from, to string) (int, error) {
// 	switch {
// 	case from == "c" && to == "f":
// 		return CelsiusToFahrenheit(data), nil
// 	}
// 	return 0, errors.New("No valid temperature unit provided, should be k, c, f")
// }
