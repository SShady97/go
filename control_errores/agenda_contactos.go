package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contacto struct {
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"`
}

func guardarContacto(contactos []Contacto) error {
	file, err := os.Create("contactos.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contactos)
	if err != nil {
		return err
	}
	return nil
}

func cargarContactos(contactos *[]Contacto) error {
	file, err := os.Open("contactos.json")
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contactos)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var contactos []Contacto
	var opcion int

	//Cargar contactos existentes del archivo
	err := cargarContactos(&contactos)
	if err != nil {
		fmt.Println("Ha ocurrido un error al cargar el archivo de contactos: ", err)
	}

	// Crear instancia de bufio
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n================ Gestor de contactos ================\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opción:")

		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Error al leer la opción:", err)
		}

		switch opcion {
		case 1:
			var c Contacto
			fmt.Print("\nNombre:")
			c.Nombre, _ = reader.ReadString('\n')
			fmt.Print("Email:")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Teléfono:")
			c.Telefono, _ = reader.ReadString('\n')

			//Agregar contacto a Slice
			contactos = append(contactos, c)

			//Guardar en archivo json
			err = guardarContacto(contactos)
			if err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			}
		case 2:
			fmt.Println("========================================")
			for index, contact := range contactos {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Nombre, contact.Email, contact.Telefono)
			}
			fmt.Println("========================================")
		case 3:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Escoja una opción válida")
		}
	}

}
