package main

import "fmt"

func main() {

	fmt.Println("Mapas")
	mapa1 := map[string]int{"Uno": 1, "Dos": 2, "Tres": 3, "Cuatro": 4}
	fmt.Println(mapa1)

	// Agregar elemento a un mapa
	fmt.Println("\nAgregar elemento a un mapa")
	mapa1["Cinco"] = 5
	fmt.Println(mapa1)

	// Eliminar elemento de un mapa
	fmt.Println("\nEliminar elemento de un mapa")
	delete(mapa1, "Uno")
	fmt.Println(mapa1)

	//Obtener elemento y saber si existe
	//value, exist := mapa1["Dos"]
	//fmt.Println(value, exist)

	fmt.Println("\nAnalizando si existe una clave")
	if value, exist := mapa1["Uno"]; exist {
		fmt.Println(value)
	} else {
		fmt.Println("La clave no existe")
	}

	//Recorriendo un mapa
	fmt.Println("\nRecorriendo un mapa")

	for key, value := range mapa1 {
		fmt.Printf("Clave: %s, Valor: %d\n", key, value)
	}
}
