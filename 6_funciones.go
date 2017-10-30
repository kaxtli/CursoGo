package main
/** Funciones

Una función es una manera de agrupar código de Go en un mismo bloque.

Go implementa dos tipos de funciones:
//- Nombradas
//- Anónimas
*/

import(
	"fmt"
	"time"
)

// Función nombrada sin parámetros
func mostrarHora() {
	currentTime := time.Now().Local()
	fmt.Println("La hora es: ", currentTime.Format("2006-01-02"))
}

// Functión nombrada con parámetros
func mostrarHoraCadena(formato string) {
	currentTime := time.Now().Local()
	fmt.Println("La hora es: ", currentTime.Format(formato))
}

// Función nombrada con parámetros y retorno
func mostrarHoraDevuelveCadena(formato string) string {
	currentTime := time.Now().Local()
	return currentTime.Format(formato)
}

// Función nombrada con parámetros y nombrado del valor de retorno.
func mostrarHoraDevueveNombrado(formato string)(fecha string) {
	currentTime := time.Now().Local()
	fecha = currentTime.Format(formato)
	return
	// NOTA: Terminar indicando las variables a retornar es opcional
	// Tampoco es necesario declarar la variable currentTime.
}

// Función con múltiple valor de retorno.
func multipleReturn(formato string) (fecha string, ok bool){
	if formato == "" {
		fecha = ""
		ok = false
		return
	}
	currentTime := time.Now().Local()
	fecha = currentTime.Format(formato)
	ok = true
	return
}

// Función con punteros
func agregarMiembro(nuevo string, equipo *[]string, miembros *int) {
	*equipo = append(*equipo, nuevo)
	*miembros++
}

func main() {
	// Invocando funciones definidas
	mostrarHora()

	mostrarHoraCadena("2006/01/02")

	fecha := mostrarHoraDevuelveCadena("2006|01|02")
	fecha2 := mostrarHoraDevueveNombrado("2006.01.02")

	fmt.Println(fecha)
	fmt.Println(fecha2)

	// Múltiples valores de retorno.
	fecha3, ok := multipleReturn("")

	if ok {
		fmt.Println(fecha3)
	} else {
		fmt.Println("Cadena de formato vacía")
	}

	// Funciones con punteros
	equipo := []string{"Juan", "Manuel", "Jessica"}
	integrantes := len(equipo)

	fmt.Println(equipo, integrantes)
	agregarMiembro("Sam", &equipo, &integrantes)
	fmt.Println(equipo, integrantes)


	

}