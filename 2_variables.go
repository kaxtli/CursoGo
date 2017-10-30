/** Declaración, uso y convenciones de variables

Una variable puede ser vista como un contenedor de información en nuestro programa.
Al igual que otros lenguajes de programación como C/C++ o Java, Go posee tipado estático:

	Al incializar una variable con un tipo de dato en específico, este no puede variar.
*/
package main

import(
	"fmt"
	"strings"
)

/** Constantes
Las constantes son útiles para definir valores que no cambiarán a lo largo del programa.
Son de acceso más eficiente y suelen ir al nivel del paquete.
*/
const participante = "Jimena"
const numeroDeLaSuerte float32 = 12

func main() {

	fmt.Println(participante, numeroDeLaSuerte)

	// Declaración de variables con un valor inicial
	var entero int = 10
	var flotante float32 = 10.121
	var cadena string = "hola"
	var booleano bool = true
	var caracter rune = '😊'

	fmt.Println(entero, flotante, cadena, booleano, string(caracter))

	// Declaración de variables sin un valor inicial
	// Al no tener un valor inicial, toman un valor por defecto.
	// ¡En Go no existen las variables vacías!
	var enteroDefault int
	var flotanteDefault float32
	var cadenaDefault string
	var booleanoDefault bool
	var caracterDefault rune

	fmt.Println(enteroDefault, flotanteDefault, cadenaDefault, booleanoDefault, string(caracterDefault))

	// Inferencia de tipo 
	// Usando el operador := es posible inferir el tipo a partir de la asignación.
	nombre := "Juan"
	edad := 12
	salario := 12341.1

	var localidad = "Tepatitlán"

	fmt.Println(nombre, edad, "años, gana", salario, "en", localidad)

	/** PUNTEROS
	
	Al igual que otros lenguajes de programación Go tiene el concepto de punteros

	Un puntero es una variable que almacena la dirección de memoria donde se ubica 
	un valor.

	UTILIDAD
	+ Valor por referencia. Es más barato utilizar referencias que copias de valores.
	+ Modificación de valores. Si una función recibe un puntero como parámetro, 
		toda modificación que se haga a dicha variable se realizará también en la original.
	
	DESVENTAJAS
	- Complejidad añadida al programa.
	- Más suceptible a bugs extraños.
	- No es recomendable para novatos.
	*/

	var miCadena string = "Hola mundo"

	var puntero *string = &miCadena

	fmt.Println("Dirección de memoria: ", puntero, "Valor: ", *puntero)

	// Como 'puntero' es una referencia a 'miCadena', toda modificación que
	// se realice desde 'puntero' se verá reflejada en 'miCadena'.
	fmt.Println(miCadena)

	*puntero = strings.ToUpper(*puntero)

	fmt.Println(miCadena)

	// Es posible cambiar la referencia de un puntero

	fmt.Println(puntero, *puntero)

	var miCadena2 string = "Adios Mundo"
	puntero = &miCadena2

	fmt.Println(puntero, *puntero)
}