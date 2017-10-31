package main
/** Apertura de un archivo

En Go abrir un archivo es una tarea relativamente sencilla. En principio es una tarea 
de tres pasos: 
- Abrir el archivo
- Leer el archivo
- Procesar la información leida.
*/


import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {

	// Abrir un archivo.
	file, err := os.Open("texto.txt")
	// Revisión de errores
    if err != nil {
        log.Fatal(err)
	}
	// Cerrar el archivo al final
    defer file.Close()

	// Abrir un nuevo Scanner (salto de línea por defecto)
	scanner := bufio.NewScanner(file)
	
	// Iterar por cada una de las líneas
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

	// Verificar errores durante el proceso
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}