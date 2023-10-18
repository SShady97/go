package main

import (
	"fmt"
	"strconv"
)

// Error personalizado1
func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		// Usando errors
		//return 0, errors.New("No se puede dividir por cero!")

		//Usando fmt
		return 0, fmt.Errorf("No se puede dividir por cero!")
	}
	return dividendo / divisor, nil
}

func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Número: ", num)

	result, err := divide(15, 0)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Resultado: ", result)

}
