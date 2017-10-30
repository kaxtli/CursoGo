package main
/** Bloques condicionales e iteraciones

En Go la estructura condicional básica es el 'if'.
Para la selección entre múltiples opciones se puede usar 'switch'.

En Go no existen ni while ni do... while. En su lugar existen
tres estilos de iteraciones.
*/

import "fmt"

func main() {
	var edad int = 34
	var permiso bool = false

	// Condicional simple con bloque else
	if edad >= 18 {
		fmt.Println("Eres mayor de edad.")
	} else if edad > 12 && permiso {
		fmt.Println("No eres mayor de edad, pero puedes pasar.")
	} else {
		fmt.Println("No puedes pasar")
	}

	/** Estructura de selección switch.
	Existen dos maneras de switch: Una de comparación y otra libre.
	Por defecto, al encontrar el bloque apropiado SÓLO EJECUTARÁ 
	DICHO BLOQUE por lo que no es necesario agregar una instrucción
	'break' al final de cada bloque.

	Si se desea obtener el comportamiento similar a C, se usa
	la palabra reservada  "fallthrough"
	*/

	// Primer forma de switch: Comparación de valores.
	// En esta forma se compara la variable dada con un listado específico.
	switch edad {
	case 18:
		fmt.Println("Recién eres mayor de edad")
	case 15:
		fmt.Println("Recién eres adolescente")
	default:
		fmt.Println("Estás en otro rango")
	}

	// Segunda forma de switch: Estructura libre.
	// En esta forma, en cada 'case' se puede poner cualquier condición
	switch {
	case edad > 60:
		fmt.Println("Tienes descuento")
		fallthrough
	case edad > 18:
		fmt.Println("Puedes entrar")
	default:
		fmt.Println("No puedes entrar")
	}


	/** Estructura de iteración for

	En Go, a diferencia de otros lenguajes comunes no existe ni while ni do... while. 
	En su lugar, for puede tomar cuatro formas distintas de iteraciones:
	*/

	// Iteración infinita. Equivalente a while(true) o for(;;)
	// Es necesario controlar su salida por errores o una instrucción 'break'.
	var usuario int
	for {
		fmt.Println("Inserte un número:")
		fmt.Scanf("%d", &usuario)

		if usuario == -1 {
			break
		} else {
			fmt.Println(usuario)
		}
	}

	// Iteración con condición. Equivalente a while(condicion)
	// Es necesario controlar su salida por errores, instrucciones break o 
	// modificando la información en la condición.

	for usuario < 10 {
		fmt.Println("Inserte su ID:")
		fmt.Scanf("%d", &usuario)

		if usuario == -1 {
			fmt.Println("ID inválido.")
			break
		} else {
			fmt.Println(usuario)
		}
	}

	// Iteración con contador. Equivalente a for(inicial; condicion; incremento)
	// Además de controlar su salida por errores o instrucciones break, automáticamente
	// incrementara o decrementará el valor en 'inicial'.
	// Aunque es posible modificar la variable 'inicial' directamente, no es recomendable
	fmt.Println("IDs NO válidos: ")
	for id := 0; id < 10; id++ {
		fmt.Printf("%d ", id)
	}
	fmt.Println()
}