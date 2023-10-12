package main

import "fmt"

func main() {

	//Matriz -> Equivalente a lista con longitud definida
	var matriz1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(matriz1)

	//Matriz con longitud indefinido
	var matriz2 = [...]int{6, 7, 8, 9}
	fmt.Println(matriz2)

	//Recorrer una matriz
	for index, value := range matriz1 {
		fmt.Printf("Indice: %d, Valor: %d\n", index, value)
	}

	//Forma tradicional de recorrer
	for i := 0; i < len(matriz2); i++ {
		fmt.Println(matriz2[i])
	}

	//Matriz bidimensional
	var matrizBi = [3][3]int{{12, 34, 52}, {1, 455, 7}, {0, 65, 23}}
	fmt.Println(matrizBi)

}
