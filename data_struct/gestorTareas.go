package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (lt *ListaTareas) agregarTarea(t Tarea) {
	lt.tareas = append(lt.tareas, t)
}

func (l *ListaTareas) completarTarea(index int) {
	l.tareas[index-1].completado = true
}

func (l *ListaTareas) editarTarea(index int, tarea Tarea) {
	l.tareas[index-1] = tarea
}

func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index-1], l.tareas[index:]...)
}

func (l *ListaTareas) listarTareas() {
	fmt.Println("======================================================")
	for i, t := range l.tareas {
		fmt.Printf("%d. %s - %s - Completado: %t\n", i+1, t.nombre, t.desc, t.completado)
	}
	fmt.Println("======================================================\n")
}

func main() {

	lista := ListaTareas{}

	var opcion int

	//Instancia de bufio para entrada de datos
	leer := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Seleccione una de las siguientes opciones:\n",
			"1. Agregar tarea.\n",
			"2. Marcar tarea como completada.\n",
			"3. Editar tarea.\n",
			"4. Eliminar tarea.\n",
			"5. Listar tareas.\n",
			"6. Salir.")

		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Ingrese el nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea:")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("La tarea fue agregada con éxito!\n")
		case 2:
			var index int
			fmt.Println("Ingrese el numero de la tarea que quiere marcar como completada:")
			fmt.Scanln(&index)
			lista.completarTarea(index)
			fmt.Printf("La tarea %d fue marcada como completada!\n", index)
		case 3:
			var index int
			var t Tarea
			fmt.Println("Ingrese el numero de la tarea que quiere editar:")
			fmt.Scanln(&index)
			fmt.Println("Ingrese el nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea:")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Printf("La tarea %d fue editada con éxito!\n", index)
		case 4:
			var index int
			fmt.Println("Ingrese el numero de la tarea que quiere eliminar:")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Printf("La tarea %d fue eliminada con éxito!\n")
		case 5:
			fmt.Println("Listado de tareas:")
			lista.listarTareas()
		case 6:
			fmt.Println("Saliendo del programa.....")
			return
		default:
			fmt.Println("Opción incorrecta!")
		}
	}
}
