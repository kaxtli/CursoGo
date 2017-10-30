/** Arreglos

Un arreglo es una estructura de datos que puede contener varios valores de un mismo tipo.
Un ejemplo de un arreglo sería un conjunto de estudiantes en un equipo.
Su longitud no puede variar durante la ejecución del programa.
*/

package main

import "fmt"

func main() {
	// Declaración de un arreglo con cinco componentes pre-declarados
	var equipo1 [5]string = [5]string{"Luis", "Jimena", "Ana", "Pablo", "Sam"}

	// Declaración de un arreglo con valores por defecto
	var calificacionesEquipo1 [5]float32

	// Podemos acceder a un arreglo por su índice:
	// arreglo[pos] Posición en específico
	// arreglo[inicio:fin] Desde inicio (incluyendo) hasta fin (no incluyendo)
	// arreglo[:fin] Desde el inicio del arreglo hasta fin (no incluyendo)
	// arreglo[inicio:] Desde inicio (incluyendo) hasta el final del arreglo
	fmt.Println("Lider", equipo1[0], "apoyo", equipo1[1:3], "Itegrantes ", len(equipo1))
	
	// Iterando sobre un arreglo
	for posicion, valor := range equipo1 {
		fmt.Println(valor, calificacionesEquipo1[posicion])
	}

	// Asignando valores a un array
	for posicion, valor := range equipo1 {
		fmt.Println("Ingrese la calificación de", valor)
		fmt.Scan(&calificacionesEquipo1[posicion])
	}

	for posicion, valor := range equipo1 {
		fmt.Println(valor, calificacionesEquipo1[posicion])
	}
	/************ SLICES
	Los 'slice' son una alternativa dinámica a los arreglos. Pueden modificar su longitud para
	los datos que sean necesarios. Consiste en una estructura que trabaja sobre un arreglo.
	NOTA: Un slice es un puntero.
	*/
	var equipo2 []string // Un slice vacío

	var salida bool salida = false
	var respuesta string

	for !salida {
		fmt.Println("NUevo integrante del equipo 2 (q para salir): ")
		fmt.Scan(&respuesta)

		if respuesta == "q" || respuesta == "Q" {
			salira = true
			fmt.Println("¡Listo!")
		} else {
			// La función append agrega un elemento a un slice y devuelve el slice aumentado
			// NO MODIFICA EL SLICE ORIGINAL
			equipo2 = append(equipo2, respuesta)
		}
	}

	//-- EJERCICIO: Definir un slice que contenga las calificaciones para el equipo 2, recibe 
	//-- el resultado del usuario y posteriormente muestra los resultados como en ejemplos anteriores.

	// TU CODIGO 

	// FIN DE TU CODIGO
}