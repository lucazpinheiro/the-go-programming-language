// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"flag"
	"fmt"
	"strings"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func showResult(initial, final string) {
	fmt.Println("From " + initial + " to " + final)
}

var inputUnit = flag.String("i", "c", "temperature unit, options: C, F, K. default = Celsius")
var outputUnit = flag.String("o", "k", "temperature unit, options: C, F, K. default = Fahrenheit")
var temperature = flag.Float64("t", 24, "temperature")

func main() {
	flag.Parse()
	input := *inputUnit
	output := *outputUnit
	temperature := *temperature

	switch strings.ToLower(input) {
	case "f":
		f := Fahrenheit(temperature)
		if strings.ToLower(output) == "c" {
			c := FToC(f)
			showResult(f.String(), c.String())
		}
		if strings.ToLower(output) == "k" {
			k := CToK(FToC(f))
			showResult(f.String(), k.String())
		}
	case "k":
		k := Kelvin(temperature)
		if strings.ToLower(output) == "c" {
			c := KToC(k)
			showResult(k.String(), c.String())
		}
		if strings.ToLower(output) == "f" {
			f := CToF(KToC(k))
			showResult(k.String(), f.String())
		}
	default:
		c := Celsius(temperature)
		if strings.ToLower(output) == "k" {
			k := CToK(c)
			showResult(c.String(), k.String())
		}
		if strings.ToLower(output) == "f" {
			f := CToF(c)
			showResult(c.String(), f.String())
		}
	}
}
