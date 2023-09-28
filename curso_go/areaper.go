package main

import (
	"fmt"
	"math"
)

const precision = 2

func main() {

	var cateto1, cateto2 float64
	var hip float64

	fmt.Print("Ingrese el lado 1 del triangulo rectángulo:")
	fmt.Scanln(&cateto1)
	fmt.Print("Ingrese el lado 2 del triangulo rectángulo:")
	fmt.Scanln(&cateto2)

	sumaCatetos := math.Pow(cateto1, 2) + math.Pow(cateto2, 2)
	hip = math.Sqrt(sumaCatetos)

	area := (cateto1 * cateto2) / 2
	per := cateto1 + cateto2 + hip

	fmt.Printf("El area del triangulo es: %.*f\n", precision, area) // Usando constante de precision
	fmt.Printf("El perímetro del triangulo es: %.2f\"\n", per)

}
