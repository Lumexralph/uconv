package main

import (
	"github.com/Lumexralph/unit-converter/temperature"
)

func main() {
	if v, err := temperature.Reporter(30, "c", "f"); err != nil {
		println(err)
	} else {
		println(v, "°F")
	}
}
