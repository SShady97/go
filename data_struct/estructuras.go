package main

import "fmt"

type Persona struct {
	nombre    string
	edad      int
	ocupacion string
}

func main() {

	//Forma 1 de crear una instancia y darle valor a sus atributos
	var persona Persona
	persona.nombre = "Jose"
	persona.edad = 26
	persona.ocupacion = "Ingeniero"

	fmt.Println(persona)

	//Forma 2 de crear una instancia y darle valores a sus atributos
	persona2 := Persona{"Manuel", 32, "Fisico"}
	fmt.Println(persona2)

}
