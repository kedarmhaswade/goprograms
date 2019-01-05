// Demonstrates the use of package creation (gopl.io chapter 2) for temperature conversion
package tempconv

import "fmt"

// type name underlying-type-name
type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
	AbsoluteZero Kelvin = 0
)

// methods on types

func (c Celsius) String() string  {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string  {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string  {
	return fmt.Sprintf("%gK", k)
}