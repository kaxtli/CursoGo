/** Declaración de paquetes

Un paquete es un contenedor de código y funcionalidades de nuetro programa.
Existe un paquete especial llamado 'main' que es el que será ejecutado como una aplicación.
*/
package main

/** Importación de paquetes

Go permite reutilizar paquetes, descargar paquetes nuevos o utilizar los definidos en la
librería estándar.
*/
import "fmt"

/** Función main

La función main es una función especial, y sólo puede existir una por ejecutable.
Esta será la función que se ejecutará al lanzar nuestro programa.
*/
func main() {
	/** Cuerpo
	
	Dentro de {   } se cre un nuevo 'bloque'. Aquí es donde agregamos el comportamiento deseado
	de nuetro código
	 */
	 fmt.Println("Hola Gophers!")
}