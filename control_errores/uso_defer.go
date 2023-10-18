package main

import (
	"fmt"
	"os"
)

func main() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola, Jose Alvarez"))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
