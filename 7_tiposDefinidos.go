package main
/** Tipos definidos
En Go se tiene el concepto de estructuras y definición de nuevos tipos.

A los tipos se les pueden asignar operaciones para simular la programación orientada a objetos.
*/

// CREACIÓN DE UN NUEVO TIPO 
type listaEquipo []string

// Creación de un nuevo tipo con estructura.
// NOTAR COMO ES POSIBLE UTILIZAR LOS NUEVOS TIPOS
type salon struct {
	nombre		string
	miembros	[]listaEquipo
}


func main() {
	// CREANDO NUESTROS NUEVOS TIPOS
	var equipos listaEquipo = [["Manuel", "Alberto"], ["Sam", "Jessica"]]
	salon1 := salon{"COMP", listaEquipo()}
}