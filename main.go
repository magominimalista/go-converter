package main

import (
	"fmt"
)

func main() {
	var kelvin float64
	fmt.Print("Digite a temperatura em Kelvin: ")
	_, _ = fmt.Scanf("%f", &kelvin)

	celsius := kelvinParaCelsius(kelvin)
	fmt.Printf("A temperatura em Celsius é: %.2f\n", celsius)
}

func kelvinParaCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
