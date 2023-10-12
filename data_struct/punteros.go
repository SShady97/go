package main

import "fmt"

type Celular struct {
	marca  string
	modelo string
	ram    int
	precio int
}

func (c *Celular) mostrarDatos() {
	fmt.Printf("El celular es de marca %s, es un %s, su memoria RAM es de %d GB y su precio es de %d dólares. \n",
		c.marca, c.modelo, c.ram, c.precio)
}

func editarX(x *int) {
	*x = 25
}

func main() {

	celular := Celular{"Apple", "iPhone 12", 4, 700}
	celular.mostrarDatos()

	//Trabajando con punteros
	x := 15
	var y *int = &x
	//Imprime el valor de x y su referencia en memoria que fue almacenada en la variable y
	fmt.Println(x, y)

	//Modificando el valor de x
	editarX(&x)
	fmt.Println(x)
}
