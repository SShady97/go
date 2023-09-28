package main

import "fmt"

func main() {

	// var i int

	// // Bucle con condición -> Equivalente a While
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// // Bucle for tipico con break
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// Bucle for con continue
	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		continue // No se ejecuta lo de abajo y pasa a la siguiente iteración
	// 	}
	// 	fmt.Println(i)
	// }

	// saludo := hello("Jose")
	// fmt.Println(saludo)

	suma, mult := calc2(5, 4)

	fmt.Println(suma)
	fmt.Println(mult)

}

func hello(name string) string { // Funcion que retorna string
	return "Hola, " + name
}

func calc1(a, b int) (int, int) { // Funcion que retorna dos enteros
	sum := a + b
	mul := a * b
	return sum, mul
}

// Otra forma de la función
func calc2(a, b int) (sum, mul int) { // Funcion que retorna dos enteros
	sum = a + b
	mul = a * b
	return
}
