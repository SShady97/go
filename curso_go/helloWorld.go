package main

import (
	"fmt"
)

// Declaracion de constantes
const pi float32 = 3.1415

const (
	x = 100
	y = 0b1010 // Valor binario
	z = 0o12   // Valor Octal
	w = 0xFF   // Valor Hex
)

const (
	domingo = iota + 1
	lunes
	martes
	miercoles
	jueves
	viernes
	sabado
)

func main() {
	// fmt.Println("Hello World!")
	// fmt.Println(quote.Go())

	// var (
	// 	firstName = "Jose"
	// 	lastName  = "Alvarez"
	// 	age       = 26
	// )

	// var firstName, lastName, age = "Jose", "Alvarez", 26

	//Sin usar var. Sólo se declaran así dentro de las funciones
	// firstName, lastName, age := "Jose", "Alvarez", 26

	// fmt.Println(firstName, lastName, age)
	// fmt.Println(pi)
	// fmt.Println(x, y, z, w)
	// fmt.Println(sabado)

	// var integer16 int16 = 50
	// var integer32 int32 = 100

	// fmt.Println(int32(integer16) + integer32)

	// Convertir String a Entero
	// s := "100"
	// i, _ := strconv.Atoi(s)
	// fmt.Println(i)

	// // Convertir Entero a String
	// n := 42
	// s2 := strconv.Itoa(n) // Itoa sólo recibe tipo int, no tipo int8, int32, etc.
	// fmt.Println(s2)

	// // Usando Printf
	// name := "Jose"
	// age := 26
	// fmt.Printf("Hola, me llamo %s y tengo %d años. \n", name, age)

	// // Usando Sprintf
	// greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.", name, age)
	// fmt.Println(greeting)

	// Ingresando datos desde consola a las variables
	var firstName, lastName string
	var age int
	fmt.Print("Ingrese su nombre:")
	fmt.Scanln(&firstName, &lastName)
	fmt.Print("Ingrese su edad:")
	fmt.Scanln(&age)

	fmt.Printf("Hola, soy %s %s y tengo %d años \n", firstName, lastName, age)
	fmt.Printf("Hola, soy %v %v y tengo %v años \n", firstName, lastName, age) // Se usa %v cuando no sé el tipo de dato a mostrar
	fmt.Printf("El tipo de dato de la variable age es: %T\n", age)             //Para conocer el tipo de dato de una variable

}
