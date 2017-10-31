package main

/** Tipos de usuario definidos

En Go la programación orientada a objetos se realiza a través de la definición
y composición de estructuras, así como a la asignación de operaciones.
*/

import "fmt"

// Creación de un nuevo tipo a partir de un alias
type Celsius float64

// Creación de un nuevo tipo a partir de una estructura
type Punto struct {
	X	int
	Y 	int
}

// Composición de una estructura dentro de otra (nombrado)
type Circulo struct {
	Centro 	Punto
	Radio	float64
}

// Composición de una estructura dentro de otra (sin nombre)
// equivalente a la herencia
type Llanta struct {
	Circulo
	Marca	string
}

// Declaración de un método para Circulo
func (c Circulo) Area() float64 {
	return c.Radio * 3.1416 * 3.1416
}

// Declaración de un método para Llanta (String)
func (l Llanta) String() string {
	cadena := fmt.Sprintf("Llanta %s en (%d, %d).", l.Marca, l.Centro.X, l.Centro.Y)
	return cadena
}

// Declaración de un método con puntero. Necesario para poder modificar elementos
// dentro de la estructura.
func (l *Llanta) Mover(x, y int) {
	l.Centro.X = x 
	l.Centro.Y = y 
}

func main() {
	// Creando estructuras.
	circulo := Circulo{Punto{10, 10}, 5.4}

	// Aunque sea un tipo compuesto, es necesario declarar cada estructura por separado.
	miLlanta := Llanta{Circulo{Punto{14, 10}, 34.5}, "Goodyear"}

	// Como Llanta contiene a Circulo, no es necesario hacer referencia a el ya que no se
	// le asignó un nombre.
	fmt.Println(miLlanta.Area())
	fmt.Println(circulo.Area())

	// Al implementar el método String en Llanta, ahora se puede imprimir directamente.
	fmt.Println(miLlanta)

	// El método Mover fue definido con un puntero a la estructura, por lo que los valores
	// pueden ser modificados.
	miLlanta.Mover(5, 5)

	fmt.Println(miLlanta)
}