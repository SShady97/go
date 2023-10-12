package main

import "fmt"

func main() {

	// Definir una rebanada
	reb1 := []int{1, 2, 3, 4, 5}

	// Imprimir contenido
	fmt.Println("Contenido Rebanada1: ", reb1)
	// Imprimir logitud
	fmt.Println("Rebanada1 - Longitud: ", len(reb1))
	// Imprimir capacidad
	fmt.Printf("Rebanada1 - Capacidad: %d \n\n", reb1)

	// Definir rebanada con longitud y capacidad. Se inicializa en 0's
	reb2 := make([]int, 6, 10)

	// Imprimir contenido
	fmt.Println("Contenido Rebanada2: ", reb2)
	// Imprimir logitud
	fmt.Println("Rebanada2 - Longitud: ", len(reb2))
	// Imprimir capacidad
	fmt.Printf("Rebanada2 - Capacidad: %d \n\n", cap(reb2))

	// Obtener un fragmento de una rebanada (otra rebanada)
	reb3 := []string{"Otra cancion larga", "Tararea", "Sancocho en leña", "La tipica", "Aranjuez"}
	fmt.Println("Rebanada3: ", reb3)
	reb4 := reb3[:2]
	fmt.Println("Rebanada4: ", reb4)
	// Capacidad de la nueva rebanada (reb4)
	fmt.Printf("Capacidad de Rebanada4: %d \n\n", cap(reb4)) //Conserva la capacidad de la rebanada de donde se hizo la extracción

	// Agregar un elemento a una rebanada
	fmt.Println("Agregar un elemento a una rebanada")
	fmt.Println("Capacidad Rebanada3: ", cap(reb3))
	reb3 = append(reb3, "Los legendarios")
	fmt.Println(reb3)
	fmt.Println("Nueva capacidad Rebanada3: ", cap(reb3)) // Al sobrepasar la capacidad inicial que era 5 la capacidad se duplica dinamicamente

	// Agregar un elemento a una rebanada
	fmt.Println("\nAgregar n elementos a una rebanada")
	reb3 = append(reb3, "Reflujo", "Fruko y sus presos", "El malo de la pelicula")
	fmt.Println("Rebanada3: ", reb3)

	// Eliminar un elemento de una rebanad
	fmt.Println("\nEliminar elemento de una rebanada")
	fmt.Println("Rebanada3: ", reb3)
	reb3 = append(reb3[:3], reb3[4:]...) //Se elimina el elemento en la posición 3
	fmt.Println("Posición 3 eliminada: ", reb3)

	// Copiar rebanadas
	fmt.Println("\nCopiar rebanadas")
	fmt.Println("Rebanada2: ", reb2)
	copy(reb2, reb1) // Copia la rebanada 1 en la 2
	fmt.Println("Rebanada2 despues: ", reb2)

	// Reemplazar elemento de una rebanada
	fmt.Println("\nReemplazar elemento de una rebanada")
	reb2[5] = 6
	fmt.Println(reb2)

}
