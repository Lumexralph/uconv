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
	Long: `temperature is a sub-command for uconv (unit converter).

	It helps to convert temperature from one unit to another, Kelvin - k,
	Fahrenheit - f and Celsius - c.

	If no flags are specified, it converts Celsius - c to Fahrenheit - f by default.

	Usage:
		uconv temperature 45
		uconv temperature 45 --from f --to c
		uconv temperature 45 -f f -t c
	`,
	Args: cobra.ExactArgs(1),
	Run:  executeTemperatureCmd,
}

// executeTemperatureCmd - handles the temperature commands with respective flags
func executeTemperatureCmd(cmd *cobra.Command, args []string) {
	// convert the string to int
	arg, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch {
	// TODO: Make the flags caps insensitive
	case from == "c" && to == "f":
		fmt.Printf("temperature: %d°C ==> %d°F\n", arg, celsiusToFahrenheit(arg))
	case from == "f" && to == "c":
		fmt.Printf("temperature: %d°F ==> %d°C\n", arg, fahrenheitToCelsius(arg))
	case from == "c" && to == "k":
		fmt.Printf("temperature: %d°C ==> %dK\n", arg, celsiusToKelvin(arg))
	case from == "f" && to == "k":
		fmt.Printf("temperature: %d°F ==> %dK\n", arg, fahrenheitToKelvin(arg))
	case from == "k" && to == "c":
		fmt.Printf("temperature: %dK ==> %d°C\n", arg, kelvinToCelsius(arg))
	case from == "k" && to == "f":
		fmt.Printf("temperature: %dK ==> %d°F\n", arg, kelvinToFahrenheit(arg))
	case from == to:
		fmt.Printf("temperature: %d%s\n", arg, to)
	}
}

// celsiusToFahrenheit - converts Celsius to Fahrenheit
func celsiusToFahrenheit(data int64) int64 {
	return (data * 9 / 5) + 32
}

// celsiusToKelvin - converts Celsius to Kelvin
func celsiusToKelvin(data int64) int64 {
	return data + 273
}

// fahrenheitToCelsius - converts Fahrenheit to Celsius
func fahrenheitToCelsius(data int64) int64 {
	return (data - 32) * 5 / 9
}

// fahrenheitToKelvin - converts Fahrenheit to Kelvin
func fahrenheitToKelvin(data int64) int64 {
	return (data-32)*5/9 + 273
}

// kelvinToCelsius - converts Kelvin to Celsius
func kelvinToCelsius(data int64) int64 {
	return data - 273
}

// kelvinToFahrenheit - converts Kelvin to Fahrenheit
func kelvinToFahrenheit(data int64) int64 {
	return (data-273)*9/5 + 32
}
