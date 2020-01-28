package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var lengthFrom, lengthTo string

func init() {
	lengthCmd.Flags().StringVarP(&lengthFrom, "from", "f", "m", "the unit to convert from")
	lengthCmd.Flags().StringVarP(&lengthTo, "to", "t", "km", "the unit to convert to")

	rootCmd.AddCommand(lengthCmd)
}

var lengthCmd = &cobra.Command{
	Use:   "length",
	Short: "Convert length for different units",
	Long: `length is a sub-command for uconv (unit converter).

	It helps to convert length from one unit to another, Kilometer - km,
	Centimeter - cm, Meter - m, Millimeter - mm, Yard - y, Foot - f and Inches - in.

	If no flags are specified, it converts Meter - m to Kilometer - km by default.

	Usage:
		uconv length 45
		uconv length 45 --from m --to km
		uconv length 45 -f m -t km
	`,
	Args: cobra.ExactArgs(1),
	Run:  executeLengthCmd,
}

// executeTemperatureCmd - handles the temperature commands with respective flags
func executeLengthCmd(cmd *cobra.Command, args []string) {
	// convert the string to int
	arg, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch {
	// TODO: Make the flags caps insensitive
	case lengthFrom == "m" && lengthTo == "km":
		fmt.Printf("length: %dM ==> %dKM\n", arg, meterToKilometer(arg))
	case lengthFrom == "km" && lengthTo == "m":
		fmt.Printf("length: %dKM ==> %dM\n", arg, kilometerToMeter(arg))
	case lengthFrom == "m" && lengthTo == "cm":
		fmt.Printf("length: %dM ==> %dCM\n", arg, meterToCm(arg))
	case lengthFrom == "cm" && lengthTo == "m":
		fmt.Printf("length: %dCM ==> %dM\n", arg, cmToMeter(arg))
	}
}

func meterToKilometer(data int64) int64 {
	return data / 1000
}

func kilometerToMeter(data int64) int64 {
	return data * 1000
}

func cmToMeter(data int64) int64 {
	return data / 100
}

func meterToCm(data int64) int64 {
	return data / 100
}
