package main

import "fmt"

func convertCelsiusToFahrenheit(c float64) float64 {
	return c*9.0/5.0 + 32.0
}

func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32.0) * 5.0 / 9.0
}

func main() {
	c1 := 0.0
	c2 := 37.0
	c3 := -40.0

	f1 := convertCelsiusToFahrenheit(c1)
	f2 := convertCelsiusToFahrenheit(c2)
	f3 := convertCelsiusToFahrenheit(c3)

	fmt.Printf("%.1f°C -> %.1f°F\n", c1, f1)
	fmt.Printf("%.1f°C -> %.1f°F\n", c2, f2)
	fmt.Printf("%.1f°C -> %.1f°F\n", c3, f3)

	fmt.Printf("%.1f°F -> %.1f°C\n", f1, convertFahrenheitToCelsius(f1))
	fmt.Printf("%.1f°F -> %.1f°C\n", f2, convertFahrenheitToCelsius(f2))
	fmt.Printf("%.1f°F -> %.1f°C\n", f3, convertFahrenheitToCelsius(f3))
}
